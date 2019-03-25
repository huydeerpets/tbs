package contributions

import "github.com/huydeerpets/tbs/models"

// GetImageIDAndAdd 
func GetImageIDAndAdd(userContributionID int) (uint, error) {
	l := models.LogContributionImage{
		UserContributionID: userContributionID,
	}

	return l.GetIDAndAdd()
}
