package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/luoluoluo/ws_api/config"
	"github.com/luoluoluo/ws_api/library"
)

type Session struct {
	Id        int    `json:"id"`
	OpenId    string `json:"openid"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Gender    int    `json:"gender"`
	SessionId string `json:"sessionid"`
	Time      int64  `json:"time"`
}
type WxSession struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
}
type registerReq struct {
	Code   string `json:"code"`
	Avatar string `json:"avatar"`
	Gender int    `json:"gender"`
	Name   string `json:"name"`
}
type UserController struct {
}

var sessions map[string]*Session = make(map[string]*Session)

// 注册
func (h *UserController) Login(c *gin.Context) {
	go clearSession()
	var req registerReq
	nowTime := time.Now().Unix()
	err := c.BindJSON(&req)
	if err != nil {
		glog.Error(err)
		c.JSON(400, gin.H{})
		return
	}
	if req.Code == "" {
		c.JSON(400, gin.H{})
		return
	}
	session, err := wxJscodeToSession(req.Code)
	if err != nil {
		glog.Error(err)
		c.JSON(500, gin.H{})
		return
	}
	if session.OpenId == "" || session.SessionKey == "" {
		c.JSON(500, gin.H{})
		return
	}
	db := c.MustGet("db").(*library.DB)
	user, err := db.SelectOne("SELECT * FROM user WHERE openid=?", session.OpenId)
	if err != nil {
		glog.Error(err)
		c.JSON(500, gin.H{})
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
			c.JSON(500, gin.H{})
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
			c.JSON(500, gin.H{})
			return
		}
	}

	user, err = db.SelectOne("SELECT * FROM user WHERE openid=?", session.OpenId)

	if err != nil || len(user) == 0 {
		glog.Error(err)
		c.JSON(500, gin.H{})
		return
	}

	sessionid := library.Md5(session.SessionKey)

	res := &Session{
		library.ParseInt(user["id"]),
		user["openId"],
		user["name"],
		user["avatar"],
		library.ParseInt(user["gender"]),
		sessionid,
		nowTime,
	}

	setSession(sessionid, res)

	c.JSON(200, res)
	return
}

// js code 换取 session
func wxJscodeToSession(code string) (*WxSession, error) {
	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + config.Wx["id"] + "&secret=" + config.Wx["secret"] + "&js_code=" + code + "&grant_type=authorization_code")

	if err != nil {
		glog.Error(err)
		return &WxSession{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	glog.Infoln("wxresp:" + string(body))
	if err != nil {
		glog.Error(err)
		return &WxSession{}, err
	}
	wxSession := &WxSession{}
	json.Unmarshal(body, wxSession)
	if wxSession.OpenId == "" || wxSession.SessionKey == "" {
		glog.Error("wsres:" + string(body))
		return &WxSession{}, err
	}
	return wxSession, nil
}

// 获取session
func GetSession(sessionid string) *Session {
	session, exist := sessions[sessionid]
	if exist {
		return session
	}
	return &Session{}
}

// 设置session
func setSession(sessionid string, session *Session) {
	sessions[sessionid] = session
}

// 清除过期session
func clearSession() {
	nowTime := time.Now().Unix()
	for sessionid, session := range sessions {
		if session.Time < nowTime-7200 {
			delete(sessions, sessionid)
		}
	}
}
