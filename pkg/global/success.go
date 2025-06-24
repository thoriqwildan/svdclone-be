package global

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    any        `json:"data,omitempty"`
}