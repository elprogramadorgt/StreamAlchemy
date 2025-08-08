package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	obsmodel "github.com/elprogramadorgt/StreamAlchemy/models/obs"
	scenemodel "github.com/elprogramadorgt/StreamAlchemy/models/scene"
	"github.com/elprogramadorgt/StreamAlchemy/utils"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func ChangeSceneController(ctx *gin.Context, request *scenemodel.SceneRequest) (scenemodel.SceneResponse, error) {
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

	client, err := utils.NewOBSClient(WS_SERVER)
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

func GetScenesController(ctx *gin.Context) (scenemodel.SceneListResponse, error) {

	client, err := utils.NewOBSClient(WS_SERVER)

	if err != nil {
		logrus.Error("Failed to create OBS client:", err)
		return scenemodel.SceneListResponse{Ok: false, Message: "OBS connection failed"}, err
	}

	defer client.Close()

	request := obsmodel.OBSRequest{
		Op: 6,
		D: obsmodel.RequestOp{
			RequestType: "GetSceneList",
			RequestId:   "req-get-scenes",
		},
	}

	responseBytes, err := client.SendRequest(request)

	if err != nil {
		logrus.Error("Failed to get scenes:", err)
		return scenemodel.SceneListResponse{Ok: false, Message: "Failed to retrieve scenes"}, err
	}

	fmt.Println("OBS Scene List Response:", string(responseBytes))
	response := obsmodel.ObsResponse{}

	if err := json.NewDecoder(bytes.NewBuffer(responseBytes)).Decode(&response); err != nil {
		logrus.Error("Failed to decode OBS response:", err)
		return scenemodel.SceneListResponse{Ok: false, Message: "Failed to decode OBS response"}, err
	}

	return scenemodel.SceneListResponse{
		Ok:     true,
		Scenes: response.D.ResponseData.Scenes,
	}, nil

}

func SetSceneItemVisibilityController(ctx *gin.Context, request *scenemodel.SceneItemRequest) (scenemodel.SceneResponse, error) {
	obsRequest := obsmodel.OBSRequest{
		Op: 6,
		D: obsmodel.RequestOp{
			RequestType: "SetSceneItemEnabled",
			RequestId:   "req-set-scene-item-visibility",
			RequestData: obsmodel.RequestData{
				SceneName:        request.SceneName,
				SceneItemId:      request.SceneItem,
				SceneItemEnabled: false,
			},
		},
	}

	client, err := utils.NewOBSClient(WS_SERVER)
	if err != nil {
		log.Println("OBS client init failed:", err)
		return scenemodel.SceneResponse{Ok: false, Message: "OBS connection failed"}, err
	}
	defer client.Close()

	responseBytes, err := client.SendRequest(obsRequest)
	if err != nil {
		log.Println("Set scene item visibility failed:", err)
		return scenemodel.SceneResponse{Ok: false, Message: "Set scene item visibility failed"}, err
	}

	log.Printf("OBS Set Scene Item Visibility Response: %s", responseBytes)

	return scenemodel.SceneResponse{
		Ok:      true,
		Message: "success",
	}, nil
}
