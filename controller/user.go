package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/luoluoluo/ws_api/library"
	"github.com/luoluoluo/ws_api/model"
)

type loginReq struct {
	Code   string `json:"code"`
	Avatar string `json:"avatar"`
	Gender int    `json:"gender"`
	Name   string `json:"name"`
}
type UserController struct {
	Controller
}

// 登录
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
	session, err := library.WxJscodeToSession(req.Code)
	if err != nil {
		glog.Error(err)
		u.resp(c, 1000, nil)
		return
	}
	if session.OpenId == "" || session.SessionKey == "" {
		u.resp(c, 1000, nil)
		return
	}
	db := c.MustGet("db").(*library.DB)
	user, err := db.SelectOne("SELECT * FROM user WHERE openid=?", session.OpenId)
	if err != nil {
		glog.Error(err)
		u.resp(c, 500, nil)
		return
	}
	// 新增用户信息
	if len(user) == 0 {
		_, err = db.Insert(
			"INSERT INTO user(openid,name,avatar,gender,update_time,create_time) VALUES(?,?,?,?,?,?)",
			session.OpenId,
			req.Name,
			req.Avatar,
			req.Gender,
			nowTime,
			nowTime,
		)
		if err != nil {
			glog.Error(err)
			u.resp(c, 500, nil)
			return
		}
	} else { // 修改用户信息
		_, err = db.Insert(
			"UPDATE user SET name=?,avatar=?,gender=?,update_time=? WHERE id=?",
			req.Name,
			req.Avatar,
			req.Gender,
			nowTime,
			user["id"],
		)
		if err != nil {
			glog.Error(err)
			u.resp(c, 500, nil)
			return
		}
	}

	user, err = db.SelectOne("SELECT * FROM user WHERE openid=?", session.OpenId)

	if err != nil || len(user) == 0 {
		glog.Error(err)
		u.resp(c, 500, nil)
		return
	}

	sessionid := library.Md5(session.SessionKey)

	res := &model.Session{
		library.ParseInt(user["id"]),
		user["openId"],
		user["name"],
		user["avatar"],
		library.ParseInt(user["gender"]),
		sessionid,
		nowTime,
	}
	s := &model.Session{}
	s.Set(sessionid, res)

	u.resp(c, 200, res)
	return
}
