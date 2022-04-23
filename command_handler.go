package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func commandHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if update.Message.Command() == "start" {
		go checkUserExists(update)

		msg.Text = greeting
		msg.ReplyMarkup = mainMenuKeyboard
	} else if update.Message.Command() == "menu" {
		msg.Text = menu
		msg.ReplyMarkup = mainMenuKeyboard
	} else if update.Message.Command() == "guide" {
		msg.Text = guide
		msg.ReplyMarkup = backToMainMenuKeyboard
	} else {
		msg.Text = wrongCommand
		msg.ReplyMarkup = backToMainMenuKeyboard
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
