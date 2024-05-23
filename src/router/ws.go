package router

import (
	"server/src/controller/chat"
	"server/src/controller/ws"

	"github.com/gin-gonic/gin"
)

func Ws(r *gin.RouterGroup) {
	r.GET("/chat/*any", gin.WrapH(chat.Chat()))
	r.GET("/test", ws.Test)
}
