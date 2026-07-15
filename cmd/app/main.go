package main

import (
	"context"
	"cryptobot/internal/config"
	"cryptobot/internal/fetcher"
	"cryptobot/internal/repository/postgres"
	"cryptobot/internal/scheduler"
	"cryptobot/internal/service"
	transporthttp "cryptobot/internal/transport/http"
	"cryptobot/internal/transport/telegram"
	"log"
	"net/http"
	"time"
)

func main() {

	cfg := config.Load()

	db, err := postgres.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := postgres.NewRateRepository(db)
	rateService := service.NewRateService(repo)

	subRepo := postgres.NewSubscriptionRepository(db)
	subscriptionService := service.NewSubscriptionService(
		subRepo,
	)
	bot, err := telegram.NewBot(
		cfg.TelegramAPIKey,
	)

	if err != nil {
		log.Fatal(err)
	}

	notifier := scheduler.NewNotifier(
		rateService,
		subscriptionService,
		bot,
	)

	go telegram.Run(bot,
		rateService,
		subscriptionService)

	go notifier.Start()

	client := fetcher.NewClient(cfg)

	updater := scheduler.NewUpdater(
		rateService,
		func(
			ctx context.Context,
			coin string,
		) (float64, error) {
			return client.FetchRate(ctx, coin)
		},
		time.Minute,
	)
	go updater.Start()

	router := transporthttp.NewRouter(rateService)

	go func() {
		log.Println("HTTP server started: 8081")
		err := http.ListenAndServe(
			":8081",
			router,
		)
		if err != nil {
			log.Fatal(err)
		}

	}()
	select {}
}
