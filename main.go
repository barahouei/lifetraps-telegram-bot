package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("LIFETRAPS_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil && update.Message.IsCommand() {
			userID := update.Message.Chat.ID
			msgID := update.Message.MessageID

			for i := msgID; i > 0; i-- {
				go bot.Request(tgbotapi.NewDeleteMessage(userID, i))
			}

			go commandHandling(bot, update)
		} else {
			go bot.Request(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID))

			go messageHandling(bot, update)
		}
	}
}
