package api

import "github.com/gin-gonic/gin"

func ServeAPI(ginServer *gin.Engine) {
	// No authentication required
	ginServer.GET("/_meta", getMeta)
	ginServer.GET("/_content/*catchAll", getContent)
}
