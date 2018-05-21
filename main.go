package main

import (
	"flag"

	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/luoluoluo/ws_api/boot"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	flag.Parse()
	boot.Run(host + ":" + port)
}
