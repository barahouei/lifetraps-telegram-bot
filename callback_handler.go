package main

import (
	"fmt"
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
		msg.Text = menu
		msg.ReplyMarkup = mainMenuKeyboard
	case "quiz":
		user := users{telegramID: update.CallbackQuery.From.ID}

		go resetTest(user.telegramID)

		countChan := make(chan int)

		go func() {
			num := numberOfQuestions()
			countChan <- num
		}()

		count := <-countChan

		questionChan := make(chan string)
		cIDChan := make(chan int)
		errorChan := make(chan error)

		go func() {
			question, cID, err := questionWalker(1)

			questionChan <- question
			cIDChan <- cID
			errorChan <- err
		}()

		question := fmt.Sprintf("سوال شماره %d از مجموع %d سوال:\n", 1, count)
		question += <-questionChan
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

		countChan := make(chan int)

		go func() {
			num := numberOfQuestions()
			countChan <- num
		}()

		count := <-countChan

		questionChan := make(chan string)
		cIDChan := make(chan int)
		errorChan := make(chan error)

		go func() {
			question, cID, err := questionWalker(qID + 1)

			questionChan <- question
			cIDChan <- cID
			errorChan <- err
		}()

		question := fmt.Sprintf("سوال شماره %d از مجموع %d سوال:\n", qID+1, count)
		question += <-questionChan
		categotyID := <-cIDChan
		err := <-errorChan

		if err != nil {
			go setTestCompleted(user.telegramID)

			ltChan := make(chan []string)

			go func() {
				lt := showLifetraps(user.telegramID)
				ltChan <- lt
			}()

			lifetraps := <-ltChan

			if len(lifetraps) == 0 {
				msg.Text = noLifetrap
				msg.ReplyMarkup = backToMainMenuKeyboard
			} else {
				var lifetrapsInText string

				for i, lifetrap := range lifetraps {
					lifetrapsInText += fmt.Sprintf("تله شماره %d:\n%s\n", i+1, lifetrap)
				}

				msg.Text = testEnded + lifetrapsInText
				msg.ReplyMarkup = backToMainMenuKeyboard
			}
		} else {
			msg.Text = question
			msg.ReplyMarkup = scoreButtons(qID+1, categotyID)
		}
	case "myLifeTraps":
		user := users{telegramID: update.CallbackQuery.From.ID}

		testedChan := make(chan bool)
		go func() {
			t := isTestCompleted(user.telegramID)

			testedChan <- t
		}()

		user.tested = <-testedChan

		if user.tested {
			ltChan := make(chan []string)

			go func() {
				lt := showLifetraps(user.telegramID)
				ltChan <- lt
			}()

			lifetraps := <-ltChan

			if len(lifetraps) == 0 {
				msg.Text = noLifetrap
				msg.ReplyMarkup = backToMainMenuKeyboard
			} else {
				var lifetrapsInText string

				for i, lifetrap := range lifetraps {
					lifetrapsInText += fmt.Sprintf("تله شماره %d:\n%s\n", i+1, lifetrap)
				}

				msg.Text = showingLifetraps + lifetrapsInText
				msg.ReplyMarkup = backToMainMenuKeyboard
			}
		} else {
			msg.Text = testNotCompleted
			msg.ReplyMarkup = backToMainMenuKeyboard
		}
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
