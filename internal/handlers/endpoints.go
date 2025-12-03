package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kodra-pay/fee-service/internal/dto"
	"github.com/kodra-pay/fee-service/internal/services"
)

type FeeHandler struct {
	svc *services.FeeService
}

func NewFeeHandler(svc *services.FeeService) *FeeHandler { return &FeeHandler{svc: svc} }

func (h *FeeHandler) Quote(c *fiber.Ctx) error {
	var req dto.FeeQuoteRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	return c.JSON(h.svc.Quote(c.Context(), req))
}

func (h *FeeHandler) Rules(c *fiber.Ctx) error {
	return c.JSON(h.svc.Rules(c.Context()))
}
