package scheduler

import (
	"context"
	"cryptobot/internal/service"
	"log"
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

func (u *Updater) Start() {
	u.update()
	ticker := time.NewTicker(u.Interval)

	defer ticker.Stop()

	for range ticker.C {
		u.update()
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

		cancel()

		if err != nil {
			log.Printf(
				"failed to fetch %s: %v",
				coin, err,
			)
			continue
		}
		err = u.Service.Save(
			coin,
			price,
		)
		if err != nil {

			log.Printf(
				"failed save %s: %v",
				coin,
				err,
			)

			continue
		}
		log.Printf(
			"%s updated: %.2f",
			coin,
			price,
		)
	}
}
