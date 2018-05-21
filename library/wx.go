package library

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"

	"github.com/golang/glog"
)

// WXSession 微信 session_key
type WXSession struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

// WXToken 微信 access_token
type WXToken struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

// WXUser 微信 userinfo
type WXUser struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	OpenID  string `json:"openid"`
	Name    string `json:"nickname"`
	Gender  string `json:"sex"`
	Avatar  string `json:"headimgurl"`
}

// WX 微信
type WX struct {
	ID     string
	Secret string
}

// JSCodeToSession js code 换取 session
func (wx *WX) JSCodeToSession(code string) (*WXSession, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	h := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	resp, err := h.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + wx.ID + "&secret=" + wx.Secret + "&js_code=" + code + "&grant_type=authorization_code")

	if err != nil {
		glog.Error(err)
		return &WXSession{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	glog.Infoln("wxresp:" + string(body))
	if err != nil {
		glog.Error(err)
		return &WXSession{}, err
	}
	wxSession := &WXSession{}
	json.Unmarshal(body, wxSession)
	if wxSession.ErrCode != 0 {
		glog.Error("wsres:" + string(body))
		return &WXSession{}, errors.New(wxSession.ErrMsg)
	}
	return wxSession, nil
}

// Token 获取access token
func (wx *WX) Token() (*WXToken, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	h := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	resp, err := h.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + wx.ID + "&secret=" + wx.Secret)

	if err != nil {
		glog.Error(err)
		return &WXToken{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	glog.Infoln("wxresp:" + string(body))
	if err != nil {
		glog.Error(err)
		return &WXToken{}, err
	}
	wxToken := &WXToken{}
	json.Unmarshal(body, wxToken)
	if wxToken.ErrCode != 0 {
		glog.Error("wsres:" + string(body))
		return &WXToken{}, errors.New(wxToken.ErrMsg)
	}
	return wxToken, nil
}

// User 获取user info
func (wx *WX) User(accessToken, openID string) (*WXUser, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	h := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	resp, err := h.Get("https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + accessToken + "&openid=" + openID + "&lang=zh_CN")

	if err != nil {
		glog.Error(err)
		return &WXUser{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	glog.Infoln("wxresp:" + string(body))
	if err != nil {
		glog.Error(err)
		return &WXUser{}, err
	}
	wxUser := &WXUser{}
	json.Unmarshal(body, wxUser)
	if wxUser.ErrCode != 0 {
		glog.Error("wsres:" + string(body))
		return &WXUser{}, errors.New(wxUser.ErrMsg)
	}
	return wxUser, nil
}
