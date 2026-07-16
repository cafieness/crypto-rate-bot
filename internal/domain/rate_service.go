package domain

import "context"

type RateService interface {
	GetLatest(
		ctx context.Context,
		currency string,
	) (Rate, error)
}
