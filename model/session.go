package model

import "time"

var sessions map[string]*Session = make(map[string]*Session)

type Session struct {
	Id        int    `json:"id"`
	OpenId    string `json:"openid"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Gender    int    `json:"gender"`
	SessionId string `json:"sessionid"`
	Time      int64  `json:"time"`
}

// 获取session
func (s *Session) Get(sessionid string) *Session {
	s.clearSession()
	session, exist := sessions[sessionid]
	if exist {
		return session
	}
	return &Session{}
}

// 设置session
func (s *Session) Set(sessionid string, session *Session) {
	sessions[sessionid] = session
}

// 清除过期session
func (s *Session) clearSession() {
	nowTime := time.Now().Unix()
	for sessionid, session := range sessions {
		if session.Time < nowTime-7200 {
			delete(sessions, sessionid)
		}
	}
}
