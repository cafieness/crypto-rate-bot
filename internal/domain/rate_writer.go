package domain

import "context"

type RateWriter interface {
	Save(
		ctx context.Context,
		currency string,
		price float64,
	) error
}
