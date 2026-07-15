package telegram

import (
	"cryptobot/internal/service"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleRates(
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
	rateService *service.RateService,
) {
	args := strings.Fields(update.Message.Text)

	if len(args) == 1 {

		btc, err := BuildRateMessage(
			"bitcoin",
			rateService,
		)
		if err != nil {
			SendText(update, bot, "Нет данных")
			return
		}

		eth, err := BuildRateMessage(
			"ethereum",
			rateService,
		)
		if err != nil {
			SendText(update, bot, "Нет данных")
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
		args[1],
		rateService,
		update,
		bot,
	)
}

func SendSingleRate(
	currency string,
	rateService *service.RateService,
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
) {

	text, err := BuildRateMessage(
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
