package user

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 获取用户列表
func List(ctx *gin.Context) {
	type Params struct {
		PageNumber int    `form:"pageNumber" binding:"required"`
		PageSize   int    `form:"pageSize" binding:"required"`
		Name       string `form:"name"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	users := spider.User.List(params.PageNumber, params.PageSize, params.Name)
	service.State.SuccessData(ctx, users)
}
