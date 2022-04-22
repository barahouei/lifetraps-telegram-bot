package main

import "log"

//This function takes a question ID and returns the question and category ID it belongs to.
func questionWalker(questionID int) (question string, categoryID int) {
	db := dbConnect()
	defer db.Close()

	err := db.QueryRow("SELECT question, cid FROM questions WHERE qid=?", questionID).Scan(&question, &categoryID)
	if err != nil {
		log.Fatal(err)
	}

	return question, categoryID
}
