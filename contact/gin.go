package contact

import (
	"bytes"
	"fmt"
	lib "github.com/maxiloEmmmm/go-tool"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GinHelp struct {
	*gin.Context
}

type GinHelpHandlerFunc func(c *GinHelp)

type H map[string]interface{}

func GinRouteAuth() gin.HandlerFunc {
	return GinHelpHandle(func(c *GinHelp) {
		token := c.GetToken()

		jwt := JwtNew()

		jwt.SetSecret(Config.Jwt.Secret)

		if err := jwt.ParseToken(token); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		} else {
			c.Set("auth", jwt)
			c.Next()
		}
	})
}

func GinCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
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

var ServerErrorWrite = new(ServerErrorIO)

type ServerErrorIO struct{}

func (sew ServerErrorIO) Write(p []byte) (n int, err error) {
	buffer := new(bytes.Buffer)
	buffer.Write([]byte("[SERVER_ERROR]:"))
	buffer.Write(p)
	return gin.DefaultWriter.Write(buffer.Bytes())
}

func GinHelpHandle(h GinHelpHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		help := &GinHelp{c}
		defer func(c *GinHelp) {
			if err := recover(); err != nil {
				switch err.(type) {
				case ResponseAbortError:
					{
						return
					}
				default:
					{
						errMsg := ""
						if e, ok := err.(error); ok {
							errMsg = e.(error).Error()
						} else {
							errMsg = fmt.Sprintf("%+v", err)
						}

						md5 := lib.Md5(fmt.Sprintf("%d%s", time.Now().Unix(), errMsg))
						buffer := new(bytes.Buffer)
						buffer.WriteString(md5)
						buffer.WriteString("-")
						buffer.WriteString(errMsg)
						ServerErrorWrite.Write(buffer.Bytes())
						c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"msg": md5})
					}
				}
			}
		}(help)
		h(help)
	}
}

type ResponseAbortError struct{}

func Error() string {
	return "abort"
}

func (help *GinHelp) Response(code int, jsonObj interface{}) {
	help.AbortWithStatusJSON(code, jsonObj)
	panic(ResponseAbortError{})
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
	help.Response(http.StatusInternalServerError, map[string]interface{}{"msg": msg})
}

func (help *GinHelp) InValidInput(msg string, error string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString(msg)
	buffer.WriteString(", ")
	buffer.WriteString(error)
	help.Response(http.StatusUnprocessableEntity, buffer.String())
}

func (help *GinHelp) InValid(msg string, error error) {
	if error != nil {
		help.InValidInput(msg, error.Error())
	}
}

// 当post时 InValidBind 获取query显得苍白无力
func (help *GinHelp) InValidBindQuery(query interface{}) {
	help.InValid("query bind err", help.ShouldBindQuery(query))
}

func (help *GinHelp) InValidBindUri(query interface{}) {
	help.InValid("uri bind err", help.ShouldBindUri(query))
}

func (help *GinHelp) InValidBind(json interface{}) {
	help.InValid("body bind err", help.ShouldBind(json))
}

func (help *GinHelp) Unauthorized(msg string) {
	help.Response(http.StatusUnauthorized, map[string]interface{}{"msg": msg})
}

func (help *GinHelp) BadRequest(msg string) {
	help.Response(http.StatusBadRequest, map[string]interface{}{"msg": msg})
}

func (help *GinHelp) Forbidden(msg string) {
	help.Response(http.StatusForbidden, map[string]interface{}{"msg": msg})
}

func (help *GinHelp) GetToken() string {
	token := help.GetHeader("Authorization")

	if len(token) == 0 {
		token, _ = help.GetQuery("token")
	}

	return token
}
