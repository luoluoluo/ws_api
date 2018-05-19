package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luoluoluo/ws_api/config"
)

type Controller struct {
}

func (controller *Controller) resp(context *gin.Context, code int, data interface{}) {
	context.JSON(code, gin.H{
		"code": code,
		"text": config.StatusText(code),
		"data": data,
	})
}
