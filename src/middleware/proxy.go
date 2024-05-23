package middleware

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// 权限应用代理
func ProxyPermissions(ctx *gin.Context) {
	targetURL, _ := url.Parse("http://localhost:20020/")
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	ctx.Request.URL.Scheme = targetURL.Scheme
	ctx.Request.URL.Host = targetURL.Host
	ctx.Request.Host = targetURL.Host
	ctx.Request.Header.Set("Open-Id", "1hendj97f")

	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
