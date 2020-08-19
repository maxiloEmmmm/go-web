package contact

import (
	"github.com/gin-gonic/gin"
	go_tool "github.com/maxiloEmmmm/go-tool"
	"reflect"
)

type Model interface {
	List() (interface{}, int, error)
	Get(string) (interface{}, error)
	Patch(string, interface{}) error
	Fill() interface{}
	Create(interface{}) (interface{}, error)
	Delete(string) error
	PrimaryKey() string
}

func isPtr(s interface{}) bool {
	return reflect.ValueOf(s).Kind() == reflect.Ptr
}

// ref: https://developer.github.com/v3/#http-verbs
func CURD(r *gin.Engine, prefix string, model Model) {
	g := r.Group(go_tool.StringJoin("/", prefix))

	g.GET("", GinHelpHandle(func(c *GinHelp) {
		items, total, err := model.List()
		c.AssetsInValid("list", err)
		c.ResourcePage(items, total)
	}))

	g.GET("/:id", GinHelpHandle(func(c *GinHelp) {
		uri := struct {
			Id string `uri:"id"`
		}{}

		c.InValidBindUri(&uri)

		item, err := model.Get(uri.Id)
		c.AssetsInValid("get", err)
		c.Resource(item)
	}))

	g.POST("", GinHelpHandle(func(c *GinHelp) {
		fill := model.Fill()
		if !isPtr(fill) {
			c.InValidBind(&fill)
		} else {
			c.InValidBind(fill)
		}

		item, err := model.Create(fill)
		c.AssetsInValid("patch", err)
		c.ResourceCreate(item)
	}))

	g.PATCH("/:id", GinHelpHandle(func(c *GinHelp) {
		uri := struct {
			Id string `uri:"id"`
		}{}

		c.InValidBindUri(&uri)

		fill := model.Fill()
		if !isPtr(fill) {
			c.InValidBind(&fill)
		} else {
			c.InValidBind(fill)
		}

		c.AssetsInValid("patch", model.Patch(uri.Id, fill))
		c.Resource(nil)
	}))

	g.DELETE("/:id", GinHelpHandle(func(c *GinHelp) {
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
}

func (g GORMM) List() (interface{}, int, error) {
	items := g.ResolveList()

	if !isPtr(items) {
		items = &items
	}
	total := GinGormPageHelp(Db, items)
	return items, total, nil
}

func (g GORMM) Get(id string) (interface{}, error) {
	item := g.ResolveOne()

	if !isPtr(item) {
		item = &item
	}
	go_tool.AssetsError(Db.Where(go_tool.StringJoin(g.PrimaryKey(), " = ?"), id).First(item).Error)
	return item, nil
}

func (g GORMM) PrimaryKey() string {
	return "id"
}

func (g GORMM) Delete(id string) error {
	item := g.ResolveOne()

	if !isPtr(item) {
		item = &item
	}

	go_tool.AssetsError(Db.Where(go_tool.StringJoin(g.PrimaryKey(), " = ?"), id).First(item).Error)
	go_tool.AssetsError(Db.Delete(item).Error)
	return nil
}

func (g GORMM) Patch(id string, data interface{}) error {
	item := g.ResolveOne()

	if !isPtr(item) {
		item = &item
	}

	go_tool.AssetsError(Db.Where(go_tool.StringJoin(g.PrimaryKey(), " = ?"), id).First(item).Error)
	go_tool.AssetsError(Db.Model(item).Updates(data).Error)
	return nil
}
