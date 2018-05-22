package model

import (
	"encoding/json"

	"github.com/luoluoluo/ws_api/global"
)

// Token token model
type Token struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	ExpireTime int64  `json:"expire_time"`
}

// Encrypt 生成token
func (t *Token) Encrypt() (string, error) {
	token, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	text, err := global.AES.NewCBCEncrypter(string(token))
	if err != nil {
		return "", err
	}
	return text, nil
}

// Decrypt 解析token
func (t *Token) Decrypt(token string) (*Token, error) {
	text, err := global.AES.NewCBCDecrypter(token)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(text), t)
	if err != nil {
		return nil, err
	}
	return t, nil
}
