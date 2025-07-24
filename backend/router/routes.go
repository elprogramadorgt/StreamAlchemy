package router

import (
	"github.com/elprogramadorgt/StreamAlchemy/handlers"
	"github.com/gin-gonic/gin"
)

func GinRoutes(engine *gin.Engine) {

	health := engine.Group("/healthcheck")
	{
		health.GET("", handlers.GetHealthCheckHandler)
	}

	scenes := engine.Group("/scene")
	{
		scenes.POST("", handlers.ChangeSceneHandler)
	}

}
