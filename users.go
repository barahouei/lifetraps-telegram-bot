package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type users struct {
	telegramID int64
	username   string
	firstname  string
	lastname   string
	tested     bool
}

//This function check if user already existed or not,
//and if user can not be found in the database it add the user to database.
func checkUserExists(update tgbotapi.Update) {
	user := users{}
	var userID int64
	var existed bool

	user.telegramID = update.Message.From.ID
	user.username = update.Message.From.UserName
	user.firstname = update.Message.From.FirstName
	user.lastname = update.Message.From.LastName
	user.tested = false

	db := dbConnect()
	defer db.Close()

	err := db.QueryRow("SELECT telegram_id FROM users WHERE telegram_id=?", user.telegramID).Scan(&userID)
	if err != nil {
		existed = false
	} else {
		existed = true
	}

	if !existed {
		stmt, err := db.Prepare("INSERT INTO users(telegram_id, username, firstname, lastname, tested) VALUES(?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmt.Exec(user.telegramID, user.username, user.firstname, user.lastname, user.tested)
		if err != nil {
			log.Fatal(err)
		}
	}
}
