package rate

import (
	"context"
)

type Rate struct {
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
}

type RateRepository interface {
	Save(
		ctx context.Context,
		currency string,
		price float64,
	) error

	GetLatest(
		ctx context.Context,
		currency string,
	) (Rate, error)

	GetDailyMinMax(
		ctx context.Context,
		currency string,
	) (float64, float64, error)

	GetHourlyChange(
		ctx context.Context,
		currency string,
	) (float64, error)
}

type Service interface {
	GetLatest(
		ctx context.Context,
		currency string,
	) (Rate, error)
}

type RateWriter interface {
	Save(
		ctx context.Context,
		currency string,
		price float64,
	) error
}
