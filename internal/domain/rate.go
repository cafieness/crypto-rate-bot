package domain

type Rate struct {
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
}
