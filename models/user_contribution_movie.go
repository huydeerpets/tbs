package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

const (
	// MovieTypeYoutube 
	MovieTypeYoutube = 1
)

// UserContributionMovie 
type UserContributionMovie struct {
	BaseModel
	UserContributionID int    `json:"user_contribution_id"`
	MovieType          int    `json:"movie_type"`
	MovieID            string `json:"movie_id"`
	MovieStatus        int    `json:"movie_status"`
}

// Add
func (u *UserContributionMovie) Add() error {
	return Create(u)
}

// Save Save
func (u *UserContributionMovie) Save() error {
	return Save(u)
}

// GetByUserContributionID
func (u *UserContributionMovie) GetByUserContributionID(uID int, t int) (userContributionMovie UserContributionMovie, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
		{"MovieType": t},
	}
	option := make(map[string]interface{})

	db, err = GetWhere(&userContributionMovie, "User_contribution_ID = :UserContributionID AND movie_type = :MovieType", whereList, option)

	return
}

// GetListByUserContributionIDList
func (u *UserContributionMovie) GetListByUserContributionIDList(uID []int, t int) (userContributionMovie []UserContributionMovie, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{
		{"UserContributionID": uID},
		{"MovieType": t},
	}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionMovie, "User_contribution_ID IN :UserContributionID AND movie_type = :MovieType AND movie_status = "+strconv.Itoa(StatusPublic), whereList, option)

	return
}

// GetListByMovieStatusPublic 
func (u *UserContributionMovie) GetListByMovieStatusPublic() (userContributionMovie []UserContributionMovie, db *gorm.DB, err error) {
	whereList := []map[string]interface{}{}
	option := make(map[string]interface{})

	db, err = GetListWhere(&userContributionMovie, "movie_status = "+strconv.Itoa(StatusPublic), whereList, option)

	return
}
