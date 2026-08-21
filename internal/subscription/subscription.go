package subscription

import (
	"context"
	"time"
)

type Subscription struct {
	ChatID          int64
	IntervalMinutes int
	Currency        string
	Active          bool
	LastSentAt      *time.Time
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
