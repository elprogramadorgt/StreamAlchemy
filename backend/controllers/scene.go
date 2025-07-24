package controllers

import (
	"fmt"
	"log"

	obsmodel "github.com/elprogramadorgt/StreamAlchemy/models/obs"
	scenemodel "github.com/elprogramadorgt/StreamAlchemy/models/scene"
	"github.com/elprogramadorgt/StreamAlchemy/utils"

	"github.com/gin-gonic/gin"
)

func ChangeSceneController(ctx *gin.Context, request *scenemodel.SceneRequest) (interface{}, error) {
	obsRequest := obsmodel.OBSRequest{
		Op: 6,
		D: obsmodel.RequestOp{
			RequestType: "SetCurrentProgramScene",
			RequestId:   "req-change-scene",
			RequestData: obsmodel.RequestData{
				SceneName: request.Name,
			},
		},
	}

	client, err := utils.NewOBSClient("ws://192.168.1.2:4455")
	if err != nil {
		log.Println("OBS client init failed:", err)
		return scenemodel.SceneResponse{Ok: false, Message: "OBS connection failed"}, err
	}
	defer client.Close()

	responseBytes, err := client.SendRequest(obsRequest)
	if err != nil {
		log.Println("Set scene failed:", err)
		return scenemodel.SceneResponse{Ok: false, Message: "Scene change failed"}, err
	}

	log.Printf("OBS Scene Change Response: %s", responseBytes)

	return scenemodel.SceneResponse{
		Ok:      true,
		Message: fmt.Sprintf("Scene changed to %s", request.Name),
	}, nil
}
