package global

import (
	"github.com/jmoiron/sqlx"
	"github.com/luoluoluo/ws_api/config"
)

// DB db全局变量
var DB = getDB()

func getDB() *sqlx.DB {
	db, err := sqlx.Connect("mysql", config.DB["user"]+":"+config.DB["password"]+"@tcp("+config.DB["host"]+":"+config.DB["port"]+")/"+config.DB["dbname"])
	if err != nil {
		panic(err)
	}
	return db
}
