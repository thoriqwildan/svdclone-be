package router

import (
	"github.com/gofiber/fiber/v2"
	paymentchannel "github.com/thoriqwildan/svdclone-be/app/payment_channel"
	paymentmethod "github.com/thoriqwildan/svdclone-be/app/payment_method"
)

func PaymentMethodRoutes(app *fiber.App) {
	p := app.Group("/api/payment-methods")

	p.Post("/", paymentmethod.CreatePaymentMethod)
	p.Get("/", paymentmethod.GetPaymentMethods)
	p.Get("/:id", paymentmethod.GetPaymentMethodById)
	p.Put("/:id", paymentmethod.UpdatePaymentMethod)
	p.Delete("/:id", paymentmethod.DeletePaymentMethod)
}

func PaymentChannelRoutes(app *fiber.App) {
	p := app.Group("/api/payment-channels")

	p.Post("/", paymentchannel.CreatePaymentChannel)
	p.Get("/", paymentchannel.GetPaymentChannels)
	p.Get("/:id", paymentchannel.GetPaymentChannelById)
	p.Put("/:id", paymentchannel.UpdatePaymentChannel)
	p.Delete("/:id", paymentchannel.DeletePaymentChannel)
}