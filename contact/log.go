package contact

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

func init() {
	if err := os.MkdirAll("logs", 0777); err != nil {
		log.Fatalln("日志文件夹创建失败: " + err.Error())
	}

	f, err := os.OpenFile(fmt.Sprintf("logs/%s-access.log", time.Now().Format("2020-01-01")), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		log.Fatalln("日志文件打开失败: " + err.Error())
	}

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
