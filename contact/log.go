package contact

import (
	"context"
	"encoding/json"
	"fmt"
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/gin-gonic/gin"
	lib "github.com/maxiloEmmmm/go-tool"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
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
		{
			p, err := json.Marshal(li)
			if err != nil {
				LogLog(ErrorLevel, AppLogCode, err.Error())
			}
			return string(p)
		}
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
	fmt.Fprint(os.Stdout, li.RawString())
	configInstance.Write([]byte(li.String()))
	return li
}

func LogLog(level string, code string, message string) {
	fmt.Fprint(os.Stdout, (&LogInfo{
		Message: message,
		Code:    code,
		Time:    time.Now().Format("2006-01-02 15:04:05.000"),
		Level:   level,
	}).String())
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
		{
			return config.GetELog().Write(p)
		}
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
		client, err := elasticsearch.NewClient(elasticsearch.Config{
			Addresses: []string{Config.Log.Info["address"]},
		})

		if err != nil {
			log.Fatal(lib.StringJoin("elastic search err: ", err.Error()))
		}
		config.pipe = ElasticSearch{client: client, index: Config.Log.Info["index"]}
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

type ElasticSearch struct {
	client *elasticsearch.Client
	index  string
}

func (e ElasticSearch) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func (e ElasticSearch) Write(p []byte) (n int, err error) {
	go func() {
		res, err := esapi.IndexRequest{
			Index: e.index,
			Body:  strings.NewReader(string(p)),
		}.Do(context.Background(), e.client)

		if err != nil {
			LogLog(ErrorLevel, AppLogCode, err.Error())
			return
		}
		defer res.Body.Close()

		if res.IsError() {
			LogLog(ErrorLevel, AppLogCode, res.String())
		} else {
			// Deserialize the response into a map.
			var r map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				LogLog(ErrorLevel, AppLogCode, err.Error())
			}
		}
	}()
	return len(p), nil
}

func (e ElasticSearch) Close() error {
	return nil
}
