package controller

import (
	"time"

	"github.com/luoluoluo/ws_api/config"
	"github.com/luoluoluo/ws_api/library"
	"github.com/luoluoluo/ws_api/model"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type loginReq struct {
	Code   string `json:"code"`
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Gender int    `json:"gender"`
}

// UserController user控制器
type UserController struct {
	Controller
}

// Login 登录
func (u *UserController) Login(c *gin.Context) {
	var req loginReq
	nowTime := time.Now().Unix()
	err := c.BindJSON(&req)
	if err != nil {
		glog.Error(err)
		u.resp(c, 400, nil)
		return
	}
	if req.Code == "" {
		u.resp(c, 400, nil)
		return
	}
	wx := &library.WX{
		ID:     config.WX["id"],
		Secret: config.WX["secret"],
	}
	wxSession, err := wx.JSCodeToSession(req.Code)
	if err != nil {
		u.resp(c, 500, nil)
		return
	}
	user := &model.User{
		OpenID:     wxSession.OpenID,
		Name:       req.Name,
		Avatar:     req.Avatar,
		Gender:     req.Gender,
		SessionKey: wxSession.SessionKey,
	}
	user, err = user.Insert()
	if err != nil {
		glog.Error(err)
		u.resp(c, 500, nil)
		return
	}
	token := &model.Token{
		OpenID:     user.OpenID,
		SessionKey: wxSession.SessionKey,
		ExpireTime: nowTime + 7200,
	}
	tokenStr, err := token.Encrypt()
	if err != nil {
		glog.Error(err)
		u.resp(c, 500, nil)
		return
	}
	u.resp(c, 200, gin.H{
		"id":     user.ID,
		"openid": user.OpenID,
		"name":   user.Name,
		"avatar": user.Avatar,
		"gender": user.Gender,
		"time":   nowTime,
		"token":  tokenStr,
	})
	return
}
