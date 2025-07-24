package scene

type SceneRequest struct {
	Name string `json:"name" binding:"required"`
}
