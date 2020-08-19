package contact

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/gorm"
	lib "github.com/maxiloEmmmm/go-tool"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type page struct {
	PageKey         string
	PageSizeKey     string
	PageSizeDefault int
	CurrentDefault  int
}

type pageInfo struct {
	Current int
	Size    int
}

var GinPage = page{PageKey: "page", PageSizeKey: "page_size", PageSizeDefault: 15, CurrentDefault: 1}

func (help *GinHelp) GetPageInfo() pageInfo {
	var gp pageInfo
	if val, exist := help.Get("page"); exist {
		gp = val.(pageInfo)
	} else {
		gp = pageInfo{
			Current: GinPage.CurrentDefault,
			Size:    GinPage.PageSizeDefault,
		}
	}
	return gp
}

func (help *GinHelp) GinGormPageHelp(db *gorm.DB, data interface{}) int {
	var gp = help.GetPageInfo()
	return GinGormPageBase(db, data, gp.Current, gp.Size)
}

func (help *GinHelp) GinGormPageHelpWithOptionSize(db *gorm.DB, data interface{}, size int) int {
	var gp = help.GetPageInfo()
	return GinGormPageBase(db, data, gp.Current, size)
}

func GinGormPageHelpWithOption(db *gorm.DB, data interface{}, current int, size int) int {
	return GinGormPageBase(db, data, current, size)
}

func GinGormPageBase(db *gorm.DB, data interface{}, current int, size int) (total int) {
	lib.AssetsError(db.Model(data).Count(&total).Error)
	lib.AssetsError(db.Offset((current - 1) * size).Limit(size).Find(data).Error)
	return
}

type Cors struct {
	AllowOrigin  []string
	AllowHeaders []string
	AllowMethods []string
	// response header without Cache-Control、Content-Language、Content-Type、Expires、Last-Modified、Pragma
	// if client will get...
	ExposeHeaders    []string
	AllowCredentials bool
}

var CorsConfig = Cors{
	AllowOrigin:      []string{"*"},
	AllowHeaders:     []string{"Content-Type", "AccessToken", "X-CSRF-Token", "Authorization", "X-Requested-With"},
	AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "PATCH", "DELETE"},
	AllowCredentials: false,
}

func GinCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.GetHeader("Origin")
		if lib.InArray(CorsConfig.AllowOrigin, "*") || lib.InArray(CorsConfig.AllowOrigin, origin) {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Headers", strings.Join(CorsConfig.AllowHeaders, ","))
			c.Header("Access-Control-Allow-Methods", strings.Join(CorsConfig.AllowMethods, ","))
			c.Header("Access-Control-Expose-Headers", strings.Join(CorsConfig.ExposeHeaders, ","))
			c.Header("Access-Control-Allow-Credentials", lib.AssetsReturn(CorsConfig.AllowCredentials, "true", "false").(string))
			if method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
			}
		}
		c.Next()
	}
}

func InitOpenTracing() (opentracing.Tracer, io.Closer) {
	cfg, err := jaegercfg.FromEnv()
	cfg.ServiceName = Config.OpenTracing.Service
	cfg.Sampler.Type = Config.OpenTracing.Sampler.Type
	cfg.Sampler.Param = Config.OpenTracing.Sampler.Param
	cfg.Reporter.LogSpans = Config.OpenTracing.Reporter.LogSpans
	cfg.Reporter.LocalAgentHostPort = Config.OpenTracing.Reporter.LocalAgentHostPort
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		gin.DefaultWriter.Write([]byte(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err)))
	}

	// opentracing.setGlobalTracer is singleton, not better
	return tracer, closer
}

func SpanWrap(tracer opentracing.Tracer, key string, span opentracing.Span) opentracing.Span {
	var s opentracing.Span
	if span != nil {
		s = tracer.StartSpan(key, opentracing.ChildOf(span.Context()))
	} else {
		s = tracer.StartSpan(key)
	}

	return s
}

func ClientSpanWrap(tracer opentracing.Tracer, span opentracing.Span) *resty.Client {
	client := resty.New()
	client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		s := SpanWrap(tracer, request.Method, span)
		s.SetTag("url", request.URL)
		client.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
			s.Finish()
			return nil
		})
		tracer.Inject(
			s.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(request.Header))
		return nil
	})
	return client
}

func GinOpenTracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer, closer := InitOpenTracing()
		var serverSpan opentracing.Span
		wireContext, err := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			serverSpan = tracer.StartSpan("request")
		} else {
			serverSpan = tracer.StartSpan("request", opentracing.ChildOf(wireContext))
		}
		c.Set("open-tracing.span", serverSpan)
		c.Set("open-tracing.tracer", tracer)
		defer func() {
			serverSpan.Finish()
			closer.Close()
		}()
		c.Next()
	}
}

func InitGin() {
	switch Config.App.Mode {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

var ServerErrorWrite = new(ServerErrorIO)

type ServerErrorIO struct{}

func (sew ServerErrorIO) Write(p []byte) (n int, err error) {
	buffer := new(strings.Builder)
	buffer.WriteString("[SERVER_ERROR]:")
	buffer.Write(p)
	return gin.DefaultWriter.Write([]byte(buffer.String()))
}

func GinHelpHandle(h GinHelpHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		gp := pageInfo{Current: 1, Size: 10}

		if page, err := strconv.Atoi(c.DefaultQuery(GinPage.PageKey, "1")); err == nil {
			gp.Current = page
		}

		if pageSize, err := strconv.Atoi(c.DefaultQuery(GinPage.PageSizeKey, string(GinPage.PageSizeDefault))); err == nil {
			gp.Size = pageSize
		}

		c.Set("page", gp)

		help := &GinHelp{Context: c, AppContext: context.Background()}
		help.AppContext = context.WithValue(help.AppContext, "app", help)
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
						ServerErrorWrite.Write([]byte(lib.StringJoin(md5, "-", errMsg)))
						if span, exist := c.Get("open-tracing.span"); exist {
							span.(opentracing.Span).LogKV(
								"msg", errMsg,
								"error_code", md5,
							)
							span.(opentracing.Span).SetTag("status", "error")
						}
						c.AbortWithStatusJSON(http.StatusUnprocessableEntity, InValidFunc("server", md5))
					}
				}
			}
		}(help)
		h(help)
	}
}

type ResponseAbortError struct{}

func (r ResponseAbortError) Error() string {
	return "abort"
}

type GinHelp struct {
	*gin.Context
	AppContext context.Context
}

type GinHelpHandlerFunc func(c *GinHelp)

type H map[string]interface{}

func GinRouteAuth() gin.HandlerFunc {
	return GinHelpHandle(func(c *GinHelp) {
		token := c.GetToken()

		jwt := JwtNew()

		jwt.SetSecret(Config.Jwt.Secret)

		if err := jwt.ParseToken(token); err != nil {
			c.InValidError("auth:token:parse", err)
		} else {
			c.Set("auth", jwt)
			c.Next()
		}
	})
}

func (help *GinHelp) SpanWrapWithApp(key string) opentracing.Span {
	return SpanWrap(help.TracingTracer(), key, help.TracingSpan())
}

func (help *GinHelp) ClientSpanWrapWithApp() *resty.Client {
	return ClientSpanWrap(help.TracingTracer(), help.TracingSpan())
}

func (help *GinHelp) TracingSpan() opentracing.Span {
	if context, exist := help.Get("open-tracing.span"); exist {
		return context.(opentracing.Span)
	} else {

		return nil
	}
}

func (help *GinHelp) TracingTracer() opentracing.Tracer {
	if context, exist := help.Get("open-tracing.tracer"); exist {
		return context.(opentracing.Tracer)
	} else {

		return nil
	}
}

// 响应
func (help *GinHelp) Response(code int, jsonObj interface{}) {
	help.AbortWithStatusJSON(code, jsonObj)
	panic(ResponseAbortError{})
}

// 分页响应辅助
func (help *GinHelp) ResourcePage(data interface{}, total int) {
	help.Resource(H{
		"data":  data,
		"total": total,
	})
}

// 成功响应
func (help *GinHelp) Resource(data interface{}) {
	help.Response(http.StatusOK, data)
}

// 创建成功响应
func (help *GinHelp) ResourceCreate(data interface{}) {
	help.Response(http.StatusCreated, data)
}

// 删除成功响应
func (help *GinHelp) ResourceDelete() {
	help.Response(http.StatusNoContent, "")
}

// 资源丢失响应 不推荐使用 太过于底层 推荐InValid*
func (help *GinHelp) ResourceNotFound() {
	help.Response(http.StatusNotFound, "")
}

type GinInValidFunc func(code string, message string) interface{}

var InValidFunc GinInValidFunc = defaultInValidFunc

var defaultInValidFunc GinInValidFunc = func(code string, message string) interface{} {
	return H{
		"code":    code,
		"message": message,
	}
}

// 客户端错误响应
func (help *GinHelp) InValid(code string, msg string) {
	help.Response(http.StatusUnprocessableEntity, InValidFunc(code, msg))
}

// 断言客户端错误
func (help *GinHelp) AssetsInValid(code string, err error) {
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			{
				if errors, ok := err.(validator.ValidationErrors); ok && len(errors) > 0 {
					e := errors[0]
					help.InValid(code, lib.StringJoin(
						e.Translate(Tranintance.Tran),
						", 类型: ", e.Type().Name(),
						", 当前值: ", fmt.Sprintf("%v", e.Value()),
						", 数据结构路径: ", e.StructNamespace()))
				} else {
					help.InValidError(code, err)
				}
			}
		default:
			help.InValidError(code, err)
		}
	}
}

// 客户端错误响应
func (help *GinHelp) InValidError(code string, err error) {
	help.InValid(code, err.Error())
}

// 客户端query错误响应
func (help *GinHelp) InValidBindQuery(query interface{}) {
	help.AssetsInValid("input:query", help.ShouldBindQuery(query))
}

// 客户端uri错误响应
func (help *GinHelp) InValidBindUri(query interface{}) {
	help.AssetsInValid("input:uri", help.ShouldBindUri(query))
}

// 客户端body错误响应
func (help *GinHelp) InValidBind(json interface{}) {
	help.AssetsInValid("input:body", help.ShouldBind(json))
}

// 客户端未认证 不推荐使用 太过于底层 推荐InValid*
func (help *GinHelp) Unauthorized(msg string) {
	help.Response(http.StatusUnauthorized, map[string]interface{}{"msg": msg})
}

// 客户端错误请求 不推荐使用 太过于底层 推荐InValid*
func (help *GinHelp) BadRequest(msg string) {
	help.Response(http.StatusBadRequest, map[string]interface{}{"msg": msg})
}

// 客户端权限不足 不推荐使用 太过于底层 推荐InValid*
func (help *GinHelp) Forbidden(msg string) {
	help.Response(http.StatusForbidden, map[string]interface{}{"msg": msg})
}

// 服务端错误响应 不推荐使用 太过于底层 推荐InValid*
func (help *GinHelp) ServerError(msg string) {
	help.Response(http.StatusInternalServerError, map[string]interface{}{"msg": msg})
}

// 获取token
func (help *GinHelp) GetToken() string {
	token := help.GetHeader("Authorization")

	if len(token) == 0 {
		token, _ = help.GetQuery("token")
	}

	return token
}
