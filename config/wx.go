package config

import "os"

// WX wx配置
var WX = map[string]string{
	"id":     os.Getenv("WX_ID"),
	"secret": os.Getenv("WX_SECRET"),
}
