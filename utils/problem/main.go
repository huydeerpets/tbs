package problem

import "github.com/huydeerpets/tbs/models"

// Add 
func Add(userID int, uID int, programType int) error {
	log := models.LogProblemContributionReport{
		UserID:             userID,
		UserContributionID: uID,
		Type:               programType,
	}

	return log.Add()
}
