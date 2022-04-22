package main

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func callbackHandling(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

	callbackText := update.CallbackQuery.Data
	var score, qID, cID string

	if strings.Contains(callbackText, "score") {
		spl := strings.Split(callbackText, "-")

		s := strings.Split(spl[0], "=")
		score = s[1]

		q := strings.Split(spl[1], "=")
		qID = q[1]

		c := strings.Split(spl[2], "=")
		cID = c[1]
	}

	switch update.CallbackQuery.Data {
	case "backToMainMenu":
		msg.Text = greeting
		msg.ReplyMarkup = mainMenuKeyboard
	case "quiz":
		questionChan := make(chan string)
		cIDChan := make(chan int)
		errorChan := make(chan error)

		go func() {
			question, cID, err := questionWalker(1)

			questionChan <- question
			cIDChan <- cID
			errorChan <- err
		}()

		question := <-questionChan
		cID := <-cIDChan
		err := <-errorChan

		if err != nil {
			log.Fatal(err)
		}

		msg.Text = question
		msg.ReplyMarkup = scoreButtons(1, cID)
	case "score=" + score + "-qid=" + qID + "-cid=" + cID:
		user := users{}
		user.telegramID = update.CallbackQuery.From.ID
		score, _ := strconv.Atoi(score)
		qID, _ := strconv.Atoi(qID)
		cID, _ := strconv.Atoi(cID)

		go setScore(user.telegramID, score, qID, cID)

		questionChan := make(chan string)
		cIDChan := make(chan int)
		errorChan := make(chan error)

		go func() {
			question, cID, err := questionWalker(qID + 1)

			questionChan <- question
			cIDChan <- cID
			errorChan <- err
		}()

		question := <-questionChan
		categotyID := <-cIDChan
		err := <-errorChan

		if err != nil {
			msg.Text = testEnded
			msg.ReplyMarkup = backToMainMenuKeyboard
		} else {
			msg.Text = question
			msg.ReplyMarkup = scoreButtons(qID+1, categotyID)
		}
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
