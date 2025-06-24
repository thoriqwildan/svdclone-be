package global

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Errors    any    `json:"errors,omitempty"`
}