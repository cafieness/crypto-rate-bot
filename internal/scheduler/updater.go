package scheduler

import (
	"context"
	"cryptobot/internal/domain"
	"log/slog"
	"time"
)

type Updater struct {
	Service  domain.RateWriter
	Fetch    func(context.Context, string) (float64, error)
	Interval time.Duration
}

func NewUpdater(service domain.RateWriter,
	fetch func(context.Context, string) (float64, error),
	interval time.Duration,
) *Updater {
	return &Updater{
		Service:  service,
		Fetch:    fetch,
		Interval: interval,
	}
}

func (u *Updater) Start(ctx context.Context) {

	u.Update()

	ticker := time.NewTicker(u.Interval)
	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:
			u.Update()

		case <-ctx.Done():
			slog.Info("updater stopped")
			return
		}
	}
}

func (u *Updater) Update() {
	coins := []string{
		"bitcoin",
		"ethereum",
	}

	for _, coin := range coins {

		ctx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)

		price, err := u.Fetch(ctx, coin)

		defer cancel()

		if err != nil {
			slog.Error(
				"failed to fetch",
				"currency", coin,
				"error", err,
			)
			continue
		}

		err = u.Service.Save(
			ctx,
			coin,
			price,
		)

		if err != nil {
			slog.Error(
				"failed to save",
				"currency", coin,
				"error", err,
			)
			continue
		}

		slog.Info(
			"rate updated",
			"currency", coin,
			"price", price,
		)
	}
}
