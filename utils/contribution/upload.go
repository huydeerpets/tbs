package contributions

import (
	"github.com/huydeerpets/tbs/models"
)

// AddUpload 
func AddUpload(uID int, t string) error {
	u := models.UserContributionUpload{
		UserContributionID: uID,
		Token:              t,
	}

	return u.Add()
}

// GetUploadByUserContributionID
func GetUploadByUserContributionID(uID int) (models.UserContributionUpload, error) {
	u := models.UserContributionUpload{}

	r, _, err := u.GetByUserContributionID(uID)

	return r, err
}
