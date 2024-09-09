package middleware

import (
	"Go_server/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的Authorization
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "未登录系统",
			})
			// 终止访问
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "当前登录已失效请重新登录",
			})
			c.Abort()
			return
		}
		// 验证token
		tokenClaims, err := helper.ValidateToken(parts[1])
		if tokenClaims == nil || err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "当前登录已失效请重新登录",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
