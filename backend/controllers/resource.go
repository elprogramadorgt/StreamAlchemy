package controllers

import (
	obsmodel "github.com/elprogramadorgt/StreamAlchemy/models/obs"
	"github.com/elprogramadorgt/StreamAlchemy/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetResourcesController(ctx *gin.Context) (interface{}, error) {
	// This function will handle the logic to retrieve resources
	// For now, we return an empty response

	client, err := utils.NewOBSClient(WS_SERVER)

	obsRequest := obsmodel.OBSRequest{
		Op: 6,
		D: obsmodel.RequestOp{
			RequestType: "GetSourceActive",
			RequestId:   "req-change-GetSourceActive",

			RequestData: obsmodel.RequestData{
				SourceName: "taskmanager",
			},
		},
	}

	if err != nil {
		logrus.Error("Failed to create OBS client:", err)
		return nil, err
	}
	defer client.Close()

	responseBytes, err := client.SendRequest(obsRequest)
	if err != nil {
		logrus.Error("Failed to send OBS request:", err)
		return nil, err
	}

	logrus.Info("OBS Response:", string(responseBytes))

	return responseBytes, nil
}
