package services

import (
	"context"

	"github.com/kodra-pay/fee-service/internal/dto"
)

type FeeService struct{}

func NewFeeService() *FeeService { return &FeeService{} }

func (s *FeeService) Quote(_ context.Context, req dto.FeeQuoteRequest) dto.FeeQuoteResponse {
	// simple stub: 1.5% fee
	return dto.FeeQuoteResponse{
		TotalFee: req.Amount * 0.015,
		Currency: req.Currency,
	}
}

func (s *FeeService) Rules(_ context.Context) []string {
	return []string{"flat 1.5% processing fee (stub)"}
}
