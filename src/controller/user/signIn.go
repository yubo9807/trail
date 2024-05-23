package user

import (
	"server/src/service"
	"server/src/spider"
	"server/src/utils"

	"github.com/gin-gonic/gin"
)

// 登录
func SignIn(ctx *gin.Context) {
	type Params struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	// 验证用户名密码
	user := spider.User.Detail(params.Username)
	if user == nil {
		service.State.ErrorCustom(ctx, "用户名不存在")
		return
	}
	if utils.Md5Encipher(params.Password) != user.Pass {
		service.State.ErrorCustom(ctx, "密码错误")
		return
	}

	// 将一些重要信息存在 token 中
	info := map[string]interface{}{
		"roleId":   user.Role_id,
		"userId":   user.Id,
		"username": params.Username,
	}
	token := service.Jwt.Publish(info)
	service.State.SuccessData(ctx, token)
}
