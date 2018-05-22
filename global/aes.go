package global

import (
	"github.com/luoluoluo/ws_api/config"
	"github.com/luoluoluo/ws_api/library"
)

// AES aes全局变量
var AES = &library.AES{Key: config.App["key"]}
