package middleware

import (
	"server/src/service"

	"github.com/gin-gonic/gin"
)

const KEY = "user_info"

func Authorization(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		service.State.ErrorUnauthorized(ctx, "Unauthorized")
		ctx.Abort()
		return
	}

	info, err := service.Jwt.Verify(auth)
	if err != nil {
		if err.Error() == "Token is expired" {
			// 可尝试刷新 token
			service.State.ErrorTokenFailure(ctx)
		} else {
			service.State.ErrorUnauthorized(ctx, err.Error())
		}
		ctx.Abort()
		return
	}
	ctx.Set(KEY, info)
}

// 获取 token 储存信息
func GetTokenInfo(ctx *gin.Context) map[string]interface{} {
	info, _ := ctx.Get(KEY)
	return info.(map[string]interface{})
}

// 角色校验
func RoleVerify(roleId string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		info := GetTokenInfo(ctx)
		if info["roleId"] != roleId {
			service.State.ErrorCustom(ctx, "The current user has no permission")
			return
		}
	}
}
