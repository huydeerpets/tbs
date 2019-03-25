package follows

import "github.com/huydeerpets/tbs/models"

// TruncateTotal 
func TruncateTotal() error {
	u := models.ContributionTotalFollows{}

	return u.Truncate()
}

// AddTotal 
func AddTotal(uID int, count int) error {
	u := models.ContributionTotalFollows{
		UserContributionID: uID,
		Count:              count,
	}

	return u.Add()
}

// AddTotalMap
func AddTotalMap(m map[int]int) error {
	for id, count := range m {
		if err := AddTotal(id, count); err != nil {
			return err
		}
	}

	return nil
}

// GetTotalListByUserContributionIDList
func GetTotalListByUserContributionIDList(idList []int) ([]models.ContributionTotalFollows, error) {
	u := models.ContributionTotalFollows{}

	r, _, err := u.GetListByUserContributionID(idList)
	return r, err
}

// getToatlMap 
func getToatlMap(u []models.ContributionTotalFollows) map[int]int {
	r := map[int]int{}

	for _, v := range u {
		r[v.UserContributionID] = v.Count
	}

	return r
}

// GetTotalMapByUserContributionIDList
func GetTotalMapByUserContributionIDList(idList []int) (map[int]int, error) {
	u, err := GetTotalListByUserContributionIDList(idList)
	if err != nil {
		return map[int]int{}, err
	}

	return getToatlMap(u), nil
}
