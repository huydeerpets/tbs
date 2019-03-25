package controllersUser

import (
	"encoding/json"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils/contribution"
)

// ContributionListControllerPostListControlerー
type ContributionListController struct {
	controllers.BaseController
}

// ContributionListRequestPostListrequest
type ContributionListRequest struct {
	Order int `form:"order"`
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

// ContributionListResponsePostListresponse
type ContributionListResponse struct {
	PrivteList []models.UserContribution `json:"privtes"`
	List       []models.UserContribution `json:"list"`
	TitleList  []string                  `json:"titles"`
	Count      int                       `json:"count"`
}

// Post Get user post list
func (c *ContributionListController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := ContributionListRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	orderMap := map[int]string{
		1: "ID desc",
		2: "ID asc",
	}

	limit := 1000
	offset := 0

	userContributionlist, err := contributions.GetListByUserID(userID, orderMap[request.Order], limit, offset)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	title := []string{}

	for _, v := range userContributionlist {
		title = append(title, v.Title)
	}

	privte := []models.UserContribution{}

	for _, v := range userContributionlist {
		if v.ViewStatus == models.ViewStatusPrivate {
			privte = append(privte, v)
		}
	}

	count, err := contributions.GetCountByUserID(userID, orderMap[request.Order])
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = ContributionListResponse{
		PrivteList: privte,
		List:       userContributionlist,
		Count:      count,
		TitleList:  title,
	}
	c.ServeJSON()
}
