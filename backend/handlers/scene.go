package handlers

import (
	"fmt"

	"github.com/elprogramadorgt/StreamAlchemy/controllers"
	model_request "github.com/elprogramadorgt/StreamAlchemy/models/scene"
	"github.com/gin-gonic/gin"
)

func ChangeSceneHandler(ctx *gin.Context) {
	var request model_request.SceneRequest
	fmt.Println("ChangeSceneHandler called")
	if err := ctx.ShouldBindJSON(&request); err != nil {

		fmt.Println("Error binding JSON:", err)

		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := controllers.ChangeSceneController(ctx, &request)

	defaultResponse(ctx, response, err)
}
