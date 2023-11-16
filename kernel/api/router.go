package api

import "github.com/gin-gonic/gin"

func ServeAPI(ginServer *gin.Engine) {
	// No authentication required
	ginServer.GET("/meta", getMeta)
	ginServer.GET("/content/*catchAll", getContent)
}
