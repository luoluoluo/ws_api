package model

import (
	"time"

	"github.com/golang/glog"
	"github.com/luoluoluo/ws_api/global"
)

// User user model
type User struct {
	ID         int    `json:"id" db:"id"`
	OpenID     string `json:"openid" db:"openid"`
	Name       string `json:"name" db:"name"`
	Avatar     string `json:"avatar" db:"avatar"`
	Gender     int    `json:"gender" db:"gender"`
	SessionKey string `json:"session_key" db:"session_key"`
	UpdateTime int    `json:"update_time" db:"update_time"`
	CreateTime int    `json:"create_time" db:"create_time"`
}

// Insert 新增用户
func (u *User) Insert() (*User, error) {
	nowTime := time.Now().Unix()

	var count int
	err := global.DB.Get(&count, "SELECT count(*) FROM user WHERE openid=?", u.OpenID)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	// 新增用户信息
	if count == 0 {
		_, err := global.DB.Exec(
			`INSERT INTO user(openid,name,avatar,gender,session_key, update_time,create_time) 
			VALUES(?,?,?,?,?,?,?)`,
			u.OpenID,
			u.Name,
			u.Avatar,
			u.Gender,
			u.SessionKey,
			nowTime,
			nowTime,
		)
		if err != nil {
			glog.Error(err)
			return nil, err
		}
	} else { // 修改用户信息
		_, err := global.DB.Exec(
			"UPDATE user SET name=?,avatar=?,gender=?,session_key=?,update_time=? WHERE openid=?",
			u.Name,
			u.Avatar,
			u.Gender,
			u.SessionKey,
			nowTime,
			u.OpenID,
		)
		if err != nil {
			glog.Error(err)
			return nil, err
		}
	}
	return u.GetByOpenID()
}

// GetByOpenID 根据openid获取用户信息
func (u *User) GetByOpenID() (*User, error) {
	err := global.DB.Get(u, "SELECT * FROM user WHERE openid=?", u.OpenID)
	return u, err
}
