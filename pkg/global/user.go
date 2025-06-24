package global

type UserResponse struct {
	Id int	`json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	ProfileUrl string `json:"profile_url,omitempty"`
	Admin bool `json:"admin"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}