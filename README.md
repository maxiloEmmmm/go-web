# go-web
go好像完全没必要用框架

# 快速开始

```go
// main.go
contact.Init()
defer contact.Close()

// db := contact.InitDB("")
// defer db.Close()

InitRoute().Run(fmt.Sprintf(":%d", contact.Config.App.Port))

//route.go
r := gin.Default()
r.Use(contact.GinCors())
r.GET("/", contact.GinHelpHandle(func(c *contact.GinHelp) {
	c.Resource("hello world!")
}))
```

# 配置

新建config.yml
数据结构: contact/config.go

## database
> 默认环境 Config.App.Mode
```yaml
test:
  engine: mysql
  source: root:pass@tcp(localhost:3306)/test
release:
  engine: mysql
  source: root:pass@tcp(localhost:3306)/test
debug:
  engine: mysql
  source: root:pass@tcp(localhost:3306)/test
```


