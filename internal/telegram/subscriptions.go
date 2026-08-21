package telegram

import (
	"context"
	"cryptobot/internal/subscription"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartAuto(
	ctx context.Context,
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
	subService *subscription.SubscriptionService,
) {

	args := strings.Fields(update.Message.Text)

	if len(args) < 3 {
		SendText(
			update,
			bot,
			"Usage: /subscribe <currency> <minutes>",
		)
		return
	}

	currency := args[1]

	if currency != "bitcoin" &&
		currency != "ethereum" {

		SendText(
			update,
			bot,
			"Unsupported currency. Available: bitcoin, ethereum.",
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
			"Minutes must be greater than 0.",
		)
		return
	}

	err = subService.Subscribe(
		ctx,
		update.Message.Chat.ID,
		minutes,
		currency,
	)

	if err != nil {
		SendText(
			update,
			bot,
			"Failed to create subscription.",
		)
		return
	}

	SendText(
		update,
		bot,
		fmt.Sprintf(
			"Subscription enabled: %s every %d minutes.",
			currency,
			minutes,
		),
	)

}

func StopAuto(
	ctx context.Context,
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
	service *subscription.SubscriptionService,
) {

	err := service.Unsubscribe(
		ctx,
		update.Message.Chat.ID,
	)

	if err != nil {
		SendText(update, bot, "Failed to unsubscribe")
		return
	}

	SendText(
		update,
		bot,
		"Subscription disabled",
	)

}
