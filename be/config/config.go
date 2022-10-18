package config

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	JWTContextKey = "user"
)

var (
	// C 全局配置文件，在Init调用前为nil
	C *Config
)

const (
	DurationCodeExpire = time.Minute
)

// Config 配置
type Config struct {
	AppInfo     appInfo  `json:"app_info"`
	Postgres    postgres `json:"postgres"`
	Redis       redis    `json:"redis"`
	EnableDebug bool     `json:"debug"`
	JWT         jwt      `json:"jwt"`
	Wx          wx       `json:"wx"`
}

type appInfo struct {
	Addr      string `json:"addr"`
	APIPrefix string `json:"api_prefix"`
}

type postgres struct {
	Host     string `json:"host"`
	DB       string `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
}

type redis struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type jwt struct {
	Secret       string   `json:"secret"`
	SkipperPaths []string `json:"skipper_paths"`
}

type wx struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

func init() {
	configFile := "default.json"

	var data []byte
	var err error
	// 如果有设置 ENV ，则使用ENV中的环境
	if v, ok := os.LookupEnv("ENV"); ok {
		configFile = v + ".json"
		data, err = ioutil.ReadFile(fmt.Sprintf("config/%s", configFile))
	} else {
		data, err = ioutil.ReadFile(fmt.Sprintf("../env/config/%s", configFile))
	}

	// 读取配置文件

	if err != nil {
		log.Println("Read config error!")
		log.Panic(err)
		return
	}

	config := &Config{}

	err = jsoniter.Unmarshal(data, config)

	if err != nil {
		log.Println("Unmarshal config error!")
		log.Panic(err)
		return
	}

	C = config

	log.Println("Config " + configFile + " loaded.")
	if C.EnableDebug {
		log.Printf("%+v\n", C)
	}
}
