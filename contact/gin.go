package contact

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GinHelp struct {
	*gin.Context
}

type GinHelpHandlerFunc func(c *GinHelp)

type H map[string]interface{}

func RouteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		jwt := JwtNew()
		jwt.SetSecret(Config.Jwt.Secret)

		if err := jwt.ParseToken(token); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		} else {
			c.Next()
		}
	}
}

func InitGin() {
	switch Config.App.Mode {
	case "debug", "":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

func GinHelpHandle(h GinHelpHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(&GinHelp{c})
	}
}

func (help *GinHelp) Response(code int, jsonObj interface{}) {
	help.AbortWithStatusJSON(code, jsonObj)
}

func (help *GinHelp) Resource(data interface{}) {
	help.Response(http.StatusOK, data)
}

func (help *GinHelp) ResourceCreate(data interface{}) {
	help.Response(http.StatusCreated, data)
}

func (help *GinHelp) ResourceDelete() {
	help.Response(http.StatusNoContent, "")
}

func (help *GinHelp) ResourceNotFound() {
	help.Response(http.StatusNotFound, "")
}

func (help *GinHelp) ServerError(msg string) {
	help.Response(http.StatusInternalServerError, msg)
}

type ValidInputSet struct {
	Msg   string
	Valid []InvalidError
}

func (help *GinHelp) InValidInput(msg string, data []InvalidError) {
	help.Response(http.StatusUnprocessableEntity, ValidInputSet{
		msg, data,
	})
}

type InvalidError struct {
	Field string
	Param interface{}
	Tag   string
	Msg   string
}

func (help *GinHelp) InValid(err error) {
	if errors, ve := err.(validator.ValidationErrors); ve {
		var errs []InvalidError
		for i := 0; i < len(errors); i++ {
			errItem := errors[i]
			errs = append(errs, InvalidError{
				Field: errItem.Field(),
				Param: errItem.Param(),
				Tag:   errItem.Tag(),
				Msg:   errItem.Translate(Tranintance.Tran),
			})
		}
		help.InValidInput("", errs)
	}
}

// 当post时 InValidBind 获取query显得苍白无力
func (help *GinHelp) InValidBindQuery(query interface{}) {
	help.InValid(help.ShouldBindQuery(query))
}

func (help *GinHelp) InValidBind(json interface{}) {
	help.InValid(help.ShouldBind(json))
}

func (help *GinHelp) Unauthorized(msg string) {
	help.Response(http.StatusUnauthorized, msg)
}

func (help *GinHelp) BadRequest(msg string) {
	help.Response(http.StatusBadRequest, msg)
}

func (help *GinHelp) Forbidden(msg string) {
	help.Response(http.StatusForbidden, msg)
}
