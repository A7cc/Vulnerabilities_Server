package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 配置跨域 Cors
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求方式
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, AccessToken")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 如果是OPTIONS则立即停止当前的处理流程并返回特定的 HTTP 状态码
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
