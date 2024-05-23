package user

import (
	"encoding/base64"
	"encoding/json"
	"server/configs"
	"server/src/service"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 刷新 token
func RefreshToken(ctx *gin.Context) {
	type Params struct {
		Token string `binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	_, err := service.Jwt.Verify(params.Token)
	if err != nil && err.Error() != "Token is expired" {
		service.State.ErrorUnauthorized(ctx, err.Error())
		return
	}

	startIndex := strings.Index(params.Token, ".") + 1
	endIndex := startIndex + strings.Index(params.Token[startIndex:], ".")
	tokenBody := params.Token[startIndex:endIndex]
	deficiency := strings.Repeat("=", 4-len(tokenBody)%4)
	decodedData, err1 := base64.StdEncoding.DecodeString(tokenBody + deficiency)
	if err1 != nil {
		service.State.ErrorCustom(ctx, err1.Error())
		return
	}

	var data map[string]interface{}
	err2 := json.Unmarshal(decodedData, &data)
	if err2 != nil {
		service.State.ErrorCustom(ctx, err2.Error())
		return
	}

	// 超过__时间未登录，拒绝更新 token，通知退出
	poor := time.Now().Unix() - int64(data["exp"].(float64))
	if poor > configs.Config.TokenExceedRefreshTime {
		service.State.ErrorUnauthorized(ctx, "Leave too long, please log in again")
		return
	}

	token := service.Jwt.Publish(data["info"].(map[string]interface{}))
	service.State.SuccessData(ctx, token)
}
