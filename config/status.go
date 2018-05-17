package config

import "net/http"

// 自定义错误码 从1000开始
var statusText = map[int]string{
	1000: "微信信息获取失败",
}

func StatusText(code int) string {
	text := http.StatusText(code)
	if text != "" {
		return text
	}
	text, ok := statusText[code]
	if ok {
		return text
	}
	return http.StatusText(500)
}
