package dto

type FeeQuoteRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Channel  string  `json:"channel"`
}

type FeeQuoteResponse struct {
	TotalFee float64 `json:"total_fee"`
	Currency string  `json:"currency"`
}
