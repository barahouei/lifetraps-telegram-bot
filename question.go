package main

import "log"

//This function takes a question ID and returns the question and category ID it belongs to,
//and if there were an error it returns the error too.
func questionWalker(questionID int) (question string, categoryID int, err error) {
	db := dbConnect()
	defer db.Close()

	err = db.QueryRow("SELECT question, cid FROM questions WHERE qid=?", questionID).Scan(&question, &categoryID)

	return question, categoryID, err
}

//If user completed the test this function sets a true value for the user.
func setTestCompleted(telegramID int64) {
	db := dbConnect()
	defer db.Close()

	user := users{}
	user.telegramID = telegramID
	user.tested = true

	stmt, err := db.Prepare("UPDATE users SET tested=? WHERE telegram_id=?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(user.tested, user.telegramID)
	if err != nil {
		log.Fatal(err)
	}
}
