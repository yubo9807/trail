package user

import (
	"server/src/service"
	"server/src/spider"
	"server/src/utils"

	"github.com/gin-gonic/gin"
)

// 注册
func SignUp(ctx *gin.Context) {
	type Params struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	user := spider.User.Detail(params.Username)
	if user != nil {
		service.State.ErrorCustom(ctx, "用户名已存在")
		return
	}

	pass := utils.Md5Encipher(params.Password)
	spider.User.Add(params.Username, pass, "10")
	service.State.Success(ctx)
}
