package characters

import (
	"errors"
	"strconv"

	"github.com/huydeerpets/tbs/models"
)

// ImageImage
type Image struct {
	ID          uint   `json:"id"`
	CharacterID int    `json:"characterID"`
	Priority    int    `json:"priority"`
	VoiceType   int    `json:"voiceType"`
	FileName    string `json:"fileName"`
}

// AddImageImage
func AddImage(uID int, cID int, p int) (uint, error) {
	u := &models.UserCharacterImage{
		UserID:      uID,
		CharacterID: cID,
		Priority:    p,
	}

	e := u.Add()

	return u.ID, e
}

// SaveToVoiceType 
func SaveToVoiceType(id int, v int, userID int) error {
	u := models.UserCharacterImage{}

	user, _, err := u.GetByID(id)
	if err != nil {
		return err
	}

	if userID != user.UserID {
		return errors.New("User_ID is wrong")
	}

	user.VoiceType = v

	return user.Save()
}

// GetImageListByUserID 
func GetImageListByUserID(uID int) ([]Image, error) {
	u := models.UserCharacterImage{}
	image := []Image{}
	r, _, err := u.GetListByUserID(uID)

	for _, v := range r {
		st := Image{
			ID:          v.ID,
			CharacterID: v.CharacterID,
			Priority:    v.Priority,
			VoiceType:   v.VoiceType,
			FileName:    GetImageName(v.ID),
		}

		image = append(image, st)
	}

	return image, err
}

// DeleteByID 
func DeleteByID(id int, userID int) error {
	u := models.UserCharacterImage{}
	userCharacterImage, _, err := u.GetByID(id)
	if err != nil {
		return err
	}

	if userCharacterImage.UserID != userID {
		return errors.New("User_ID is wrong")
	}

	return userCharacterImage.Delete()
}

// GetImageName
func GetImageName(id uint) string {
	return strconv.Itoa(int(id)) + ".jpg"
}
