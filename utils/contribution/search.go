package contributions

import "github.com/huydeerpets/tbs/models"

// SearchValue 
type SearchValue struct {
	UserContributionID int
	Search             string
	Order              int
}

// SearchWord 
type SearchWord struct {
	Title string
	Body  string
	Tag   string
}

// JoinSearchWord 
func JoinSearchWord(s SearchWord) string {
	return s.Title + "/" + s.Body + "/" + s.Tag
}

// GetSearchWordBody 
func GetSearchWordBody(body string) (s string, err error) {
	b, err := StirngToGetBody(body)
	if err != nil {
		return "", err
	}

	for _, v := range b {
		s += v.Body
	}

	return s, nil
}

// AddSearch 
func AddSearch(uID int, search string) error {
	u := models.UserContributionSearch{
		UserContributionID: uID,
		Search:             search,
	}

	return u.Add()
}

// GetSearchByUserContributionID
func GetSearchByUserContributionID(uID int) (models.UserContributionSearch, error) {
	u := models.UserContributionSearch{}
	r, _, err := u.GetByUserContributionID(uID)

	return r, err
}

// GetSearchListByUserContributionIDList
func GetSearchListByUserContributionIDList(uID []int) ([]models.UserContributionSearch, error) {
	u := models.UserContributionSearch{}
	r, _, err := u.GetListByUserContributionIDList(uID)

	return r, err
}

// AddOrSaveSearch 
func AddOrSaveSearch(uID int, s string) error {
	u, err := GetSearchByUserContributionID(uID)
	if err != nil {
		return err
	}

	if u.ID == uint(0) {
		return AddSearch(uID, s)
	}

	u.Search = s
	return u.Save()
}

// DeleteSearchByUserContributionID
func DeleteSearchByUserContributionID(uID int) error {
	u, err := GetSearchByUserContributionID(uID)
	if err != nil {
		return err
	}

	if u.ID == uint(0) {
		return nil
	}

	return u.Delete()
}

// GetSearchValueListBySearch 
func GetSearchValueListBySearch(search string, order string, limit int, offset int) ([]SearchValue, error) {
	s := []SearchValue{}

	u := models.UserContributionSearch{}
	user, _, err := u.GetListBySearch(search, order, limit, offset)
	if err != nil {
		return s, err
	}

	if len(user) == 0 {
		return s, nil
	}

	for key, v := range user {
		tmp := SearchValue{
			UserContributionID: v.UserContributionID,
			Search:             v.Search,
			Order:              key,
		}

		s = append(s, tmp)
	}

	return s, nil
}

// SaveToFollowCountFollow
func SaveToFollowCount(u []models.UserContributionSearch, m map[int]int) error {
	for _, v := range u {
		if v.FollowCount != m[v.UserContributionID] {
			v.FollowCount = m[v.UserContributionID]
			if err := v.Save(); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetCountBySearch
func GetCountBySearch(search string, order string) (int, error) {
	u := models.UserContributionSearch{}

	return u.GetCountBySearch(search, order)
}

// TruncateSearch 
func TruncateSearch() error {
	u := models.UserContributionSearch{}

	return u.Truncate()
}
