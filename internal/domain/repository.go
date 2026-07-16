package domain

import "context"

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

type SubscriptionRepository interface {
	Create(
		ctx context.Context,
		sub Subscription,
	) error
	Deactivate(
		ctx context.Context,
		chatID int64,
	) error
	GetActive(context.Context) (
		[]Subscription,
		error)

	UpdateLastSent(
		ctx context.Context,
		chatID int64,
		currency string,
	) error
	Exists(
		ctx context.Context,
		chatID int64,
		currency string,
	) (bool, error)
}
