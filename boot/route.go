package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/luoluoluo/ws_api/handler"
	"github.com/luoluoluo/ws_api/util"
)

func route(r *gin.Engine) {
	userHanlder := &handler.UserHanlder{}
	r.POST("login", userHanlder.Login)

	r.GET("/ping", func(c *gin.Context) {
		db := c.MustGet("db").(*util.DB)
		num, _ := db.Exec("select count(*) from user")
		c.JSON(200, gin.H{
			"message": num,
		})
	})
}
