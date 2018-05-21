package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/luoluoluo/ws_api/controller"
	"github.com/luoluoluo/ws_api/middleware"
)

var userController = &controller.UserController{}
var taskController = &controller.TaskController{}

func route(r *gin.Engine) {
	r.POST("login", userController.Login)
	r.GET("task", middleware.Auth(), taskController.List)
	r.GET("task/{id}", middleware.Auth(), taskController.Info)
}
