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

## <span id="generate">Generate</span>
1. entc init User
2. edit User schema (must include id)
    ```go
    func (User) Fields() []ent.Field {
        return []ent.Field{
            field.Int("id"),
            field.String("username"),
            field.String("password"),
        }
    }
    ```
3. touch gen.go
    ```go
    //go:generate go run github.com/maxiloEmmmm/go-web/generate {{ent_schema_path}}
    ```
4. generate
    ```shell script
    go generate gen.go
    ```

## curd

### gin
```go
// ent
client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
if err != nil {
    log.Fatalf("failed opening connection to sqlite: %v", err)
}
defer client.Close()
// Run the auto migration tool.
if err := client.Schema.Create(context.Background()); err != nil {
    log.Fatalf("failed creating schema resources: %v", err)
}

engine := gin.Default()
curd := ent.NewCurdBuilder(client)
curd.Route("/api", engine)
engine.Run(":8000")
```
    
```shell script
[GIN-debug] GET    /api/user                 --> github.com/maxiloEmmmm/go-web/contact.GinHelpHandle.func1 (3 handlers)
[GIN-debug] GET    /api/user/:id             --> github.com/maxiloEmmmm/go-web/contact.GinHelpHandle.func1 (3 handlers)
[GIN-debug] POST   /api/user                 --> github.com/maxiloEmmmm/go-web/contact.GinHelpHandle.func1 (3 handlers)
[GIN-debug] PATCH  /api/user/:id             --> github.com/maxiloEmmmm/go-web/contact.GinHelpHandle.func1 (3 handlers)
[GIN-debug] DELETE /api/user/:id             --> github.com/maxiloEmmmm/go-web/contact.GinHelpHandle.func1 (3 handlers)
```


