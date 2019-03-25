package contributions

import (
	"github.com/huydeerpets/tbs/models"
)

// AddLog 
func AddLog(userID int, uID int) error {
	u := models.LogUserContribution{
		UserID:             userID,
		UserContributionID: uID,
	}

	return u.Add()
}
