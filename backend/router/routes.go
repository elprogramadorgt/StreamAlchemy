package router

import (
	"github.com/elprogramadorgt/StreamAlchemy/handlers"
	"github.com/elprogramadorgt/StreamAlchemy/middleware"
	"github.com/gin-gonic/gin"
)

func GinRoutes(engine *gin.Engine) {
	engine.Use(middleware.ErrorMiddleware())
	health := engine.Group("/healthcheck")
	{
		health.GET("", handlers.GetHealthCheckHandler)
	}

	scenes := engine.Group("/scene")
	{
		scenes.GET("", handlers.GetScenesHandler)
		scenes.POST("", handlers.ChangeSceneHandler)
		scenes.POST("/activate-item", handlers.SetSceneItemVisibilityHandler)
	}

	resource := engine.Group("/resource")
	{
		resource.GET("", handlers.GetResourceHandler)
	}
}
