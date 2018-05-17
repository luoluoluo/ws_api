package main

import (
	"flag"

	"os"

	"github.com/luoluoluo/ws_api/boot"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	flag.Parse()
	boot.Run(host + ":" + port)
}
