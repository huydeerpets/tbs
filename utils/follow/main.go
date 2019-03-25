package follows

import "github.com/huydeerpets/tbs/models"

// OrderValue 
type OrderValue struct {
	UserContributionID int
	Order              int
}

// Add 
func Add(uID int, cID int) error {
	u := models.UserContributionFollow{
		UserID:             uID,
		UserContributionID: cID,
	}

	return u.Add()
}

// DeleteFollow
func Delete(id uint) error {
	u := models.UserContributionFollow{}
	user, _, err := u.GetByID(id)
	if err != nil {
		return err
	}

	return user.Delete()
}

// GetListByUserContributionID
func GetListByUserContributionID(cID int) ([]models.UserContributionFollow, error) {
	u := models.UserContributionFollow{}
	r, _, err := u.GetListByUserContributionID(cID)

	return r, err
}

// GetCountByUserContributionID
func GetCountByUserContributionID(cID int) (int, error) {
	u := models.UserContributionFollow{}

	return u.GetCountByUserContributionID(cID)
}

// GetByUserIDAndUserContributionID 
func GetByUserIDAndUserContributionID(uID int, ucID int) (models.UserContributionFollow, error) {
	u := models.UserContributionFollow{}
	r, _, err := u.GetByUserIDAndUserContributionID(uID, ucID)

	return r, err
}

// GetCountByUserIDAndUserContributionID 
func GetCountByUserIDAndUserContributionID(uID int, ucID int) (int, error) {
	u := models.UserContributionFollow{}

	return u.GetCountByUserIDAndUserContributionID(uID, ucID)
}

// GetListByUserID 
func GetListByUserID(uID int, order string, limit int, offset int) ([]models.UserContributionFollow, error) {
	u := models.UserContributionFollow{}
	r, _, err := u.GetListByUserID(uID, order, limit, offset)

	return r, err
}

// GetOrderValueListByUserID 
func GetOrderValueListByUserID(uID int, order string, limit int, offset int) (o []OrderValue, err error) {
	u, err := GetListByUserID(uID, order, limit, offset)
	if err != nil {
		return o, err
	}

	if len(u) == 0 {
		return o, nil
	}

	for key, v := range u {
		tmp := OrderValue{
			UserContributionID: v.UserContributionID,
			Order:              key,
		}

		o = append(o, tmp)
	}

	return o, nil
}

// GetListByUserContributionIDList
func GetListByUserContributionIDList(cID []int) ([]models.UserContributionFollow, error) {
	u := models.UserContributionFollow{}
	r, _, err := u.GetListByUserContributionIDList(cID)

	return r, err
}

// GetFollowCountMapFollow
func GetFollowCountMap(u []models.UserContributionFollow) map[int]int {
	m := map[int]int{}

	for _, v := range u {
		id := int(v.UserContributionID)
		if _, ok := m[id]; !ok {
			m[id] = 0
		}

		m[id]++
	}

	return m
}

// GetCountByUserID 
func GetCountByUserID(uID int, order string) (int, error) {
	u := models.UserContributionFollow{}

	return u.GetCountByUserID(uID, order)
}
