package models

import "github.com/jinzhu/gorm"

// UserForgetPassword 
type UserForgetPassword struct {
	BaseModel
	Email   string `json:"email"`
	Keyword string `json:"keyword"`
}

// Add 
func (u *UserForgetPassword) Add() error {
	return Create(u)
}

// Delete Delete
func (u *UserForgetPassword) Delete() error {
	return Delete(u)
}

// DeleteList 
func (u *UserForgetPassword) DeleteList(userForgetPassword []UserForgetPassword) error {
	return Delete(userForgetPassword)
}

// GetByEmail 
func (u *UserForgetPassword) GetByEmail(email string) (userForgetPassword UserForgetPassword, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"Email": email},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userForgetPassword, "Email LIKE :Email", whereList, option)

	return
}

// GetListByEmail 
func (u *UserForgetPassword) GetListByEmail(email string) (userForgetPassword []UserForgetPassword, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"Email": email},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userForgetPassword, "Email LIKE :Email", whereList, option)

	return
}
