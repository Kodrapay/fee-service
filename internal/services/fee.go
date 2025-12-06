package services

import (
	"context"
	"strings"

	"github.com/kodra-pay/fee-service/internal/dto"
)

type FeeService struct{}

func NewFeeService() *FeeService { return &FeeService{} }

func (s *FeeService) Quote(_ context.Context, req dto.FeeQuoteRequest) dto.FeeQuoteResponse {
	channel := strings.ToLower(strings.TrimSpace(req.Channel))
	if channel == "" {
		channel = "card"
	}
	currency := strings.ToUpper(strings.TrimSpace(req.Currency))
	if currency == "" {
		currency = "NGN"
	}

	rate := 0.015
	flat := 0.0
	cap := 0.0

	switch channel {
	case "card":
		rate = 0.017
		cap = 2000 // NGN cap for cards
	case "bank_transfer", "bank":
		rate = 0.01
		flat = 50
	case "mobile_money", "momo":
		rate = 0.015
	}

	fee := req.Amount*rate + flat
	capped := false
	if cap > 0 && fee > cap {
		fee = cap
		capped = true
	}

	return dto.FeeQuoteResponse{
		TotalFee:   fee,
		BaseAmount: req.Amount,
		Currency:   currency,
		Rate:       rate,
		Flat:       flat,
		Capped:     capped,
		Cap:        cap,
		Channel:    channel,
	}
}

func (s *FeeService) Rules(_ context.Context) []string {
	return []string{
		"Card: 1.7% capped at NGN 2000",
		"Bank transfer: 1% + NGN 50 flat",
		"Mobile money: 1.5%",
	}
}
