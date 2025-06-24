package global

type PaginationData struct {
	Items any				`json:"items"`
	Meta PaginationPage `json:"meta"`
}

type PaginationPage struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	Total int64 `json:"total"`
}