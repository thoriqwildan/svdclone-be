package authorization

import (
	"github.com/thoriqwildan/svdclone-be/pkg/database"
	"github.com/thoriqwildan/svdclone-be/pkg/database/models"
	"github.com/thoriqwildan/svdclone-be/pkg/global"
)

func GetEmail(email string) global.UserResponse {
	var user models.User

	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return global.UserResponse{}
	}

	return global.UserResponse{
		Id: int(user.ID),
		Name: user.Name,
		Email: user.Email,
		ProfileUrl: user.ProfileUrl.String,
		Admin: user.Admin,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}