package obs

type IdentifyOp struct {
	RpcVersion int `json:"rpcVersion"`
}

type AuthIdentifyOp struct {
	RpcVersion int `json:"rpcVersion"`
}

type RequestData struct {
	SceneName string `json:"sceneName"`
}

type RequestOp struct {
	RequestType string      `json:"requestType"`
	RequestId   string      `json:"requestId"`
	RequestData RequestData `json:"requestData,omitempty"`
}

type OBSRequest struct {
	Op int       `json:"op"`
	D  RequestOp `json:"d"`
}

type OBSIdentify struct {
	Op int        `json:"op"`
	D  IdentifyOp `json:"d"`
}
