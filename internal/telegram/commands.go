package telegram

import (
	"cryptobot/internal/rate"
	"cryptobot/internal/subscription"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(
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
			"Hi! Use /rates to see crypto currencies rate or use /help",
		)
	case "help":
		SendText(
			update,
			bot,
			`
/rates - show all crypto rates

/rates bitcoin - show bitcoin rate

/rates ethereum - show ethereum rate

/subscribe [currency] [minutes]

Example:
/subscribe bitcoin 60

/unsubscribe: stop subscription

Available currencies:
bitcoin
ethereum
			`,
		)
	case "rates":
		HandleRates(
			update,
			bot,
			rateService,
		)

	case "subscribe":
		StartAuto(
			update,
			bot,
			subService,
		)

	case "unsubscribe":
		StopAuto(
			update,
			bot,
			subService,
		)

	default:
		SendText(
			update,
			bot,
			"Unknown command. Please use /help",
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
			Description: "subscribe and get notify you with current rates and stats",
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
