package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

const (
	// TalkTypeText 
	TalkTypeText = 1
	// TalkTypeImage 
	TalkTypeImage = 2
	// MakeStatusUncreated 
	MakeStatusUncreated = 2
	// MakeStatusMade
	MakeStatusMade = 1
)

// UserContributionSoundDetail 
type UserContributionSoundDetail struct {
	BaseModel
	UserContributionID int    `json:"user_contribution_id"`
	Priority           int    `json:"priority"`
	TalkType           int    `json:"talk_type"`
	Body               string `json:"body"`
	BodySound          string `json:"body_sound"`
	VoiceType          int    `json:"voice_type"`
	MakeStatus         int    `json:"make_status"`
}

// Add 
func (u *UserContributionSoundDetail) Add() error {
	return Create(u)
}

// Save Save
func (u *UserContributionSoundDetail) Save() error {
	return Save(u)
}

// GetListByUserContributionID
func (u *UserContributionSoundDetail) GetListByUserContributionID(uID int) (userContributionSoundDetail []UserContributionSoundDetail, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionSoundDetail, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// UpdateToMakeStatusByUserContributionID
func (u *UserContributionSoundDetail) UpdateToMakeStatusByUserContributionID(uID int, makeStatus int) (err error) {
	userContributionSoundDetail := []UserContributionSoundDetail{}

	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})
	update := []interface{}{
		"make_status",
		makeStatus,
	}

	_, err = Update(&userContributionSoundDetail, update, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// UpdatesToMakeStatusAndVoiceTypeByUserContributionID
func (u *UserContributionSoundDetail) UpdatesToMakeStatusAndVoiceTypeByUserContributionID(uID int, makeStatus int, voiceType int) (err error) {
	userContributionSoundDetail := []UserContributionSoundDetail{}

	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}

	option := make(map[string]interface{})
	update := UserContributionSoundDetail{
		MakeStatus: makeStatus,
		VoiceType:  voiceType,
	}

	_, err = Updates(&userContributionSoundDetail, update, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// GetByID 
func (u *UserContributionSoundDetail) GetByID(id uint) (userContributionSoundDetail UserContributionSoundDetail, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionSoundDetail, "ID = :ID", whereList, option)

	return
}

// GetListByMakeStatusMade 
func (u *UserContributionSoundDetail) GetListByMakeStatusMade() (userContributionSoundDetail []UserContributionSoundDetail, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionSoundDetail, "Make_status = "+strconv.Itoa(MakeStatusMade), whereList, option)

	return
}
