package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func callbackHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

	switch update.CallbackQuery.Data {
	case "backToMainMenu":
		msg.Text = greeting
		msg.ReplyMarkup = mainMenuKeyboard
	}

	errorCh := make(chan error)

	go func() {
		_, err := bot.Send(msg)

		errorCh <- err
	}()

	err := <-errorCh
	if err != nil {
		log.Panic(err)
	}
}
