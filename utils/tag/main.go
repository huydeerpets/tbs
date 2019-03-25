package tags

import (
	"errors"
	"strings"

	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils"
)

// TagMaxNumber 
const TagMaxNumber = 10

// Tag 
type Tag struct {
	ID                 uint   `json:"id"`
	UserContributionID int    `json:"userContributionID"`
	Name               string `json:"name"`
}

// Save Save
func Save(id int, n string) (err error) {
	u := models.UserContributionTag{}
	u, _, err = u.GetByID(id)
	if err != nil {
		return err
	}

	u.Name = n

	return u.Save()
}

// DeleteByIDAndUserContributionID
func DeleteByIDAndUserContributionID(id int, cID int) (err error) {
	u := models.UserContributionTag{}
	u, _, err = u.GetByID(id)

	if u.UserContributionID != cID {
		return errors.New("difference UserContributionID")
	}

	if err != nil {
		return err
	}

	return u.Delete()
}

// Add
func Add(uID int, n string) error {
	u := models.UserContributionTag{
		UserContributionID: uID,
		Name:               n,
	}

	return u.Add()
}

// AddList 
func AddList(uID int, n string) error {
	n = strings.Replace(strings.TrimSpace(n), "　", " ", -1)
	namelist := strings.Split(n, " ")

	userContributionTag := []models.UserContributionTag{}

	if len(namelist) > TagMaxNumber {
		return errors.New("max number over tag")
	}

	addName := []string{}

	for _, name := range namelist {

		if utils.InStringArray(name, addName) {
			continue
		}

		if len(name) > 20 {
			continue
		}

		u := models.UserContributionTag{
			UserContributionID: uID,
			Name:               name,
		}

		addName = append(addName, name)

		userContributionTag = append(userContributionTag, u)
	}

	uct := models.UserContributionTag{}

	return uct.AddList(userContributionTag)
}

// GetListByUserContributionID
func GetListByUserContributionID(uID int) ([]Tag, error) {
	u := &models.UserContributionTag{}
	tag := []Tag{}

	err := u.GetScanListByUserContributionID(uID, &tag)
	if err != nil {
		return tag, err
	}

	if len(tag) == 0 {
		tag = []Tag{}
	}

	return tag, nil
}

// GetMapByUserContributionIDList
func GetMapByUserContributionIDList(uIDList []int) (map[int][]Tag, error) {
	tagMap := map[int][]Tag{}

	u := &models.UserContributionTag{}
	tagList := []Tag{}
	err := u.GetScanListByUserContributionIDList(uIDList, &tagList)
	if err != nil {
		return tagMap, err
	}

	for _, tag := range tagList {
		tagMap[tag.UserContributionID] = append(tagMap[tag.UserContributionID], tag)
	}
	return tagMap, nil
}

// GetTagNameJoin 
func GetTagNameJoin(uID int) (string, error) {
	t, err := GetListByUserContributionID(uID)
	if err != nil {
		return "", err
	}

	return ToTagNameJoin(t), nil
}

// ToTagNameJoin 
func ToTagNameJoin(t []Tag) string {
	list := []string{}
	for _, v := range t {
		list = append(list, v.Name)
	}

	return strings.Join(list, ",")
}
