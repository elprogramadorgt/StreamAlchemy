package scene

type SceneRequest struct {
	Name string `json:"name" binding:"required"`
}

type SceneItemRequest struct {
	SceneName string `json:"sceneName" binding:"required"`
	SceneItem int    `json:"sceneItem" binding:"required"`
	// Visible   bool   `json:"visible" binding:"required"`
}
