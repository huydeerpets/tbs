package contributions

import (
	"encoding/json"

	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils/character"
)

// SaveBody 
type SaveBody struct {
	Priority      int           `json:"priority"`
	Body          string        `json:"body"`
	DirectionType int           `json:"directionType"`
	TalkType      int           `json:"talkType"`
	Character     SaveCharacter `json:"character"`
}

// SaveCharacter 
type SaveCharacter struct {
	ID        int    `json:"id"`
	FileName  string `json:"fileName"`
	VoiceType int    `json:"voiceType"`
}

// GetBody 
type GetBody struct {
	Priority      int          `json:"priority"`
	Body          string       `json:"body"`
	DirectionType int          `json:"directionType"`
	TalkType      int          `json:"talkType"`
	Character     GetCharacter `json:"character"`
}

// GetCharacter 
type GetCharacter struct {
	ID        int    `json:"id"`
	FileName  string `json:"fileName"`
	VoiceType int    `json:"voiceType"`
}

// SaveDetail 
func SaveDetail(userContributionID int, body string) error {
	ucd, err := GetDetailByUserContributionID(userContributionID)
	if err != nil {
		return err
	}

	b, err := StirngToSaveBody(body)
	if err != nil {
		return err
	}

	st, err := json.Marshal(b)
	if err != nil {
		return err
	}

	ucd.Body = string(st)

	ucd.Save()

	return nil
}

// StirngToSaveBody 
func StirngToSaveBody(body string) (b []SaveBody, err error) {
	bytes := []byte(body)
	err = json.Unmarshal(bytes, &b)
	if err != nil {
		return b, err
	}

	for k := range b {
		b[k].Priority = k
	}

	return b, err
}

// StirngToGetBody 
func StirngToGetBody(body string) (b []GetBody, err error) {
	bytes := []byte(body)
	err = json.Unmarshal(bytes, &b)
	if err != nil {
		return b, err
	}

	for k, v := range b {
		if v.Character.ID == 0 {
			b[k].Character.FileName = v.Character.FileName
		} else {
			b[k].Character.FileName = characters.GetImageName(uint(v.Character.ID))
		}
	}

	return b, err
}

// GetDetailByUserContributionID
func GetDetailByUserContributionID(uID int) (models.UserContributionDetail, error) {
	u := &models.UserContributionDetail{}

	r, _, err := u.GetByUserContributionID(uID)
	if err != nil {
		return r, err
	}

	return r, err
}

// GetBodyByUserContributionID
func GetBodyByUserContributionID(uID int) ([]GetBody, error) {
	b := []GetBody{}
	u, err := GetDetailByUserContributionID(uID)
	if err != nil {
		return b, err
	}

	b, err = StirngToGetBody(u.Body)
	if err != nil {
		return b, err
	}

	return b, nil
}
