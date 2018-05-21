package config

import "os"

// App config app
var App = map[string]string{
	"key": os.Getenv("key"),
}
