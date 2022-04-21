package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func messageHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	msg.Text = wrongCommand

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
