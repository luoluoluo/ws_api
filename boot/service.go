package boot

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/luoluoluo/ws_api/config"
	"github.com/luoluoluo/ws_api/util"
)

func service(c *gin.Context) {
	c.Set("db", getDb())
}

func getDb() *util.DB {
	db, err := util.NewDB("mysql", config.Db["user"]+":"+config.Db["password"]+"@tcp("+config.Db["host"]+":"+config.Db["port"]+")/"+config.Db["dbname"])
	if err != nil {
		panic(err)
	}
	return db
}
