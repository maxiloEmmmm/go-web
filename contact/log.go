package contact

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	lib "github.com/maxiloEmmmm/go-tool"
	"github.com/olivere/elastic/v7"
	"io"
	"log"
	"os"
	"runtime"
	"time"
)

var configInstance configIO

var (
	Info    *LogInfo
	Warning *LogInfo
	Error   *LogInfo
)

func InitLog() {
	configInstance = configIO{}

	system := &SystemAdapter{}
	gin.DefaultWriter = system
	log.SetOutput(system)

	Info = &LogInfo{Level: InfoLevel}
	Warning = &LogInfo{Level: WarnLevel}
	Error = &LogInfo{Level: ErrorLevel}
}

type configIO struct {
	pipe io.ReadWriteCloser
	key  string
}

type LogInfo struct {
	Message string
	Code    string
	Time    string
	Level   string
	Line    int    `json:",omitempty"`
	File    string `json:",omitempty"`
}

func (li *LogInfo) RawString() string {
	tmp := fmt.Sprintf(
		"time=%s level=%s code=%s message=%s",
		li.Time,
		li.Level,
		li.Code,
		li.Message,
	)

	if li.Line > 0 {
		tmp = fmt.Sprintf("%s file=%s line=%d", tmp, li.File, li.Line)
	}
	return tmp
}

func (li *LogInfo) String() string {
	switch Config.Log.Type {
	case "elastic_search":
		p, err := json.Marshal(li)
		if err != nil {
			LogLog(ErrorLevel, AppLogCode, err.Error())
		}
		return string(p)
	case "file":
		fallthrough
	default:
		return li.RawString()
	}
}

func (li *LogInfo) Log(code string, message string) *LogInfo {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		li.File = file
		li.Line = line
	}

	li.Code = code
	li.Message = message
	li.Time = time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Fprintf(os.Stdout, "%s\n", li.RawString())
	configInstance.Write([]byte(li.String()))
	return li
}

func LogLog(level string, code string, message string) {
	fmt.Fprintf(os.Stdout, "%s\n", (&LogInfo{
		Message: message,
		Code:    code,
		Time:    time.Now().Format("2006-01-02 15:04:05.000"),
		Level:   level,
	}).RawString())
}

type SystemAdapter struct{}

func (e SystemAdapter) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func (e SystemAdapter) Write(p []byte) (n int, err error) {
	Info.Log(AppLogCode, string(p))
	return len(p), nil
}

func (e SystemAdapter) Close() error {
	return nil
}

const (
	AppLogCode = "app"
)

const (
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
)

func (config configIO) Write(p []byte) (n int, err error) {
	switch Config.Log.Type {
	case "elastic_search":
		return config.GetELog().Write(p)
	case "file":
		fallthrough
	default:
		return config.GetFileLog().Write(p)
	}
}

func (config *configIO) Close() error {
	return config.pipe.Close()
}

func (config *configIO) GetELog() io.ReadWriteCloser {
	if config.pipe == nil {
		client, err := elastic.NewClient(
			elastic.SetURL(Config.Log.Info["address"]),
		)
		if err != nil {
			log.Fatal(err.Error())
		}

		_, _, err = client.Ping(Config.Log.Info["address"]).Do(context.Background())
		if err != nil {
			log.Fatal(err.Error())
		}

		index := Config.Log.Info["index"]
		exists, err := client.IndexExists(index).Do(context.Background())
		if err != nil {
			log.Fatal(err.Error())
		}
		if !exists {
			_, err := client.CreateIndex("twitter").BodyString(EsMapping).Do(context.Background())
			if err != nil {
				log.Fatal(err.Error())
			}
		}
		config.pipe = ElasticSearch{client: client, index: index}
	}
	return config.pipe
}

func (config *configIO) GetFileLog() (file *os.File) {
	key := time.Now().Format("2006-01-02")

	if key != config.key || config.pipe == nil {
		if err := os.MkdirAll(Config.Log.Info["path"], 0744); err != nil {
			log.Fatalln("日志文件夹创建失败: " + err.Error())
		}

		file, err := os.OpenFile(lib.StringJoin(Config.Log.Info["path"], "/access_", key, ".log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatalln("日志文件打开失败: " + err.Error())
		}

		if config.pipe != nil {
			config.pipe.Close()
		}

		config.pipe = file
		config.key = key
	}

	return config.pipe.(*os.File)
}

func LogClose() error {
	return configInstance.Close()
}

const EsMapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"properties":{
			"@timestamp":{
				"format":"strict_date_optional_time||epoch_millis",
				"type":"date"
				"enabled":true
			}
		}
	}
}`

type ElasticSearch struct {
	client *elastic.Client
	index  string
}

func (e ElasticSearch) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func (e ElasticSearch) Write(p []byte) (n int, err error) {
	go func() {
		_, err := e.client.Index().Index(e.index).BodyString(string(p)).Do(context.Background())
		if err != nil {
			LogLog(ErrorLevel, AppLogCode, err.Error())
			return
		}
	}()
	return len(p), nil
}

func (e ElasticSearch) Close() error {
	return nil
}
