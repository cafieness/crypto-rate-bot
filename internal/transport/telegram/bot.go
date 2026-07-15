package telegram

import (
	"cryptobot/internal/service"
	"errors"
	"fmt"
	"log"

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

	log.Printf("Authorized as %s", bot.Self.UserName)

	return bot, nil
}

func Run(bot *tgbotapi.BotAPI, rateService *service.RateService, subService *service.SubscriptionService) {

	err := SetCommands(bot)
	if err != nil {
		log.Println("failed to set commands:", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		HandleCommand(
			update,
			bot,
			rateService,
			subService,
		)

	}

}
