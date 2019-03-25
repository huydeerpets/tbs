package user

import (
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils"
	"time"
)

// PasswordMinLength 
const PasswordMinLength = 8

// AddForgetPassword 
func AddForgetPassword(email string, keyword string) error {
	u := models.UserForgetPassword{
		Email:   email,
		Keyword: keyword,
	}

	return u.Add()
}

// GetForgetPasswordByEmail 
func GetForgetPasswordByEmail(email string) (models.UserForgetPassword, error) {
	u := models.UserForgetPassword{}
	r, _, err := u.GetByEmail(email)

	return r, err
}

// IsUpdatePassword 
func IsUpdatePassword(email string, keyword string) (bool, error) {
	up, err := GetForgetPasswordByEmail(email)
	if err != nil {
		return false, err
	}

	if up.Keyword != keyword {
		return false, nil
	}

	if up.CreatedAt.Add(1*time.Hour).Unix() < utils.Now().Unix() {
		return false, nil
	}

	return true, nil
}

// DeleteByEmail 
func DeleteByEmail(email string) error {
	u := models.UserForgetPassword{}
	r, _, err := u.GetListByEmail(email)
	if err != nil {
		return err
	}

	if len(r) == 0 {
		return nil
	}

	return u.DeleteList(r)
}
