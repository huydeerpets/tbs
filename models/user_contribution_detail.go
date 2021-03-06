package models

import "github.com/jinzhu/gorm"

// UserContributionDetail
type UserContributionDetail struct {
	BaseModel
	UserContributionID int    `json:"user_contribution_id"`
	Body               string `json:"body"`
}

// Add
func (u *UserContributionDetail) Add() error {
	return Create(u)
}

// Save Save
func (u *UserContributionDetail) Save() error {
	return Save(u)
}

// Delete Delete
func (u *UserContributionDetail) Delete() error {
	return Delete(u)
}

// GetByUserContributionID Get from PostID
func (u *UserContributionDetail) GetByUserContributionID(uID int) (userContributionDetail UserContributionDetail, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionDetail, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}
