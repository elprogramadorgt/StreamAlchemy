package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func defaultResponse(ctx *gin.Context, response interface{}, err error) {
	if err != nil {

		logrus.Error("Error in response: ", err)
		ctx.JSON(500, gin.H{"error": "something went wrong"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
