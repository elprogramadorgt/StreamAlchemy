package scene

import (
	obs_model "github.com/elprogramadorgt/StreamAlchemy/models/obs"
)

type SceneResponse struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type SceneListResponse struct {
	Ok      bool              `json:"ok"`
	Scenes  []obs_model.Scene `json:"scenes"`
	Message string            `json:"message,omitempty"`
	Error   string            `json:"error,omitempty"`
}
