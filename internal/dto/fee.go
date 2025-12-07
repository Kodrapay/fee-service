package dto

type FeeQuoteRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Channel  string  `json:"channel"`
}

type FeeQuoteResponse struct {
	TotalFee   float64 `json:"total_fee"`
	BaseAmount float64 `json:"base_amount"`
	Currency   string  `json:"currency"`
	Rate       float64 `json:"rate"`
	Flat       float64 `json:"flat"`
	Capped     bool    `json:"capped"`
	Cap        float64 `json:"cap,omitempty"`
	Channel    string  `json:"channel,omitempty"`
}

type FeeRule struct {
	Channel string  `json:"channel"`
	Rate    float64 `json:"rate"`
	Flat    float64 `json:"flat"`
	Cap     float64 `json:"cap"`
	Note    string  `json:"note,omitempty"`
}
