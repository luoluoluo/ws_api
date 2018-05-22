package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/luoluoluo/ws_api/model"
)

// Auth 校验登录
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := &model.Token{}
		authStr := c.Request.Header.Get("Authorization")
		authSlice := strings.Split(authStr, " ")
		if len(authSlice) != 2 || authSlice[0] != "Bearer" {
			glog.Error("token不能为空")
			c.AbortWithStatus(401)
			return
		}
		token, err := token.Decrypt(authSlice[1])
		if err != nil {
			glog.Error(err)
			c.AbortWithStatus(401)
			return
		}
		user := &model.User{
			OpenID: token.OpenID,
		}
		user, err = user.GetByOpenID()
		if err != nil {
			glog.Error(err)
			c.AbortWithStatus(401)
			return
		}
		if user.SessionKey != token.SessionKey {
			glog.Error("session_key已过期")
			c.AbortWithStatus(401)
			return
		}
		if token.ExpireTime < time.Now().Unix() {
			glog.Error("token已过期")
			c.AbortWithStatus(401)
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
