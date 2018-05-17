package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/luoluoluo/ws_api/handler"
	"github.com/luoluoluo/ws_api/library"
)
var userHanlder handler.UserHanlder = &handler.UserHanlder{}

func route(r *gin.Engine) {
	r.POST("login", userHanlder.Login)
}

func func checkLogin()() gin.HandlerFunc {
    return func(c *gin.Context) {
		sessionid := c.GetHeader("sessionid")
		session := userHanlder.GetSession(sessionid)
		if session.Id == "" {
			c.JSON
			return
		}
        c.Set("request", "clinet_request")
        c.Next()
        fmt.Println("before middleware")
    }
}
