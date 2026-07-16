package scheduler

import (
	"context"
	"cryptobot/internal/service"
	"cryptobot/internal/transport/telegram"
	"log/slog"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Notifier struct {
	rateService *service.RateService
	subService  *service.SubscriptionService
	bot         *tgbotapi.BotAPI
}

func NewNotifier(
	rateService *service.RateService,
	subService *service.SubscriptionService,
	bot *tgbotapi.BotAPI,
) *Notifier {
	return &Notifier{
		rateService: rateService,
		subService:  subService,
		bot:         bot,
	}
}

func (n *Notifier) Start(ctx context.Context) {

	const notifierInterval = time.Minute

	ticker := time.NewTicker(notifierInterval)

	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:
			n.notify(ctx)

		case <-ctx.Done():
			slog.Info("notifier stopped")
			return
		}
	}
}

func (n *Notifier) notify(ctx context.Context) {
	subs, err := n.subService.GetActive(ctx)

	if err != nil {
		slog.Error("Failed get active", "error", err)
		return
	}

	for _, sub := range subs {

		if sub.LastSentAt != nil {
			if time.Since(*sub.LastSentAt) <
				time.Duration(sub.IntervalMinutes)*time.Minute {
				continue
			}
		}

		n.send(
			sub.ChatID,
			sub.Currency,
		)
	}
}

func (n *Notifier) send(
	chatID int64,
	currency string,
) {
	text, err := telegram.BuildRateMessage(
		context.Background(),
		currency,
		n.rateService,
	)
	if err != nil {
		slog.Error("Failed build rate message", "error", err)
		return
	}

	msg := tgbotapi.NewMessage(
		chatID,
		text,
	)

	_, err = n.bot.Send(msg)

	if err != nil {
		slog.Error("Failed to send message", "error", err)
		return
	}

	err = n.subService.UpdateLastSent(
		context.Background(),
		chatID,
		currency,
	)

	if err != nil {
		slog.Error("Failed update last sent", "error", err)
	}

}
