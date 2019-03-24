package bug

import "github.com/huydeerpets/tbs/models"

// Add 追加する
func Add(uID int, body string) error {
	log := models.LogBugReport{
		UserID: uID,
		Body:   body,
	}

	return log.Add()
}
