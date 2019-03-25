package models

import "github.com/jinzhu/gorm"

// UserCharacterImage User character image
type UserCharacterImage struct {
	BaseModel
	UserID      int `json:"user_id"`
	CharacterID int `json:"character_id"`
	Priority    int `json:"priority"`
	VoiceType   int `json:"voice_type"`
}

// Add Add
func (u *UserCharacterImage) Add() error {
	return Create(u)
}

// Save Save
func (u *UserCharacterImage) Save() error {
	return Save(u)
}

// GetListByUserID Get List from User ID
func (u *UserCharacterImage) GetListByUserID(uID int) (userCharacterImage []UserCharacterImage, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserID": uID},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userCharacterImage, "User_ID = :UserID", whereList, option)

	return
}

// GetByID Get from ID
func (u *UserCharacterImage) GetByID(id int) (userCharacterImage UserCharacterImage, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userCharacterImage, "ID = :ID", whereList, option)

	return
}

// Delete Delete
func (u *UserCharacterImage) Delete() error {
	return Delete(u)
}
