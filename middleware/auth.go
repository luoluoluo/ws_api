package middleware

import (
	"github.com/gin-gonic/gin"
)

// Auth 校验登录
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user", "")
	}
}
