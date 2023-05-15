package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("LIFETRAPS_APITOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		go SafeHandle(bot, update)
	}
}

// SafeHandle safely handles every request we have
// and if any error happens it will recover from error and does not let bot break.
func SafeHandle(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*Somthing is wrong!", err)
		}
	}()

	handleAll(bot, update)
}

// handleAll handles all request.
func handleAll(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil && update.Message.IsCommand() {
		userID := update.Message.Chat.ID
		msgID := update.Message.MessageID

		go func() {
			for i := msgID - 20; i < msgID; i++ {
				go bot.Request(tgbotapi.NewDeleteMessage(userID, i))
			}
		}()

		go commandHandling(bot, update)
	} else if update.CallbackQuery != nil {
		go bot.Request(tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID))

		go callbackHandling(bot, update)
	} else if update.Message != nil {
		go bot.Request(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID))

		go messageHandling(bot, update)
	} else {
		go bot.Request(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID))

		go messageHandling(bot, update)
	}
}
