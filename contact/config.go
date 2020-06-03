package contact

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigMap struct {
	Database struct {
		Host     string
		Port     int
		Prefix   string
		Database string
		Username string
		Password string
	}

	Jwt struct {
		Secret    string
		LifeCycle string `json:"life_cycle"`
	}

	App struct {
		Port int
		Mode string
	}
}

var Config ConfigMap

func InitConfig() {
	configFile, err := ioutil.ReadFile("./config.yml")

	if err != nil {
		panic(fmt.Sprintf("打开文件失败: %s", err))
	}

	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		panic(fmt.Sprintf("解析配置失败: %s", err))
	}

	if Config.App.Port == 0 {
		Config.App.Port = 8080
	}
}
