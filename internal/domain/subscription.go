package domain

import "time"

type Subscription struct {
	ChatID          int64
	IntervalMinutes int
	Currency        string
	Active          bool
	LastSentAt      *time.Time
}
