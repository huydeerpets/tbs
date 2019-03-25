package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// UserContributionSearch 
type UserContributionSearch struct {
	ID                 uint      `gorm:"primary_key"`
	UserContributionID int       `json:"user_contribution_id"`
	Search             string    `gorm:"index:search"`
	FollowCount        int       `json:"follow_count"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

// Add 
func (u *UserContributionSearch) Add() error {
	return Create(u)
}

// Save Save
func (u *UserContributionSearch) Save() error {
	return Save(u)
}

// Delete Delete
func (u *UserContributionSearch) Delete() error {
	return Delete(u)
}

// Truncate 
func (u *UserContributionSearch) Truncate() error {
	return Truncate("user_contribution_searches")
}

// GetByUserContributionID
func (u *UserContributionSearch) GetByUserContributionID(id int) (userContributionSearch UserContributionSearch, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": id},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionSearch, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// GetListByUserContributionIDList
func (u *UserContributionSearch) GetListByUserContributionIDList(id []int) (userContributionSearch []UserContributionSearch, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": id},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionSearch, "User_contribution_ID IN :UserContributionID", whereList, option)

	return
}

// GetListBySearch 
func (u *UserContributionSearch) GetListBySearch(search string, order string, limit int, offset int) (userContributionSearch []UserContributionSearch, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"Search": "+" + search},
	}

	option := map[string]interface{}{
		"order":  order,
		"limit":  limit,
		"offset": offset,
	}

	db, err = GetListWhere(&userContributionSearch, "MATCH(`search`) AGAINST( :Search IN BOOLEAN MODE)", whereList, option)

	return
}

// GetCountBySearch 
func (u *UserContributionSearch) GetCountBySearch(search string, order string) (int, error) {
	userContributionSearch := UserContributionSearch{}

	whereList := []map[string]interface{}{
		{"Search": "+" + search},
	}
	option := make(map[string]interface{})

	return GetCount(&userContributionSearch, "user_contribution_searches", "MATCH(`search`) AGAINST( :Search IN BOOLEAN MODE)", whereList, option)
}
