package user

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 注销
func Delete(ctx *gin.Context) {
	type Params struct {
		Id string `form:"id" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	spider.User.Delete(params.Id)
	service.State.Success(ctx)
}
