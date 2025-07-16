package handlers

import (
	model_base "github.com/elprogramadorgt/StreamAlchemy/models"
	"github.com/gin-gonic/gin"
)

func GetHealthCheckHandler(ctx *gin.Context) {

	var response = model_base.BaseResponse{
		Error:   false,
		Message: "Server is running",
	}

	defaultResponse(ctx, response, nil)

}
