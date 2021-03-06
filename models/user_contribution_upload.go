package models

import (
	"github.com/jinzhu/gorm"
)

// UserContributionUpload 
type UserContributionUpload struct {
	BaseModel
	UserID             int    `json:"user_id"`
	UserContributionID int    `json:"user_contribution_id"`
	Token              string `json:"token"`
}

// Add 
func (u *UserContributionUpload) Add() error {
	return Create(u)
}

// Save Save
func (u *UserContributionUpload) Save() error {
	return Save(u)
}

// GetByUserContributionID
func (u *UserContributionUpload) GetByUserContributionID(uID int) (userContributionUpload UserContributionUpload, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionUpload, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}
