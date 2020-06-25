package contact

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var configInstance configIO

var LogPath = "logs"

func InitLog() {
	configInstance := configIO{}
	gin.DefaultWriter = configInstance
}

type configIO struct {
	file *os.File
	key  string
}

func (config configIO) Write(p []byte) (n int, err error) {
	return io.MultiWriter(config.GetLogFile(), os.Stdout).Write(p)
}

func (config *configIO) Close() error {
	return config.file.Close()
}

func (config *configIO) GetLogFile() (file *os.File) {
	key := time.Now().Format("2006-01-02")

	if key != config.key || config.file == nil {
		if err := os.MkdirAll(LogPath, 0744); err != nil {
			log.Fatalln("日志文件夹创建失败: " + err.Error())
		}

		buffer := new(strings.Builder)
		buffer.WriteString(LogPath)
		buffer.WriteString("/access_")
		buffer.WriteString(key)
		buffer.WriteString(".log")
		file, err := os.OpenFile(buffer.String(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatalln("日志文件打开失败: " + err.Error())
		}

		if config.file != nil {
			config.file.Close()
		}

		config.file = file
		config.key = key
	}

	return config.file
}

func LogClose() error {
	return configInstance.Close()
}
