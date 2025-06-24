package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/thoriqwildan/svdclone-be/pkg/config"
	"github.com/thoriqwildan/svdclone-be/pkg/router"
)

func Serve() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Inisiasi Routeerr
	router.GeneralRoutes(app)
	router.AuthRoutes(app)
	router.PaymentMethodRoutes(app)
	router.PaymentChannelRoutes(app)

	app.Listen(":" + config.GetEnv("PORT", "3000"))
}