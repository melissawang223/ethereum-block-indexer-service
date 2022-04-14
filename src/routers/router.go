package routers

import (
	"fmt"
	"github.com/ethereum-block-indexer-service/src/controller"
	"github.com/ethereum-block-indexer-service/src/helpers/config"
	"github.com/gin-gonic/gin"
)

var addr = fmt.Sprintf("%s:%v", config.Get("server.host"), config.Get("server.port"))

// Run router
func Run() {

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	router.GET("/blocks", controller.GetLatestNBlockchain)
	router.GET("/blocks/:id", controller.GetBlockById)
	router.GET("/transaction/:txHash", controller.GetTxWithLogs)

	// default run on :8080
	router.Run(addr)
}
