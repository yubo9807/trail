package user

import (
	"server/src/middleware"
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 获取用户信息
func GetInfo(ctx *gin.Context) {
	info := middleware.GetTokenInfo(ctx)
	user := spider.User.Detail(info["userId"].(string))
	user.Pass = ""
	service.State.SuccessData(ctx, user)
}
