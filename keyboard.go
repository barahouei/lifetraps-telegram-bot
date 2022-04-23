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
			tgbotapi.NewInlineKeyboardButtonData("کاملا نادرسته", "score=1-qid="+qID+"-cid="+cid),
			tgbotapi.NewInlineKeyboardButtonData("تقریبا نادرسته", "score=2-qid="+qID+"-cid="+cid),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("بیشتر درسته تا غلط", "score=3-qid="+qID+"-cid="+cid),
			tgbotapi.NewInlineKeyboardButtonData("نسبتا درسته", "score=4-qid="+qID+"-cid="+cid),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("تقریبا درسته", "score=5-qid="+qID+"-cid="+cid),
			tgbotapi.NewInlineKeyboardButtonData("کاملا درسته", "score=6-qid="+qID+"-cid="+cid),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("برگشت به منوی اصلی", "backToMainMenu"),
		),
	)

	return sbKeyboard
}
