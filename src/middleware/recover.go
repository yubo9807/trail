package middleware

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// 捕获程序异常，写入日志
func Recover(ctx *gin.Context) {
	defer func() {
		if msg := recover(); msg != nil {
			debug.PrintStack()
			ctx.JSON(200, gin.H{
				"code":    500,
				"data":    nil,
				"message": msg,
			})
			message := "\npanic: " + fmt.Sprintf("%v", msg) + "\n" + string(debug.Stack())
			LogsWrite(ctx, message)
			ctx.Abort()
		}
	}()

	ctx.Next()
}
