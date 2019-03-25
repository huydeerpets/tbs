package models

import "github.com/jinzhu/gorm"

// UserMaster User information
type UserMaster struct {
	BaseModel      `model:"true"`
	Name           string `json:"name"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password"`
	ProfileImageID int    `json:"profile_image_id"`
}

// GetIDAndAddPost to get the ID
func (u *UserMaster) GetIDAndAdd() (uint, error) {
	if err := Create(u); err != nil {
		return 0, err
	}

	return u.ID, nil
}

// Save Save
func (u *UserMaster) Save() error {
	return Save(u)
}

// GetByEmail 
func (u *UserMaster) GetByEmail(email string) (userMaster UserMaster, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"Email": email},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userMaster, "Email LIKE :Email", whereList, option)

	return
}

// GetByID Get from user ID
func (u *UserMaster) GetByID(id int) (userMaster UserMaster, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userMaster, "ID = :ID", whereList, option)

	return
}

// GetScanByID Get scan from user ID
func (u *UserMaster) GetScanByID(id int, dest interface{}) error {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	return GeScanWhere(dest, "user_masters", "ID = :ID", whereList, option)
}

// GetListByIDList Get List from User ID List
func (u *UserMaster) GetListByIDList(idList []int) (userMaster []UserMaster, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"ID": idList},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userMaster, "ID IN :ID", whereList, option)

	return
}

// GetScanByIDList Get scan from user ID List
func (u *UserMaster) GetScanByIDList(idList []int, dest interface{}) error {
	whereList := []map[string]interface{}{
		{"ID": idList},
	}
	option := make(map[string]interface{})

	return GeScanWhere(dest, "user_masters", "ID IN :ID", whereList, option)
}
