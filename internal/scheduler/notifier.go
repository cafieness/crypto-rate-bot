package scheduler

import (
	"cryptobot/internal/service"
	"cryptobot/internal/transport/telegram"
	"log"
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

func (n *Notifier) Start() {

	const notifierInterval = time.Minute

	ticker := time.NewTicker(notifierInterval)

	defer ticker.Stop()

	for range ticker.C {
		n.notify()
	}
}

func (n *Notifier) notify() {
	subs, err := n.subService.GetActive()

	if err != nil {
		log.Println(err)
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
		currency,
		n.rateService,
	)
	if err != nil {
		log.Println(err)
		return
	}

	msg := tgbotapi.NewMessage(
		chatID,
		text,
	)

	_, err = n.bot.Send(msg)

	if err != nil {
		log.Println(err)
		return
	}

	err = n.subService.UpdateLastSent(
		chatID,
		currency,
	)

	if err != nil {
		log.Println(err)
	}

}
