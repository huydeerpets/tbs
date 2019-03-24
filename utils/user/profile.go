package user

import (
	"time"

	"github.com/huydeerpets/tbs/models"
)

// ProfileProfile
type Profile struct {
	ID      uint      `json:"id"`
	UserID  int       `json:"userID"`
	Created time.Time `json:"created"`
}

// GetProfileImageListByUserID ユーザIDからプロフィール画像Listを取得する
func GetProfileImageListByUserID(uID int) ([]Profile, error) {

	profile := []Profile{}
	u := models.UserProfileImage{}

	err := u.GetScanByUserID(uID, &profile)
	if err != nil {
		return profile, err
	}

	return profile, nil
}

// GetIDAndAddProfileImageProfile画像を追加してIDを取得する
func GetIDAndAddProfileImage(uID int) (uint, error) {
	u := models.UserProfileImage{
		UserID: uID,
	}

	return u.GetIDAndAdd()
}
