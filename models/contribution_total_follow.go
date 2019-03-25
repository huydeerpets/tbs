package models

import "github.com/jinzhu/gorm"

// ContributionTotalFollows PostTotalFollow
type ContributionTotalFollows struct {
	BaseModel
	UserContributionID int `json:"user_contribution_id"`
	Count              int `json:"count"`
}

// Add Add
func (c *ContributionTotalFollows) Add() error {
	return Create(c)
}

// Save Save
func (c *ContributionTotalFollows) Save() error {
	return Save(c)
}

// GetListByUserContributionID Get List from User Post ID
func (c *ContributionTotalFollows) GetListByUserContributionID(uID []int) (contributionTotalFollows []ContributionTotalFollows, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
	}
	option := map[string]interface{}{
		"order": "ID desc",
	}

	db, err = GetListWhere(&contributionTotalFollows, "User_contribution_ID IN :UserContributionID", whereList, option)

	return
}

// Truncate truncate
func (c *ContributionTotalFollows) Truncate() error {
	return Truncate("contribution_total_follows")
}
