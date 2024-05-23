package main

import (
	"server/configs"
	"server/src/middleware"
	"server/src/router"
	"strconv"

	"github.com/gin-gonic/gin"
)

func server() *gin.Engine {

	app := gin.Default()
	app.Use(middleware.Core)

	// 代理应用
	power := app.Group("/permissions")
	power.Use(middleware.BodyDispose)
	power.Use(middleware.Authorization)
	power.Use(middleware.RoleVerify("0")) // 指定特定的角色可以调以下接口
	power.Any("/*path", middleware.ProxyPermissions)

	// 自身应用
	base := app.Group(configs.Config.Prefix)
	base.Use(middleware.Recover)
	base.Use(middleware.Logs)
	base.Use(middleware.BodyDispose)

	router.Basic(base.Group("/basic/api"))
	router.V1(base.Group("/v1/api"))
	router.Ws(base.Group("/ws"))

	return app
}

func main() {

	port := ":" + strconv.Itoa(configs.Config.Port)
	server().Run(port)

}
