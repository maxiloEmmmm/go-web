# go-web
go好像完全没必要用框架

# 快速开始

```go
// main.go
contact.Init()
defer contact.Close()

// contact.InitDB()
// defer contact.DbClose()

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

# curd

### 数据结构

```sql
CREATE TABLE `post`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `context` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;
```

### 模型

```go
// model/post.go
package model

import (
	"github.com/maxiloEmmmm/go-web/contact"
)

type Post struct {
	Id int
	Title string
	Context string
}

type PostModel struct {
	*contact.GORMM
}

func NewPostModel() *PostModel {
	return &PostModel{
		GORMM: &contact.GORMM{
			ResolveList: func() interface{} {
				tmp := make([]Post, 0)
				return &tmp
			},
			ResolveOne: func() interface{} {
				return new(Post)
			},
		},
	}
}

func (post PostModel)Fill(create bool) interface{} {
	return &Post{}
}
```

### 路由

```go
// route.go
contact.CURD(r, "post", model.NewPostModel())
```

### 测试

#### http 配置文件

```json
{
  "dev": {
    "ct": "application/json",
    "host": "http://localhost:8000/"
  }
}
```

#### http请求文件
```http request
### create
POST {{host}}/post
Content-Type: {{ct}}

{"payload": {"title": "3"}}

### list
GET {{host}}/post

### get
GET {{host}}/post/1

### update
PATCH {{host}}/post/1
Content-Type: {{ct}}

{"payload": {"title": "4"}}

### del
DELETE {{host}}/post/1
Content-Type: {{ct}}
```

