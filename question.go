package main

//This function takes a question ID and returns the question and category ID it belongs to,
//and if there were an error it returns the error too.
func questionWalker(questionID int) (question string, categoryID int, err error) {
	db := dbConnect()
	defer db.Close()

	err = db.QueryRow("SELECT question, cid FROM questions WHERE qid=?", questionID).Scan(&question, &categoryID)

	return question, categoryID, err
}
