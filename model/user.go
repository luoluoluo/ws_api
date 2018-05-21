package model

import (
	"time"

	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"github.com/luoluoluo/ws_api/config"
	"github.com/luoluoluo/ws_api/library"
)

// User user model
type User struct {
	ID          int    `json:"id" db:"id"`
	OpenID      string `json:"openid" db:"openid"`
	Name        string `json:"name" db:"name"`
	Avatar      string `json:"avatar" db:"avatar"`
	Gender      int    `json:"gender" db:"gender"`
	SessionKey  string `json:"session_key" db:"session_key"`
	AccessToken string `json:"access_token" db:"access_token"`
	UpdateTime  int    `json:"update_time" db:"update_time"`
	CreateTime  int    `json:"create_time" db:"create_time"`
}

// Insert 新增用户
func (u *User) Insert(db *sqlx.DB, code string, avatar string, name string, gender int) (*User, error) {
	nowTime := time.Now().Unix()
	wx := &library.WX{
		ID:     config.WX["id"],
		Secret: config.WX["secret"],
	}
	wxSession, err := wx.JSCodeToSession(code)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	var count int
	err = db.Get(&count, "SELECT count(*) FROM user WHERE openid=?", wxSession.OpenID)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	// 新增用户信息
	if count == 0 {
		_, err := db.Exec(
			`INSERT INTO user(openid,name,avatar,gender,session_key, update_time,create_time) 
			VALUES(?,?,?,?,?,?,?)`,
			wxSession.OpenID,
			name,
			avatar,
			gender,
			wxSession.SessionKey,
			nowTime,
			nowTime,
		)
		if err != nil {
			glog.Error(err)
			return nil, err
		}
	} else { // 修改用户信息
		_, err := db.Exec(
			"UPDATE user SET name=?,avatar=?,gender=?,session_key=?,update_time=? WHERE openid=?",
			name,
			avatar,
			gender,
			wxSession.SessionKey,
			nowTime,
			wxSession.OpenID,
		)
		if err != nil {
			glog.Error(err)
			return nil, err
		}
	}
	return u.GetByOpenID(db, wxSession.OpenID)
}

// GetByOpenID 根据openid获取用户信息
func (u *User) GetByOpenID(db *sqlx.DB, openID string) (*User, error) {
	user := &User{}
	err := db.Get(user, "SELECT * FROM user WHERE openid=?", openID)
	return user, err
}
