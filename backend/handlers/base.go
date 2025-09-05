package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func defaultResponse(ctx *gin.Context, response interface{}, err error) {
	if err != nil {
		logrus.Error("Error in response: ", err)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
