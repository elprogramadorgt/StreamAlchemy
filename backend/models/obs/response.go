package obs

// TODO: Add the authentication field
// right now it assumes that the websocket plugins doesnt have authentication
type ConnectionResponse struct {
	D  D   `json:"d"`
	Op int `json:"op"`
}

type ObsResponse struct {
	D  D   `json:"d"`
	Op int `json:"op"`
}
