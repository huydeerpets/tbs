package contributions

import (
	"github.com/huydeerpets/tbs/models"
)

// AddUpload uploadを追加する
func AddUpload(uID int, t string) error {
	u := models.UserContributionUpload{
		UserContributionID: uID,
		Token:              t,
	}

	return u.Add()
}

// GetUploadByUserContributionIDPostIDからuploadを情報を取得する
func GetUploadByUserContributionID(uID int) (models.UserContributionUpload, error) {
	u := models.UserContributionUpload{}

	r, _, err := u.GetByUserContributionID(uID)

	return r, err
}
