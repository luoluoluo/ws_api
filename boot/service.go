package boot

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/luoluoluo/ws_api/util"
)

func service(c *gin.Context) {
	c.Set("db", getDb())
}

func getDb() *util.DB {
	user := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	db, err := util.NewDB("mysql", user+":"+passwd+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		panic(err)
	}
	return db
}
