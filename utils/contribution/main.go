package contributions

import (
	"errors"
	"time"

	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils/follow"
	"github.com/huydeerpets/tbs/utils/tag"
	"github.com/huydeerpets/tbs/utils/user"
)

// ContributionPost
type Contribution struct {
	ID          uint                         `json:"id"`
	User        user.User                    `json:"user"`
	Title       string                       `json:"title"`
	Tag         []tags.Tag                   `json:"tags"`
	FollowCount int                          `json:"followCount"`
	Body        []GetBody                    `json:"body"`
	ViewStatus  int                          `json:"viewStatus"`
	Search      string                       `json:"search"`
	SoundStatus int                          `json:"soundStatus"`
	Movie       models.UserContributionMovie `json:"movie"`
	UpdatedAt   time.Time                    `json:"updatedAt"`
	CreatedAt   time.Time                    `json:"createdAt"`
}

// AddPost
func Add(userID int, title string, body string, v int) (uint, error) {
	userContribution := &models.UserContribution{
		UserID:     userID,
		Title:      title,
		ViewStatus: v,
	}

	userContributionID, err := userContribution.GetIDAndAdd()
	if err != nil {
		return 0, err
	}

	userContributionDetail := &models.UserContributionDetail{
		UserContributionID: int(userContributionID),
		Body:               body,
	}
	userContributionDetail.Add()

	return userContributionID, nil
}

// Save Save
func Save(userContributionID int, userID int, title string, v int) error {
	u, err := GetByUserContributionID(userContributionID)
	if err != nil {
		return err
	}

	if u.UserID != userID {
		return errors.New("difference UserID")
	}

	u.Title = title
	u.ViewStatus = v

	return u.Save()
}

// DeleteByID Delete
func DeleteByID(userContributionID int, userID int) error {
	uc, err := GetByUserContributionID(userContributionID)
	if err != nil {
		return err
	}

	if uc.UserID != userID {
		return errors.New("difference UserID")
	}

	if e := uc.Delete(); e != nil {
		return e
	}

	ucd, err := GetDetailByUserContributionID(userContributionID)
	if err != nil {
		return err
	}

	return ucd.Delete()
}

// GetByUserContributionID
func GetByUserContributionID(userContributionID int) (models.UserContribution, error) {
	userContribution := &models.UserContribution{}

	r, _, err := userContribution.GetByID(userContributionID)

	return r, err
}

// GetListByUserID 
func GetListByUserID(userID int, order string, limit int, offset int) ([]models.UserContribution, error) {
	u := &models.UserContribution{}

	r, _, err := u.GetListByUserID(userID, order, limit, offset)

	return r, err
}

// GetCountByUserID 
func GetCountByUserID(userID int, order string) (int, error) {
	u := &models.UserContribution{}

	return u.GetCountByUserID(userID, order)
}

// GetContributionByUserContributionID
func GetContributionByUserContributionID(userContributionID int) (c Contribution, err error) {
	uc, err := GetByUserContributionID(userContributionID)
	if err != nil {
		return c, err
	}
	var u user.User

	if u, err = user.GetByUserID(uc.UserID); err != nil {
		return c, err
	}

	var tag []tags.Tag
	if tag, err = tags.GetListByUserContributionID(userContributionID); err != nil {
		return c, err
	}

	var body []GetBody
	if body, err = GetBodyByUserContributionID(userContributionID); err != nil {
		return c, err
	}

	user := user.User{
		ID:             u.ID,
		Name:           u.Name,
		ProfileImageID: u.ProfileImageID,
	}

	contribution := Contribution{
		ID:         uc.ID,
		User:       user,
		Title:      uc.Title,
		Tag:        tag,
		Body:       body,
		ViewStatus: uc.ViewStatus,
		UpdatedAt:  uc.UpdatedAt,
		CreatedAt:  uc.CreatedAt,
	}

	return contribution, nil
}

// getContributionListPostList
func getContributionList(u []models.UserContribution) (contributionList []Contribution, err error) {
	if len(u) == 0 {
		return contributionList, nil
	}

	var idList []int
	var userIDList []int
	for _, val := range u {
		idList = append(idList, int(val.ID))
		userIDList = append(userIDList, int(val.UserID))
	}

	var tagMap map[int][]tags.Tag
	if tagMap, err = tags.GetMapByUserContributionIDList(idList); err != nil {
		return contributionList, err
	}

	var userMap map[int]user.User
	if userMap, err = user.GetMaptByUserIDList(userIDList); err != nil {
		return contributionList, err
	}

	var followCountMap map[int]int
	if followCountMap, err = follows.GetTotalMapByUserContributionIDList(idList); err != nil {
		return contributionList, err
	}

	soundMap, err := GetSoundMapByUserContributionIDList(idList)
	if err != nil {
		return contributionList, err
	}

	movieMap, err := GetMovieMapByUserContributionIDList(idList, models.MovieTypeYoutube)
	if err != nil {
		return contributionList, err
	}

	for _, val := range u {
		if len(tagMap[int(val.ID)]) == 0 {
			tagMap[int(val.ID)] = []tags.Tag{}
		}

		c := Contribution{
			ID:          val.ID,
			User:        userMap[val.UserID],
			Title:       val.Title,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
			ViewStatus:  val.ViewStatus,
			Tag:         tagMap[int(val.ID)],
			FollowCount: followCountMap[int(val.ID)],
			SoundStatus: soundMap[int(val.ID)].SoundStatus,
			Movie:       movieMap[int(val.ID)],
		}
		contributionList = append(contributionList, c)
	}

	return contributionList, nil
}

// GetListByTop 
func GetListByTop(offset int, size int) ([]Contribution, error) {
	u := &models.UserContribution{}
	userContribution, _, err := u.GetByTop(offset, size)
	if err != nil {
		return []Contribution{}, err
	}

	return getContributionList(userContribution)
}

// GetListBySearchValue 
func GetListBySearchValue(s []SearchValue) ([]Contribution, error) {
	idList := []int{}
	for _, v := range s {
		idList = append(idList, v.UserContributionID)
	}

	u := &models.UserContribution{}
	contributionList := []Contribution{}
	userContribution, _, err := u.GetListByIDList(idList)
	if err != nil {
		return contributionList, err
	}

	m := map[int]models.UserContribution{}
	orderMap := map[int]int{}
	for _, v := range s {
		orderMap[v.UserContributionID] = v.Order
	}

	keyList := []int{}
	for _, v := range userContribution {
		m[orderMap[int(v.ID)]] = v
		keyList = append(keyList, int(v.ID))

	}

	userContributionList := []models.UserContribution{}
	for v := range keyList {
		userContributionList = append(userContributionList, m[v])
	}

	r, err := getContributionList(userContributionList)
	if err != nil {
		return contributionList, err
	}

	for k := range r {
		for _, v := range s {
			if r[k].ID == uint(v.UserContributionID) {
				r[k].Search = v.Search
			}
		}
	}

	return r, nil
}

// GetListByFollowOrderValue
func GetListByFollowOrderValue(f []follows.OrderValue) ([]Contribution, error) {
	idList := []int{}
	for _, v := range f {
		idList = append(idList, v.UserContributionID)
	}

	u := &models.UserContribution{}
	contributionList := []Contribution{}
	userContribution, _, err := u.GetListByIDList(idList)
	if err != nil {
		return contributionList, err
	}

	m := map[int]models.UserContribution{}
	orderMap := map[int]int{}
	for _, v := range f {
		orderMap[v.UserContributionID] = v.Order
	}

	keyList := []int{}
	for _, v := range userContribution {
		m[orderMap[int(v.ID)]] = v
		keyList = append(keyList, int(v.ID))
	}

	userContributionList := []models.UserContribution{}
	for v := range keyList {
		userContributionList = append(userContributionList, m[v])
	}

	return getContributionList(userContributionList)
}

// GetViewStatusPublicIDList 
func GetViewStatusPublicIDList() ([]int, error) {
	r := []int{}

	u := models.UserContribution{}
	user, _, err := u.GetListByViewStatusPublic()
	if err != nil {
		return r, err
	}

	for _, v := range user {
		r = append(r, int(v.ID))
	}

	return r, nil
}

// ContributionListToPublic
func ContributionListToPublic(list []Contribution) []Contribution {
	r := []Contribution{}

	for _, v := range list {
		r = append(r, ContributionToPublic(v))
	}

	return r
}

// ContributionToPublic
func ContributionToPublic(c Contribution) Contribution {
	if c.ViewStatus != models.ViewStatusPublic {
		c.Body = []GetBody{}
		c.Tag = []tags.Tag{}
	}

	return c
}
