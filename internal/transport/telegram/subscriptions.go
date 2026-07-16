package telegram

import (
	"context"
	"cryptobot/internal/service"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartAuto(
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
	subService *service.SubscriptionService,
) {

	args := strings.Fields(update.Message.Text)

	if len(args) < 3 {
		SendText(
			update,
			bot,
			"Use: /subscribe [currency] [time]",
		)
		return
	}

	currency := args[1]

	if currency != "bitcoin" &&
		currency != "ethereum" {

		SendText(
			update,
			bot,
			"Available: bitcoin, ethereum",
		)
		return
	}

	minutes, err := strconv.Atoi(args[2])

	if err != nil {
		SendText(
			update,
			bot,
			"Minutes should be a number",
		)
		return
	}
	if minutes <= 0 {
		SendText(
			update,
			bot,
			"Time should be more than 0",
		)
		return
	}

	err = subService.Subscribe(
		context.Background(),
		update.Message.Chat.ID,
		minutes,
		currency,
	)

	if err != nil {
		SendText(
			update,
			bot,
			err.Error(),
		)
		return
	}

	SendText(
		update,
		bot,
		"Subscription is enabled!",
	)

}

func StopAuto(
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
	service *service.SubscriptionService,
) {

	err := service.Unsubscribe(
		context.Background(),
		update.Message.Chat.ID,
	)

	if err != nil {
		SendText(update, bot, "Failed to unsubscribe")
		return
	}

	SendText(
		update,
		bot,
		"Subscription is disabled",
	)

}
