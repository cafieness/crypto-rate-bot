package domain

type RateRepository interface {
	Save(
		currency string,
		price float64,
	) error

	GetLatest(
		currency string,
	) (Rate, error)

	GetDailyMinMax(
		currency string,
	) (float64, float64, error)

	GetHourlyChange(
		currency string,
	) (float64, error)
}

type SubscriptionRepository interface {
	Create(
		sub Subscription,
	) error
	Deactivate(
		chatID int64,
	) error
	GetActive() (
		[]Subscription,
		error)

	UpdateLastSent(
		chatID int64,
		currency string,
	) error
	Exists(
		chatID int64,
		currency string,
	) (bool, error)
}
