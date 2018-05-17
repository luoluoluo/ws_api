package boot

import (
	"github.com/gin-gonic/gin"
)

func Run(addr string) error {
	engine := gin.Default()

	engine.Use(func(c *gin.Context) {
		service(c)
	})
	route(engine)
	return engine.Run(addr)
}
