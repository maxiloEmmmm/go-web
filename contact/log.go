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
	if err := os.MkdirAll("logs", 0744); err != nil {
		log.Fatalln("日志文件夹创建失败: " + err.Error())
	}

	f, err := os.OpenFile(fmt.Sprintf("logs/access_%s.log", time.Now().Format("2006-01-02")), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		log.Fatalln("日志文件打开失败: " + err.Error())
	}

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
