package middleware

import (
	"Go_server/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的Authorization
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			helper.ErrorResponse(c, "未登录系统", nil)
			// 终止访问
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			helper.ErrorResponse(c, "当前登录已失效请重新登录", nil)
			c.Abort()
			return
		}
		// 验证token
		tokenClaims, err := helper.ValidateToken(parts[1])
		if tokenClaims == nil || err != nil {
			helper.ErrorResponse(c, "当前登录已失效请重新登录", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
