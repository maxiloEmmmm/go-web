package contact

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	go_tool "github.com/maxiloEmmmm/go-tool"
	"reflect"
	"strings"
)

type CURDOption struct {
	CreateFields []string
	UpdateFields []string
	Model        interface{}
	Instance     interface{}
	Prefix       string
	IdTransfer   func(string) reflect.Value
}

type curd struct {
	Option CURDOption
}

func NewEntCurd(option CURDOption) *curd {
	if option.UpdateFields == nil {
		option.UpdateFields = option.CreateFields
	}
	return &curd{option}
}

func (c *curd) Route(r *gin.Engine) *gin.RouterGroup {
	if !strings.HasPrefix("/", c.Option.Prefix) {
		c.Option.Prefix = go_tool.StringJoin("/", c.Option.Prefix)
	}
	g := r.Group(c.Option.Prefix)

	g.GET("", GinHelpHandle(func(help *GinHelp) {
		c.curdList(help)
	}))
	g.GET("/:id", GinHelpHandle(func(help *GinHelp) {
		c.curdOne(help)
	}))
	g.POST("", GinHelpHandle(func(help *GinHelp) {
		c.curdPost(help)
	}))
	g.PATCH("/:id", GinHelpHandle(func(help *GinHelp) {
		c.curdPatch(help)
	}))
	g.DELETE("/:id", GinHelpHandle(func(help *GinHelp) {
		c.curdDelete(help)
	}))

	return g
}

func (c *curd) curdList(help *GinHelp) {
	help.ResourcePage(func(start int, size int) (interface{}, int) {
		pipe := reflect.ValueOf(c.Option.Model).MethodByName("Query").Call(nil)[0]
		return pipe.MethodByName("All").
				Call([]reflect.Value{reflect.ValueOf(help.AppContext)})[0].Interface(),
			pipe.MethodByName("CountX").
				Call([]reflect.Value{reflect.ValueOf(help.AppContext)})[0].Interface().(int)
	})
}

func (c *curd) curdDelete(help *GinHelp) {
	uri := &struct {
		Id string `uri:"id"`
	}{}
	help.InValidBindUri(uri)
	reflect.ValueOf(c.Option.Model).MethodByName("DeleteOneID").Call([]reflect.Value{reflect.ValueOf(c.Option.IdTransfer(uri.Id).Interface())})[0].
		MethodByName("ExecX").Call([]reflect.Value{reflect.ValueOf(help.AppContext)})
	help.ResourceDelete()
}

func structForIn(v interface{}, cb func(key string, v reflect.Value)) {
	vOf := reflect.ValueOf(v)
	for _, key := range vOf.MapKeys() {
		cb(key.String(), vOf.MapIndex(key))
	}
}

func getInstanceByProto(v interface{}) interface{} {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return reflect.New(t).Elem().Interface()
}

func strFirstUpper(str string) string {
	vv := []rune(str)
	builder := strings.Builder{}
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			vv[i] -= 32
			builder.WriteRune(vv[i])
		} else {
			builder.WriteRune(vv[i])
		}
	}
	return builder.String()
}

func entSets(fill *reflect.Value, v interface{}, pickKeys []string) {
	structForIn(v, func(key string, v reflect.Value) {
		if go_tool.InArray(pickKeys, strings.ToLower(key)) {
			if method := fill.MethodByName(fmt.Sprintf("Set%s", strFirstUpper(key))); method.IsValid() {
				method.Call([]reflect.Value{reflect.ValueOf(v.Interface().(string))})
			}

			if key == "edges" {
				structForIn(v, func(key string, v reflect.Value) {
					if method := fill.MethodByName(fmt.Sprintf("Set%s", strFirstUpper(key))); method.IsValid() {
						method.Call([]reflect.Value{reflect.ValueOf(v.Interface().(string))})
					}
				})
			}
		}
	})
}

func (c *curd) curdPost(help *GinHelp) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case ResponseAbortError:
				panic(err)
			default:
				help.InValidError("validation", err.(error))
			}
		}
	}()
	body := &struct {
		Payload interface{}
	}{}
	help.InValidBind(body)

	tmpB, err := json.Marshal(body.Payload)
	if err != nil {
		help.InValidError("encode", err)
	}

	instance := getInstanceByProto(c.Option.Instance)
	err = json.Unmarshal(tmpB, &instance)
	if err != nil {
		help.InValidError("decode", err)
	}

	pipe := reflect.ValueOf(c.Option.Model).MethodByName("Create").Call(nil)[0]
	entSets(&pipe, body.Payload, c.Option.CreateFields)
	item := pipe.MethodByName("SaveX").Call([]reflect.Value{reflect.ValueOf(help.AppContext)})[0].Interface()
	help.Resource(item)
}

func (c *curd) curdPatch(help *GinHelp) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case ResponseAbortError:
				panic(err)
			default:
				help.InValidError("validation", err.(error))
			}
		}
	}()

	uri := struct {
		Id string `uri:"id"`
	}{}
	help.InValidBindUri(&uri)

	body := &struct {
		Payload interface{}
	}{}
	help.InValidBind(body)

	tmpB, err := json.Marshal(body.Payload)
	if err != nil {
		help.InValidError("encode", err)
	}
	err = json.Unmarshal(tmpB, &c.Option.Instance)
	if err != nil {
		help.InValidError("decode", err)
	}

	item := reflect.ValueOf(c.Option.Model).MethodByName("GetX").Call([]reflect.Value{
		reflect.ValueOf(help.AppContext),
		c.Option.IdTransfer(uri.Id),
	})[0].Interface()

	if item == nil {
		help.InValid("resource", "not found")
	} else {
		pipe := reflect.ValueOf(item).MethodByName("Update").Call(nil)[0]
		entSets(&pipe, body.Payload, c.Option.UpdateFields)
		item = pipe.MethodByName("SaveX").Call([]reflect.Value{reflect.ValueOf(help.AppContext)})[0].Interface()
	}
	help.Resource(item)
}

func (c *curd) curdOne(help *GinHelp) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case ResponseAbortError:
				panic(err)
			default:
				help.InValidError("validation", err.(error))
			}
		}
	}()
	uri := struct {
		Id string `uri:"id"`
	}{}
	help.InValidBindUri(&uri)
	help.Resource(
		reflect.ValueOf(c.Option.Model).MethodByName("GetX").Call([]reflect.Value{
			reflect.ValueOf(help.AppContext),
			c.Option.IdTransfer(uri.Id),
		})[0].Interface(),
	)
}
