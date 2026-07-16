package scheduler

import (
	"context"
	"cryptobot/internal/service"
	"log/slog"
	"time"
)

type Updater struct {
	Service  *service.RateService
	Fetch    func(context.Context, string) (float64, error)
	Interval time.Duration
}

func NewUpdater(service *service.RateService,
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

	u.update()

	ticker := time.NewTicker(u.Interval)
	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:
			u.update()

		case <-ctx.Done():
			slog.Info("updater stopped")
			return
		}
	}
}

func (u *Updater) update() {
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
