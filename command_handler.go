package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func commandHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if update.Message.Command() == "start" {
		msg.Text = "سلام به بات تله‌های زندگی خوش آمدید."
	}

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
