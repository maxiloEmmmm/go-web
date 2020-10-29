package contact

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigMap struct {
	Database map[string]map[string]interface{}

	Jwt struct {
		Secret string
	}

	App struct {
		Port int
		Mode string
	}

	Redis struct {
		Host     string
		Port     int
		Db       int
		Password string
	}

	Log struct {
		Type string
		Info map[string]string
	}

	OpenTracing struct {
		Service string
		Sampler struct {
			Type  string
			Param float64
		}
		Reporter struct {
			LogSpans           bool
			LocalAgentHostPort string
		}
	}
}

var Config = ConfigMap{
	App: struct {
		Port int
		Mode string
	}{
		Port: 8000,
		Mode: "release",
	},
	OpenTracing: struct {
		Service string
		Sampler struct {
			Type  string
			Param float64
		}
		Reporter struct {
			LogSpans           bool
			LocalAgentHostPort string
		}
	}{
		Service: "App",
		Sampler: struct {
			Type  string
			Param float64
		}{
			Type:  "const",
			Param: 1,
		}, Reporter: struct {
			LogSpans           bool
			LocalAgentHostPort string
		}{
			LogSpans:           false,
			LocalAgentHostPort: "localhost:6831",
		},
	},
	Log: struct {
		Type string
		Info map[string]string
	}{Type: "file", Info: map[string]string{
		"path": "logs",
	}},
}

var ConfigPath = "./config.yaml"

func InitConfig() {
	configFile, err := ioutil.ReadFile(ConfigPath)

	if err != nil {
		return
	}

	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		panic(fmt.Sprintf("解析配置失败: %s", err))
	}
}

func ConfigFile(path string, dst interface{}) bool {
	configFile, err := ioutil.ReadFile(path)

	if err != nil {
		return false
	}

	err = yaml.Unmarshal(configFile, dst)
	if err != nil {
		return false
	}
	return true
}
