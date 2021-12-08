package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magicxiaobao/ginDemo/config"
	"github.com/magicxiaobao/ginDemo/middleware/logger"
	"github.com/magicxiaobao/ginDemo/middleware/recover"
	"github.com/magicxiaobao/ginDemo/route"
)

func main() {
	r := gin.Default()
	//r := gin.New()
	r.Use(logger.LoggerToFile(), recover.Recover())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong5",
		})
	})
	route.InitRouter(r)
	r.Run(config.SERVER_PORT) // listen and serve on 0.0.0.0:8080
}
