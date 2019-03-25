package models

import "github.com/jinzhu/gorm"

// UserContributionTag 
type UserContributionTag struct {
	BaseModel
	UserContributionID int    `json:"user_contribution_id"`
	Name               string `json:"name"`
}

// Add
func (uc *UserContributionTag) Add() (err error) {
	return Create(uc)
}

// AddList 
func (uc *UserContributionTag) AddList(u []UserContributionTag) (err error) {
	for _, user := range u {
		if err = Create(&user); err != nil {
			return err
		}
	}

	return nil
}

// GetListByUserContributionID
func (uc *UserContributionTag) GetListByUserContributionID(id int) (userContributionTag []UserContributionTag, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": id},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionTag, "User_contribution_ID = :UserContributionID", whereList, option)

	return
}

// GetScanListByUserContributionID
func (uc *UserContributionTag) GetScanListByUserContributionID(id int, dest interface{}) error {
	whereList := []map[string]interface{}{
		{"UserContributionID": id},
	}
	option := make(map[string]interface{})

	return GeScanWhere(dest, "user_contribution_tags", "User_contribution_ID = :UserContributionID", whereList, option)
}

// GetListByUserContributionIDList
func (uc *UserContributionTag) GetListByUserContributionIDList(idList []int) (userContributionTag []UserContributionTag, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": idList},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionTag, "User_contribution_ID IN :UserContributionID", whereList, option)

	return
}

// GetScanListByUserContributionIDList
func (uc *UserContributionTag) GetScanListByUserContributionIDList(idList []int, dest interface{}) error {
	whereList := []map[string]interface{}{
		{"UserContributionID": idList},
	}
	option := make(map[string]interface{})

	return GeScanWhere(dest, "user_contribution_tags", "User_contribution_ID IN :UserContributionID", whereList, option)
}

// GetByID 
func (uc *UserContributionTag) GetByID(id int) (userContributionTag UserContributionTag, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionTag, "ID = :ID", whereList, option)

	return
}

// Save Save
func (uc *UserContributionTag) Save() error {
	return Save(uc)
}

// Delete Delete
func (uc *UserContributionTag) Delete() error {
	return Delete(uc)
}
