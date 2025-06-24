package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thoriqwildan/svdclone-be/pkg/config"
)

func GenerateToken(email string, admin bool) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(30)).Unix()

	t, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET", "jwt_secret")))

	if err != nil {
		return "", err
	}
	return t, nil
}