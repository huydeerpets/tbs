package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

// ViewStatusPublic 公開状態
const ViewStatusPublic = 1

// ViewStatusPrivate プライベート状態
const ViewStatusPrivate = 2

// UserContribution ユーザ投稿
type UserContribution struct {
	BaseModel
	UserID     int    `json:"user_id"`
	Title      string `json:"title"`
	ViewStatus int    `json:"view_status"`
}

// GetIDAndAddPostしてIDを取得する
func (u *UserContribution) GetIDAndAdd() (uint, error) {
	if err := Create(u); err != nil {
		return 0, err
	}

	return u.ID, nil
}

// Save Save
func (u *UserContribution) Save() error {
	return Save(u)
}

// Delete Deleteする
func (u *UserContribution) Delete() error {
	return Delete(u)
}

// GetByIDPostIDから取得する
func (u *UserContribution) GetByID(id int) (userContribution UserContribution, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"ID": id},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContribution, "ID = :ID", whereList, option)
	return
}

// GetListByUserIDPostIDからList取得する
func (u *UserContribution) GetListByUserID(userID int, order string, limit int, offset int) (userContribution []UserContribution, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserID": userID},
	}

	option := map[string]interface{}{
		"order":  order,
		"limit":  limit,
		"offset": offset,
	}

	db, err = GetListWhere(&userContribution, "User_ID = :UserID", whereList, option)
	return
}

// GetCountByUserIDPostIDから数を取得する
func (u *UserContribution) GetCountByUserID(userID int, order string) (int, error) {
	userContribution := UserContribution{}

	whereList := []map[string]interface{}{
		{"UserID": userID},
	}

	option := make(map[string]interface{})

	return GetCount(&userContribution, "user_contributions", "User_ID = :UserID AND Deleted_at IS NULL", whereList, option)
}

// GetByTop 新着から投稿List取得する
func (u *UserContribution) GetByTop(o int, s int) (userContributionList []UserContribution, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{}

	optionMap := map[string]interface{}{
		"order":  "ID desc",
		"limit":  s,
		"offset": o,
	}

	db, err = GetListWhere(&userContributionList, "View_status = "+strconv.Itoa(ViewStatusPublic), whereList, optionMap)
	return
}

// GetListByIDList IDListから投稿Listを取得する
func (u *UserContribution) GetListByIDList(idList []int) (userContributionList []UserContribution, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"IDList": idList},
	}
	optionMap := make(map[string]interface{})

	db, err = GetListWhere(&userContributionList, "ID IN :IDList", whereList, optionMap)
	return
}

// GetListByViewStatusPublic IDListから投稿Listを取得する
func (u *UserContribution) GetListByViewStatusPublic() (userContributionList []UserContribution, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{}

	optionMap := map[string]interface{}{
		"select": "id",
	}

	db, err = GetListWhere(&userContributionList, "View_status = "+strconv.Itoa(ViewStatusPublic), whereList, optionMap)
	return
}
