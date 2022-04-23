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

//This function returns the number of test questions.
func numberOfQuestions() (count int) {
	db := dbConnect()
	defer db.Close()

	err := db.QueryRow("SELECT COUNT(qid) FROM questions").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count
}

//This function sets a true value for the user when test completed.
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

//This function checks if user already completed the test or not,
//if user already tested it will delete all user scores and updates the tested value to the false.
func resetTest(telegramID int64) {
	db := dbConnect()
	defer db.Close()

	user := users{}
	user.telegramID = telegramID

	testedChan := make(chan bool)
	go func() {
		t := isTestCompleted(user.telegramID)

		testedChan <- t
	}()

	user.tested = <-testedChan

	if user.tested {
		user.tested = false
		stmt, err := db.Prepare("UPDATE users SET tested=? WHERE telegram_id=?")
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmt.Exec(user.tested, user.telegramID)
		if err != nil {
			log.Fatal(err)
		}
	}

	stmt, err := db.Prepare("DELETE FROM scores WHERE telegram_id=?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(user.telegramID)
	if err != nil {
		log.Fatal(err)
	}
}

//This function checks is user completed the test or not and returns a false or true.
func isTestCompleted(telegramID int64) bool {
	db := dbConnect()
	defer db.Close()

	user := users{}
	user.telegramID = telegramID

	err := db.QueryRow("SELECT tested FROM users WHERE telegram_id=?", user.telegramID).Scan(&user.tested)
	if err != nil {
		log.Fatal(err)
	}

	return user.tested
}
