package library

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"

	"github.com/golang/glog"
	"github.com/luoluoluo/ws_api/config"
)

type WxSession struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

// js code 换取 session
func WxJscodeToSession(code string) (*WxSession, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	h := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	resp, err := h.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + config.Wx["id"] + "&secret=" + config.Wx["secret"] + "&js_code=" + code + "&grant_type=authorization_code")

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
