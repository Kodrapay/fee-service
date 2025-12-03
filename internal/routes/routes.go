package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kodra-pay/fee-service/internal/handlers"
	"github.com/kodra-pay/fee-service/internal/services"
)

func Register(app *fiber.App, service string) {
	health := handlers.NewHealthHandler(service)
	health.Register(app)

	svc := services.NewFeeService()
	h := handlers.NewFeeHandler(svc)
	api := app.Group("/fees")
	api.Post("/quote", h.Quote)
	api.Get("/rules", h.Rules)
}
