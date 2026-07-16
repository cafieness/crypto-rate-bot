package main

import (
	"context"
	"cryptobot/internal/config"
	"cryptobot/internal/fetcher"
	"cryptobot/internal/logger"
	"cryptobot/internal/repository/postgres"
	"cryptobot/internal/scheduler"
	"cryptobot/internal/service"
	transporthttp "cryptobot/internal/transport/http"
	"cryptobot/internal/transport/telegram"
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	initLogger()

	cfg := config.Load()

	rateService, subService, bot, db, err := buildApp(cfg)
	if err != nil {
		return
	}

	defer db.Close()

	startBackgroundWorkers(ctx, cfg, rateService, subService, bot)

	server := newHTTPServer(rateService)

	waitForShutdown(ctx, cancel, server)
}

// initialize application logger

func initLogger() {
	appLogger := logger.New()
	slog.SetDefault(appLogger)

	slog.Info("Starting app...")
}

// build application dependencies

func buildApp(cfg *config.Config) (
	rateService *service.RateService,
	subService *service.SubscriptionService,
	bot *tgbotapi.BotAPI,
	db *sql.DB,
	err error,
) {
	db, err = postgres.Connect(cfg.DatabaseURL)

	if err != nil {
		slog.Error("Failed database connection", "error", err)
		return nil, nil, nil, nil, err
	}

	repo := postgres.NewRateRepository(db)
	rateService = service.NewRateService(repo)

	subRepo := postgres.NewSubscriptionRepository(db)
	subService = service.NewSubscriptionService(
		subRepo,
	)
	bot, err = telegram.NewBot(
		cfg.TelegramAPIKey,
	)

	if err != nil {
		slog.Error("Failed telegram new bot", "error", err)
		return nil, nil, nil, nil, err
	}
	return rateService, subService, bot, db, nil

}

// start background workers

func startBackgroundWorkers(
	ctx context.Context,
	cfg *config.Config,
	rateService *service.RateService,
	subService *service.SubscriptionService,
	bot *tgbotapi.BotAPI,
) {
	notifier := scheduler.NewNotifier(
		rateService,
		subService,
		bot,
	)

	go telegram.Run(
		ctx,
		bot,
		rateService,
		subService)

	go notifier.Start(ctx)

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
	go updater.Start(ctx)
}

// start http server

func newHTTPServer(rateService *service.RateService) *http.Server {
	router := transporthttp.NewRouter(rateService)

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	go func() {
		slog.Info("Server started at :8081")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", "error", err)
			return
		}

	}()
	return server
}

//wait for shutdown signal and stop app

func waitForShutdown(
	ctx context.Context,
	cancel context.CancelFunc,
	server *http.Server,
) {
	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-quit

	slog.Info("Shutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)

	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("Failed to shut down", "error", err)
	}

	cancel()
}
