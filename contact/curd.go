package contact

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	go_tool "github.com/maxiloEmmmm/go-tool"
	"reflect"
	"strings"
)

type CURDFilterFunc func(reflect.Value) reflect.Value
type CURDFilterCheck func(ginHelp *GinHelp, id string)

type CURDFilter struct {
	Create       CURDFilterFunc
	Patch        CURDFilterFunc
	List         CURDFilterFunc
	Delete       CURDFilterFunc
	DeleteBefore CURDFilterCheck
}

type CURDOption struct {
	CreateFields []string
	UpdateFields []string
	Model        interface{}
	Instance     interface{}
	Prefix       string
	IdTransfer   func(string) reflect.Value
	Filter       CURDFilter
	Order        []reflect.Value
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

func EntResourceCheck(help *GinHelp, source interface{}, tip string) {
	if source != nil {
		v := reflect.ValueOf(source)
		if v.IsNil() {
			return
		}

		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		help.InValid("resource", fmt.Sprintf("资源存在「%s(%s)」关联",
			tip,
			methodHelp(v.FieldByName("ID"), "String", nil)[0].Interface().(string),
		))
	}
}

func (c *curd) Route(r gin.IRouter) *gin.RouterGroup {
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
		pipe := methodHelp(reflect.ValueOf(c.Option.Model), "Query", nil)[0]

		if c.Option.Filter.List != nil {
			pipe = c.Option.Filter.List(pipe)
		}

		listPipe := methodHelp(pipe, "Clone", nil)[0]

		if len(c.Option.Order) > 0 {
			listPipe = methodHelp(listPipe, "Order", c.Option.Order)[0]
		}

		listPipe = methodHelp(listPipe, "Offset", []reflect.Value{reflect.ValueOf(start)})[0]
		listPipe = methodHelp(listPipe, "Limit", []reflect.Value{reflect.ValueOf(size)})[0]

		return methodHelp(listPipe, "All", []reflect.Value{reflect.ValueOf(help.AppContext)})[0].Interface(),
			methodHelp(pipe, "CountX", []reflect.Value{reflect.ValueOf(help.AppContext)})[0].Interface().(int)
	})
}

func (c *curd) curdDelete(help *GinHelp) {
	uri := &struct {
		Id string `uri:"id"`
	}{}
	help.InValidBindUri(uri)

	if c.Option.Filter.DeleteBefore != nil {
		c.Option.Filter.DeleteBefore(help, uri.Id)
	}
	methodHelp(
		methodHelp(reflect.ValueOf(c.Option.Model), "DeleteOneID", []reflect.Value{reflect.ValueOf(c.Option.IdTransfer(uri.Id).Interface())})[0],
		"ExecX",
		[]reflect.Value{reflect.ValueOf(help.AppContext)},
	)
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

func methodCallHelp(s reflect.Value, method string, args []reflect.Value, isSliceCall bool) []reflect.Value {
	if method := s.MethodByName(method); method.IsValid() {
		if isSliceCall {
			return method.CallSlice(args)
		} else {
			return method.Call(args)
		}
	} else {
		return nil
	}
}

func methodHelp(s reflect.Value, method string, args []reflect.Value) []reflect.Value {
	return methodCallHelp(s, method, args, false)
}

func methodSliceHelp(s reflect.Value, method string, args []reflect.Value) []reflect.Value {
	return methodCallHelp(s, method, args, true)
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
			methodHelp(*fill, fmt.Sprintf("Set%s", strFirstUpper(key)), []reflect.Value{reflect.ValueOf(v.Interface().(string))})

			if key == "edges" {
				structForIn(v, func(key string, v reflect.Value) {
					methodHelp(*fill, fmt.Sprintf("Set%s", strFirstUpper(key)), []reflect.Value{reflect.ValueOf(v.Interface().(string))})
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

	pipe := methodHelp(reflect.ValueOf(c.Option.Model), "Create", nil)[0]
	if c.Option.Filter.Create != nil {
		pipe = c.Option.Filter.Create(pipe)
	}
	entSets(&pipe, body.Payload, c.Option.CreateFields)
	item := methodHelp(pipe, "SaveX", []reflect.Value{reflect.ValueOf(help.AppContext)})[0].Interface()
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

	item := methodHelp(reflect.ValueOf(c.Option.Model), "GetX", []reflect.Value{
		reflect.ValueOf(help.AppContext),
		c.Option.IdTransfer(uri.Id),
	})[0].Interface()

	if item == nil {
		help.InValid("resource", "not found")
	} else {
		pipe := methodHelp(reflect.ValueOf(item), "Update", nil)[0]
		if c.Option.Filter.Patch != nil {
			pipe = c.Option.Filter.Patch(pipe)
		}
		entSets(&pipe, body.Payload, c.Option.UpdateFields)
		item = methodHelp(pipe, "SaveX", []reflect.Value{reflect.ValueOf(help.AppContext)})[0].Interface()
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
		methodHelp(reflect.ValueOf(c.Option.Model), "GetX", []reflect.Value{
			reflect.ValueOf(help.AppContext),
			c.Option.IdTransfer(uri.Id),
		})[0].Interface(),
	)
}
