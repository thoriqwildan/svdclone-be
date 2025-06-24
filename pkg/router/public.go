package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thoriqwildan/svdclone-be/app/auth"
)

func GeneralRoutes(app *fiber.App) {
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to SVDClone API")
	})
}

func AuthRoutes(app *fiber.App) {
	a := app.Group("/api/auth")
	a.Post("/login", auth.Login)
	a.Post("/register", auth.Register)
}