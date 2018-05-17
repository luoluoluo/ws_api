package config

import "os"

var Wx = map[string]string{
	"id":     os.Getenv("WX_ID"),
	"secret": os.Getenv("WX_SECRET"),
}
