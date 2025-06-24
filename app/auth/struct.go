package auth

import "github.com/thoriqwildan/svdclone-be/pkg/global"

type LoginRequest struct {
	Email		string `json:"email" validate:"required,email"`
	Password	string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Name				string `json:"name" validate:"required"`
	Email				string `json:"email" validate:"required,email"`
	ProfileUrl	string `json:"profile_url,omitempty"`
	Password		string `json:"password" validate:"required,min=3"`
}

type AbilityRule struct {
	Action  string `json:"action"`
	Subject string `json:"subject"`
}

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
	UserData	 global.UserResponse `json:"userData"`
	UserAbilityRules []AbilityRule `json:"userAbilityRules,omitempty"`
}