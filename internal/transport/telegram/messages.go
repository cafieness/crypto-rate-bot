package telegram

import (
	"log/slog"

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
