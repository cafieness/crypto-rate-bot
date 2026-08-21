package telegram

import (
	"context"
	"cryptobot/internal/rate"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleRates(
	ctx context.Context,
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
	rateService *rate.RateService,
) {
	args := strings.Fields(update.Message.Text)

	if len(args) == 1 {

		btc, err := BuildRateMessage(
			ctx,
			"bitcoin",
			rateService,
		)
		if err != nil {
			SendText(update, bot, "Rate data is currently unavailable.")
			return
		}

		eth, err := BuildRateMessage(
			ctx,
			"ethereum",
			rateService,
		)
		if err != nil {
			SendText(update, bot, "Rate data is currently unavailable.")
			return
		}

		SendText(
			update,
			bot,
			btc+"\n\n"+eth,
		)

		return
	}

	SendSingleRate(
		ctx,
		args[1],
		rateService,
		update,
		bot,
	)
}

func SendSingleRate(
	ctx context.Context,
	currency string,
	rateService *rate.RateService,
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
) {

	text, err := BuildRateMessage(
		ctx,
		currency,
		rateService,
	)

	if err != nil {
		SendText(
			update,
			bot,
			"Currency not found. Available currencies:\nbitcoin, ethereum",
		)
		return
	}

	SendText(
		update,
		bot,
		text,
	)
}
