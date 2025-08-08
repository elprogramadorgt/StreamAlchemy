package handlers

import (
	"github.com/elprogramadorgt/StreamAlchemy/controllers"
	model_request "github.com/elprogramadorgt/StreamAlchemy/models/scene"
	"github.com/gin-gonic/gin"
)

func ChangeSceneHandler(ctx *gin.Context) {
	var request model_request.SceneRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {

		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := controllers.ChangeSceneController(ctx, &request)

	defaultResponse(ctx, response, err)
}

func GetScenesHandler(ctx *gin.Context) {

	response, err := controllers.GetScenesController(ctx)

	defaultResponse(ctx, response, err)

}

func SetSceneItemVisibilityHandler(ctx *gin.Context) {
	var request model_request.SceneItemRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := controllers.SetSceneItemVisibilityController(ctx, &request)

	defaultResponse(ctx, response, err)
}
