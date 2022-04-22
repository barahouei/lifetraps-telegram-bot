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

//This function returns all lifetraps which user has.
func showLifetraps(telegramID int64) []string {
	lt := maxScore(telegramID)

	lifetraps := []string{}

	for categoryID, maxScore := range lt {
		if maxScore > 3 {
			db := dbConnect()
			defer db.Close()

			var category string
			err := db.QueryRow("SELECT category FROM categories WHERE cid=?", categoryID).Scan(&category)
			if err != nil {
				log.Fatal(err)
			}

			lifetraps = append(lifetraps, category)
		}
	}

	return lifetraps
}

//This function gets max score of each lifetrap and returns a int map
//which key is the category ID of lifetrap and value is the max score of that lifetrap.
func maxScore(telegramID int64) map[int]int {
	db := dbConnect()
	defer db.Close()

	liftraps := make(map[int]int)
	for i := 1; i <= 2; i++ {
		var maxScore, categoryID int

		err := db.QueryRow("SELECT MAX(score), cid FROM scores WHERE telegram_id=? AND cid=?", telegramID, i).Scan(&maxScore, &categoryID)
		if err != nil {
			log.Fatal(err)
		}

		liftraps[categoryID] = maxScore
	}

	return liftraps
}
