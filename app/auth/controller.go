package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/thoriqwildan/svdclone-be/pkg/database"
	"github.com/thoriqwildan/svdclone-be/pkg/database/models"
	"github.com/thoriqwildan/svdclone-be/pkg/global"
	"github.com/thoriqwildan/svdclone-be/pkg/helper"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var req LoginRequest
	var user models.User

	c.BodyParser(&req)

	if err := helper.Validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Validation failed",
			Errors:  helper.TranslateErrorMessage(err),
		})
	}

	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Invalid credentials",
			Errors: "Email not found! Please register first.",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Invalid credentials",
			Errors:  "Incorrect password",
		})
	}



	// Simulate a successful login
	token, err := helper.GenerateToken(req.Email, false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Internal server error",
			Errors:  "",
		})
	}

	return c.Status(fiber.StatusOK).JSON(AuthResponse{
			AccessToken: token,
			UserData: global.UserResponse{
				Id: 			 int(user.ID),
				Name: 		 user.Name,
				Email: 		 user.Email,
				ProfileUrl: user.ProfileUrl.String,
				Admin: 		 user.Admin,
				CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
			UserAbilityRules: []AbilityRule{
				AbilityRule{Action: "manage", Subject: "all"},
			},
		},
	)
}

func Register(c *fiber.Ctx) error {
	var req RegisterRequest
	var user models.User

	c.BodyParser(&req)
	if err := helper.Validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Validation failed",
			Errors:  helper.TranslateErrorMessage(err),
		})
	}

	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "Email already exists",
			Errors:  "Email already registered",
		})
	}

	newPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		log.Error("Failed to hash password:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Internal server error",
			Errors:  "Failed to hash password",
		})
	}
	user = models.User{
		Name: req.Name,
		Email: req.Email,
		Password: newPassword,
		ProfileUrl: helper.ToNullString(req.ProfileUrl),
	}
	if err := database.DB.Create(&user).Error; err != nil {
		log.Error("Failed to create user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Internal server error",
			Errors:  "Failed to create user",
		})
	}
	token, err := helper.GenerateToken(user.Email, user.Admin)
	if err != nil {
		log.Error("Failed to generate token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(global.ErrorResponse{
			Success: false,
			Message: "Internal server error",
			Errors:  "Failed to generate token",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(AuthResponse{
			AccessToken: token,
			UserData: global.UserResponse{
				Id: 			 int(user.ID),
				Name: 		 user.Name,
				Email: 		 user.Email,
				ProfileUrl: user.ProfileUrl.String,
				Admin: 		 user.Admin,
				CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
			UserAbilityRules: []AbilityRule{
				AbilityRule{Action: "manage", Subject: "all"},
			},
		},
	)
}