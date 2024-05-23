package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Core(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS, UPDATE")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}
}
