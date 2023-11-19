package server

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/setcy/wiki/kernel/api"
)

var Mode string

func Serve() {
	ginServer := gin.New()

	if Mode == "allInOne" {
		serveStatic(ginServer)
	} else {
		ginServer.Use(cors())
	}

	api.ServeAPI(ginServer)

	err := ginServer.Run(":3026")
	if err != nil {
		panic(err)
	}
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func serveStatic(ginServer *gin.Engine) {
	// 正常处理静态文件
	ginServer.Static("/assets", "./static/assets")
	ginServer.StaticFile("/favicon.ico", "./static/favicon.ico")

	// 处理前端 history 模式
	ginServer.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})
}
