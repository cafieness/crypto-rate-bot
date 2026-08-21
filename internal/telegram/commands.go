package telegram

import (
	"context"
	"cryptobot/internal/rate"
	"cryptobot/internal/subscription"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(
	ctx context.Context,
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
	rateService *rate.RateService,
	subService *subscription.SubscriptionService,
) {

	command := update.Message.Command()

	switch command {

	case "start":
		SendText(
			update,
			bot,
			"Hi! Use /rates to view current crypto rates, or /help to see all commands.",
		)
	case "help":
		SendText(
			update,
			bot,
			`
Available commands:

/rates - show all crypto rates
/rates bitcoin - show Bitcoin rate
/rates ethereum - show Ethereum rate

/subscribe <currency> <minutes>
Example: /subscribe bitcoin 60

/unsubscribe - stop notifications

Supported currencies:
bitcoin, ethereum
			`,
		)
	case "rates":
		HandleRates(
			ctx,
			update,
			bot,
			rateService,
		)

	case "subscribe":
		StartAuto(
			ctx,
			update,
			bot,
			subService,
		)

	case "unsubscribe":
		StopAuto(
			ctx,
			update,
			bot,
			subService,
		)

	default:
		SendText(
			update,
			bot,
			"Unknown command. Use /help to see available commands.",
		)
	}

}

func ClearCommands(bot *tgbotapi.BotAPI) error {
	config := tgbotapi.NewDeleteMyCommands()

	_, err := bot.Request(config)

	return err
}

func SetCommands(bot *tgbotapi.BotAPI) error {

	commands := []tgbotapi.BotCommand{
		{
			Command:     "start",
			Description: "start the bot",
		},
		{
			Command:     "help",
			Description: "command list",
		},
		{
			Command:     "rates",
			Description: "show all crypto rates",
		},
		{
			Command:     "subscribe",
			Description: "subscribe to get rate updates",
		},
		{
			Command:     "unsubscribe",
			Description: "stop subscription",
		},
	}

	config := tgbotapi.NewSetMyCommands(commands...)

	_, err := bot.Request(config)

	return err
}
