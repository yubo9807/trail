package router

import (
	"server/src/controller/file"
	"server/src/controller/stream"
	"server/src/controller/text"
	"server/src/controller/user"

	"github.com/gin-gonic/gin"
)

func Basic(r *gin.RouterGroup) {

	r.GET("/stream/sse", stream.EventSource)
	r.GET("/stream/test", stream.Test)

	// 用户登录
	r.POST("/user/signIn", user.SignIn)
	r.POST("/user/signUp", user.SignUp)
	r.POST("/token/refresh", user.RefreshToken)

	// 文件读取
	r.GET("/file/catalog", file.Catalog)
	r.GET("/file/read", file.ReadFile)

	// 文本处理
	r.POST("/text/md2html", text.MdToHTML)

}
