package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/luoluoluo/ws_api/controller"
	"github.com/luoluoluo/ws_api/middleware"
)

var userController = &controller.UserController{}
var taskController = &controller.TaskController{}
var commentController = &controller.CommentController{}

func route(r *gin.Engine) {
	r.POST("login", userController.Login)
	r.GET("timeline", middleware.Auth(), taskController.Timeline)
	r.POST("task", middleware.Auth(), taskController.Add)
	r.POST("comment", middleware.Auth(), commentController.Add)
}
