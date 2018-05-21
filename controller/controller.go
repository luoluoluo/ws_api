package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luoluoluo/ws_api/config"
)

// Controller 控制器基类
type Controller struct {
}

func (controller *Controller) resp(context *gin.Context, code int, data interface{}) {
	context.JSON(code, gin.H{
		"status_code": code,
		"status_text": config.StatusText(code),
		"data":        data,
	})
}
