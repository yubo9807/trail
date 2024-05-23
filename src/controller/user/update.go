package user

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 注销
func Update(ctx *gin.Context) {
	type Params struct {
		Id       string `form:"id" binding:"required"`
		UserName string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
		RoleId   string `form:"role_id" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	spider.User.Update(params.Id, params.UserName, params.Password, params.RoleId)
	service.State.Success(ctx)
}
