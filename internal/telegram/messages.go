package telegram

import (
	"context"
	"cryptobot/internal/rate"
	"fmt"
	"log/slog"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendText(
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
	text string,
) {

	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		text,
	)

	_, err := bot.Send(msg)

	if err != nil {
		slog.Error("Failed to send message")
	}
}

func BuildRateMessage(ctx context.Context, currency string, rateService *rate.RateService) (string, error) {
	rate, err := rateService.GetLatest(ctx, currency)

	if err != nil {
		return "", err
	}

	min, max, err := rateService.GetDailyMinMax(ctx, currency)

	if err != nil {
		return "", err
	}

	change, err := rateService.GetHourlyChange(ctx, currency)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		`%s
		Current rate: %.2f$
		Daily min: %.2f$
		Daily max: %.2f$
		Hourly change: %.2f%%`,
		strings.ToUpper(currency),
		rate.Price,
		min,
		max,
		change,
	), nil
}
