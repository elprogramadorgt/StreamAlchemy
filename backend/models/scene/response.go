package scene

type SceneResponse struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
