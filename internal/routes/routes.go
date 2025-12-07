package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kodra-pay/fee-service/internal/handlers"
	"github.com/kodra-pay/fee-service/internal/services"
)

func Register(app *fiber.App, serviceName string) {
	health := handlers.NewHealthHandler(serviceName)
	health.Register(app)

	feeHandler := handlers.NewFeeHandler(services.NewFeeService())

	api := app.Group("/fees")
	api.Post("/quote", feeHandler.Quote)
	api.Get("/rules", feeHandler.Rules)
}
