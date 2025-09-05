package models

type BaseResponse struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type BaseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e BaseError) Error() string {
	return e.Message
}
