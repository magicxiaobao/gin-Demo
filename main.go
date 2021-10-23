package main

import (
	"github.com/gin-gonic/gin"
	"goDemo/config"
	"goDemo/middleware"
	"goDemo/route"
)

func main() {
	r := gin.Default()
	//r := gin.New()
	r.Use(middleware.LoggerToFile())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong5",
		})
	})
	route.InitRouter(r)
	r.Run(config.SERVER_PORT) // listen and serve on 0.0.0.0:8080
}
