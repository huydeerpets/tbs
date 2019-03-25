package models

import (
	"github.com/jinzhu/gorm"
)

// UserContributionSoundLength 
type UserContributionSoundLength struct {
	BaseModel
	UserContributionID int `json:"user_contribution_id"`
	Second             int `json:"second"`
	Length             int `json:"length"`
}

// Add 
func (u *UserContributionSoundLength) Add() error {
	return Create(u)
}

// Save Save
func (u *UserContributionSoundLength) Save() error {
	return Save(u)
}

// GetByUserContributionID
func (u *UserContributionSoundLength) GetByUserContributionID(uID int) (userContributionSoundLength UserContributionSoundLength, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionSoundLength, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// GetByTop 
func (u *UserContributionSoundLength) GetByTop(o int, s int) (userContributionSoundLength []UserContributionSoundLength, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{}

	optionMap := map[string]interface{}{
		"order":  "ID desc",
		"limit":  s,
		"offset": o,
	}

	db, err = GetListWhere(&userContributionSoundLength, "", whereList, optionMap)
	return
}
