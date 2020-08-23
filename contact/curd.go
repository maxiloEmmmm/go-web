package contact

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	go_tool "github.com/maxiloEmmmm/go-tool"
	"reflect"
)

type Model interface {
	List(interface{}) (interface{}, int, error)
	Get(string) (interface{}, error)
	Patch(string, interface{}) error
	Fill(create bool) interface{}
	Create(interface{}) (interface{}, error)
	Delete(string) error
	PrimaryKey() string
	SetContext(ctx context.Context)
}

func isPtr(s interface{}) bool {
	return reflect.ValueOf(s).Kind() == reflect.Ptr
}

func isSlice(s interface{}) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		return v.Elem().Kind() == reflect.Slice
	}
	return v.Kind() == reflect.Slice
}

type Body struct {
	Payload interface{}
}

// ref: https://developer.github.com/v3/#http-verbs
func CURD(r *gin.Engine, prefix string, model Model) {
	g := r.Group(go_tool.StringJoin("/", prefix))

	g.GET("", GinHelpHandle(func(c *GinHelp) {
		model.SetContext(c.AppContext)
		items, total, err := model.List(nil)
		c.AssetsInValid("list", err)
		c.ResourcePage(items, total)
	}))

	g.GET("/:id", GinHelpHandle(func(c *GinHelp) {
		model.SetContext(c.AppContext)
		uri := struct {
			Id string `uri:"id"`
		}{}

		c.InValidBindUri(&uri)

		item, err := model.Get(uri.Id)
		c.AssetsInValid("get", err)
		c.Resource(item)
	}))

	g.POST("", GinHelpHandle(func(c *GinHelp) {
		model.SetContext(c.AppContext)
		fill := model.Fill(true)
		body := &Body{Payload: fill}
		c.InValidBind(body)

		item, err := model.Create(body.Payload)
		c.AssetsInValid("patch", err)
		c.ResourceCreate(item)
	}))

	g.PATCH("/:id", GinHelpHandle(func(c *GinHelp) {
		model.SetContext(c.AppContext)
		uri := struct {
			Id string `uri:"id"`
		}{}

		c.InValidBindUri(&uri)

		fill := model.Fill(false)
		body := &Body{Payload: fill}
		c.InValidBind(body)

		c.AssetsInValid("patch", model.Patch(uri.Id, body.Payload))
		c.Resource(nil)
	}))

	g.DELETE("/:id", GinHelpHandle(func(c *GinHelp) {
		model.SetContext(c.AppContext)
		uri := struct {
			Id string `uri:"id"`
		}{}

		c.InValidBindUri(&uri)

		c.AssetsInValid("delete", model.Delete(uri.Id))
		c.ResourceDelete()
	}))

	//put is miss..... hhhhh
}

type GORMM struct {
	ResolveOne  func() interface{}
	ResolveList func() interface{}
	UsePage     bool
	context.Context
}

func (g GORMM) SetContext(ctx context.Context) {
	g.Context = ctx
}

func (g GORMM) PageHelp(db *gorm.DB, data interface{}) int {
	return g.Context.Value("app").(*GinHelp).GinGormPageHelp(db, data)
}

func (g GORMM) List(where interface{}) (interface{}, int, error) {
	items := g.ResolveList()

	if !isSlice(items) {
		return nil, 0, errors.New("data collection not slice")
	}

	total := 0

	db := Db
	if where != nil {
		db = db.Where(where)
	}
	if g.UsePage {
		total = g.PageHelp(db, items)
	} else {
		go_tool.AssetsError(db.Find(items).Error)
		total = reflect.ValueOf(items).Elem().Len()
	}

	return items, total, nil
}

func (g GORMM) Get(id string) (interface{}, error) {
	item := g.ResolveOne()

	if !isPtr(item) {
		return nil, errors.New("data collection not ptr")
	}
	if err := Db.Where(go_tool.StringJoin(g.PrimaryKey(), " = ?"), id).First(item).Error; gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("not found")
	} else {
		go_tool.AssetsError(err)
	}
	return item, nil
}

func (g GORMM) PrimaryKey() string {
	return "id"
}

func (g GORMM) Create(data interface{}) (interface{}, error) {
	item := g.ResolveOne()

	if !isPtr(item) {
		return nil, errors.New("data collection not ptr")
	}

	go_tool.AssetsError(Db.Create(data).Error)
	return data, nil
}

func (g GORMM) Delete(id string) error {
	item := g.ResolveOne()

	if !isPtr(item) {
		return errors.New("data collection not ptr")
	}

	if err := Db.Where(go_tool.StringJoin(g.PrimaryKey(), " = ?"), id).First(item).Error; gorm.IsRecordNotFoundError(err) {
		return errors.New("not found")
	} else {
		go_tool.AssetsError(err)
	}
	go_tool.AssetsError(Db.Delete(item).Error)
	return nil
}

func (g GORMM) Patch(id string, data interface{}) error {
	item := g.ResolveOne()

	if !isPtr(item) {
		return errors.New("data collection not ptr")
	}

	if err := Db.Where(go_tool.StringJoin(g.PrimaryKey(), " = ?"), id).First(item).Error; gorm.IsRecordNotFoundError(err) {
		return errors.New("not found")
	} else {
		go_tool.AssetsError(err)
	}
	go_tool.AssetsError(Db.Model(item).Updates(data).Error)
	return nil
}
