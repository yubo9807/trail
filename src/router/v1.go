package router

import (
	"server/src/controller/test"
	"server/src/controller/user"
	"server/src/middleware"

	"github.com/gin-gonic/gin"
)

func V1(r *gin.RouterGroup) {
	r.Use(middleware.Authorization)

	r.GET("/user/list", user.List)
	r.GET("/user/current", user.GetInfo)
	r.POST("/user/update", user.Update)
	r.DELETE("/user/delete", user.Delete)

	r.GET("/test", test.Test)
}
