package text

import (
	"server/src/service"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func MdToHTML(ctx *gin.Context) {
	type Params struct {
		Text string `form:"text" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	output := blackfriday.Run([]byte(params.Text))
	service.State.SuccessData(ctx, string(output))
}
