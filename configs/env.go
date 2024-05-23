package configs

import (
	"os"
)

type EnvParams struct {
	DEVELOPMENT bool   // 是否属于开发环境
	ENVIRONMENT string // 自定义环境
}

var Env = EnvParams{}

func init() {

	length := len(os.Args)
	if length <= 1 {
		return
	}

	for i := 1; i < length; i++ {
		if i == 1 && os.Args[i] == "dev" {
			Env.DEVELOPMENT = true
		}
		if i == 2 {
			Env.ENVIRONMENT = os.Args[i]
		}
	}

}
