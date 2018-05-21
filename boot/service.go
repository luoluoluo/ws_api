package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/luoluoluo/ws_api/config"
	"github.com/luoluoluo/ws_api/library"
)

func service(c *gin.Context) {
	c.Set("db", getDb())
	c.Set("aes", getAes())
}

func getDb() *sqlx.DB {
	db, err := sqlx.Connect("mysql", config.DB["user"]+":"+config.DB["password"]+"@tcp("+config.DB["host"]+":"+config.DB["port"]+")/"+config.DB["dbname"])
	if err != nil {
		panic(err)
	}
	return db
}

func getAes() *library.AES {
	aes := &library.AES{
		Key: []byte(config.App["key"]),
	}
	return aes
}
