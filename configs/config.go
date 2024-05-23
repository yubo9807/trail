package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigType struct {
	Port           int
	Prefix         string
	LogReserveTime int `yaml:"logReserveTime"`

	TokenValidTime         int64 `yaml:"tokenValidTime"`
	TokenExceedRefreshTime int64 `yaml:"tokenExceedRefreshTime"`

	SqlSecret string `yaml:"sqlSecret"`

	AssetsUrl string `yaml:"assetsUrl"`
}

var Config ConfigType

const template = `
prefix: "/trail"  # 路由前缀
port: 9528  # 启动端口
logReserveTime: 30  # 日志保留时间(d)

tokenValidTime: 7200  # 令牌有效时间(s)
tokenExceedRefreshTime: 86400  # 令牌超过刷新时间(s)

sqlSecret: "user:password@tcp(0.0.0.0:3306)/database?timeout=5s"  # sql 密匙

assetsUrl: ""  # 开放资源 url
`

func init() {
	configFile := "./config.yml"
	data, err := os.ReadFile(configFile)
	if err != nil {
		os.Create(configFile)
		os.WriteFile(configFile, []byte(template), 0777)
		data, _ = os.ReadFile(configFile)
	}

	if err := yaml.Unmarshal([]byte(data), &Config); err != nil {
		panic(err.Error())
	}
}
