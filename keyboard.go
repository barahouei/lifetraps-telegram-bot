package main

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var mainMenuKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("می‌خوام تست بدم", "quiz"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("می‌خوام تله‌هام رو ببینم", "myLifeTraps"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("راهنما", "guide"),
	),
)

var backToMainMenuKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("برگشت به منوی اصلی", "backToMainMenu"),
	),
)

//This functions creates an inline keyboard based on question.
func scoreButtons(questionID int, categoryID int) tgbotapi.InlineKeyboardMarkup {
	qID := strconv.Itoa(questionID)
	cid := strconv.Itoa(categoryID)

	sbKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Completely untrue of me", "score=1-qid="+qID+"-cid="+cid),
			tgbotapi.NewInlineKeyboardButtonData("Mostly untrue of me", "score=2-qid="+qID+"-cid="+cid),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Slightly more true than untrue of me", "score=3-qid="+qID+"-cid="+cid),
			tgbotapi.NewInlineKeyboardButtonData("Moderately true of me", "score=4-qid="+qID+"-cid="+cid),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Mostly true of me", "score=5-qid="+qID+"-cid="+cid),
			tgbotapi.NewInlineKeyboardButtonData("Describes me perfectly", "score=6-qid="+qID+"-cid="+cid),
		),
	)

	return sbKeyboard
}
