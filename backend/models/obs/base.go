package obs

type Scene struct {
	SceneIndex int    `json:"sceneIndex" binding:"required,min=0,max=1000"`
	SceneName  string `json:"sceneName" binding:"required"`
	SceneUuid  string `json:"sceneUuid" binding:"required,uuid"`
}

type EventData struct {
	TransitionName string `json:"transitionName"`
	TransitionUuid string `json:"transitionUuid"`
}

type ResponseData struct {
	Scenes []Scene `json:"scenes"`
}

type D struct {
	ObsWebSocketVersion string       `json:"obsWebSocketVersion"`
	RpcVersion          int          `json:"rpcVersion"`
	EventData           EventData    `json:"eventData"`
	EventIntent         int          `json:"eventIntent"`
	EventType           string       `json:"eventType"`
	RequestType         string       `json:"requestType"`
	ResponseData        ResponseData `json:"responseData"`
}

type OBSPayload interface {
	ObsResponse | ConnectionResponse
}
