package telegram

import (
	"context"
	"cryptobot/internal/service"
	"errors"
	"fmt"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewBot(token string) (*tgbotapi.BotAPI, error) {
	if token == "" {
		return nil, errors.New("telegram token is empty")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	slog.Info(
		"telegram bot authorized",
		"username",
		bot.Self.UserName,
	)

	return bot, nil
}

func Run(ctx context.Context, bot *tgbotapi.BotAPI, rateService *service.RateService, subService *service.SubscriptionService) {

	err := SetCommands(bot)
	if err != nil {
		slog.Error(
			"failed to set commands",
			"error", err,
		)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for {
		select {

		case update, ok := <-updates:
			if !ok {
				return
			}
			if update.Message == nil {
				continue
			}

			HandleCommand(
				update,
				bot,
				rateService,
				subService,
			)

		case <-ctx.Done():
			slog.Info("telegram stopped")
			return
		}
	}

}
