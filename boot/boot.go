package boot

import (
	"github.com/gin-gonic/gin"
)

// Run 运行
func Run(addr string) error {
	engine := gin.Default()
	route(engine)
	return engine.Run(addr)
}
