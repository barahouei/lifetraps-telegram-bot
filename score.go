package main

import "log"

//This function adds the scores to the database.
func setScore(telegramID int64, score int, questionID int, categoryID int) {
	db := dbConnect()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO scores(telegram_id, score, qid, cid) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(telegramID, score, questionID, categoryID)
	if err != nil {
		log.Fatal(err)
	}
}
