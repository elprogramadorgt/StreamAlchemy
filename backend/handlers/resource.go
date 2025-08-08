package handlers

import (
	"github.com/elprogramadorgt/StreamAlchemy/controllers"
	"github.com/gin-gonic/gin"
)

func GetResourceHandler(ctx *gin.Context) {

	response, err := controllers.GetResourcesController(ctx)

	defaultResponse(ctx, response, err)

}
