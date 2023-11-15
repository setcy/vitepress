package server

import "github.com/gin-gonic/gin"

func Serve() {
	ginServer := gin.New()

	serveContent(ginServer)

	err := ginServer.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func serveContent(ginServer *gin.Engine) {
	ginServer.GET("/*catchAll", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})
}
