package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/luoluoluo/ws_api/controller"
)

var userController *controller.UserController = &controller.UserController{}

func route(r *gin.Engine) {
	r.POST("login", userController.Login)
}
