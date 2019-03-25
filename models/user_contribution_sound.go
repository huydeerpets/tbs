package models

import (
	"github.com/jinzhu/gorm"
)

const (
	// SoundStatusPublic 
	SoundStatusPublic = 1
	// SoundStatusPrivate 
	SoundStatusPrivate = 2
)

// UserContributionSound 
type UserContributionSound struct {
	BaseModel
	UserContributionID int `json:"user_contribution_id"`
	SoundStatus        int `json:"sound_status"`
}

// Add 
func (u *UserContributionSound) Add() error {
	return Create(u)
}

// Save Save
func (u *UserContributionSound) Save() error {
	return Save(u)
}

// GetByUserContributionID
func (u *UserContributionSound) GetByUserContributionID(uID int) (userContributionSound UserContributionSound, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionSound, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// GetListByUserContributionIDList
func (u *UserContributionSound) GetListByUserContributionIDList(uID []int) (userContributionSound []UserContributionSound, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionSound, "User_contribution_ID IN :UserContributionID", whereList, option)

	return
}
