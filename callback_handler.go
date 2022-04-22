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
	case "quiz":
		question, cID := questionWalker(1)

		msg.Text = question
		msg.ReplyMarkup = scoreButtons(1, cID)
	case "myLifeTraps":
		msg.Text = "این یک متنی موقتی برای دکمه تله‌های من است."
		msg.ReplyMarkup = backToMainMenuKeyboard
	case "guide":
		msg.Text = guide
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
