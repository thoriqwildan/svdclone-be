package middleware

import (
	"errors"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/thoriqwildan/svdclone-be/pkg/config"
	"github.com/thoriqwildan/svdclone-be/pkg/global"
)

func JWTProtected() func(*fiber.Ctx) error {
	jwtwareConfig := jwtware.Config{	
		SigningKey: jwtware.SigningKey{Key: []byte(config.GetEnv("JWT_SECRET", "jwt_secret"))},
		ContextKey: "user",
		ErrorHandler: jwtError,
		SuccessHandler: verifyTokenExpiration,
	}
	return jwtware.New(jwtwareConfig)
}

func verifyTokenExpiration(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	expires := int64(claims["exp"].(float64))
	if time.Now().Unix() > expires {
		return jwtError(c, errors.New("token expired"))
	}

	return c.Next()
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(global.ErrorResponse{
			Success: false,
			Message: "JWT not Found",
			Errors: map[string]string{
				"error": err.Error(),
			},
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(global.ErrorResponse{
			Success: false,
			Message: "You're unauthorized",
			Errors: map[string]string{
				"error": err.Error(),
			},
		})
}