package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func commandHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	userID := update.Message.Chat.ID
	msgID := update.Message.MessageID

	for i := msgID; i > 0; i-- {
		go bot.Request(tgbotapi.NewDeleteMessage(userID, i))
	}

	msg := tgbotapi.NewMessage(userID, "")

	if update.Message.Command() == "start" {
		msg.Text = greeting
	} else if update.Message.Command() == "help" {
		msg.Text = guide
	} else {
		msg.Text = wrongCommand
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
