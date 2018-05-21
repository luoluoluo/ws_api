package controller

import (
	"time"

	"github.com/jmoiron/sqlx"
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
	db := c.MustGet("db").(*sqlx.DB)
	userModel := &model.User{}
	user, err := userModel.Insert(db, req.Code, req.Avatar, req.Name, req.Gender)
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
	})
	return
}
