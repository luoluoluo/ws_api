package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luoluoluo/ws_api/config"
)

func resp(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"code": code,
		"text": config.StatusText(code),
		"data": data,
	})
}
