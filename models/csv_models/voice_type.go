package csvModels

import "strconv"

// VoiceType voicetype
type VoiceType struct {
	ID              string
	Name            string
	VoiceSystemType string
	VoiceType       string
	Delete          string
}

const (
	// VoiceSystemTypeOpenjtalk Type：open-jtalk
	VoiceSystemTypeOpenjtalk = 1
	// VoiceSystemTypeAquesTalk Type：AquesTalk
	VoiceSystemTypeAquesTalk = 2
)

const (
	// VoiceTypeMeiNormal VoiceType:mei_normal
	VoiceTypeMeiNormal = 1
	// VoiceTypeMeiAngry VoiceType:mei_angry
	VoiceTypeMeiAngry = 2
	// VoiceTypeMeiBashful VoiceType:mei_bashful
	VoiceTypeMeiBashful = 3
	// VoiceTypeMeiHappy VoiceType:mei_happy
	VoiceTypeMeiHappy = 4
	// VoiceTypeMeiSad VoiceType:mei_sad
	VoiceTypeMeiSad = 5
	// VoiceTypeM100 VoiceType:m100
	VoiceTypeM100 = 6
	// VoiceTypeYukkuri VoiceType:slowly
	VoiceTypeYukkuri = 7
)

// GetStructAll Get All
func (c *VoiceType) GetStructAll() (r []VoiceType, err error) {
	err = GetAll("voice_type.csv", &r)

	return r, err
}

// GetStruct Get
func (c *VoiceType) GetStruct(voiceType int) (r VoiceType, err error) {
	list, err := c.GetStructAll()
	if err != nil {
		return r, err
	}

	for _, v := range list {
		id, err := strconv.Atoi(v.ID)
		if err != nil {
			return r, err
		}

		if id == voiceType {
			return v, nil
		}
	}

	return r, nil

}
