package question

import "github.com/huydeerpets/tbs/models"

// Add 追加する
func Add(uID int, body string, email string) error {
	log := models.LogQuestion{
		UserID: uID,
		Body:   body,
		Email:  email,
	}

	return log.Add()
}
