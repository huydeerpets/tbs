package controllersContribution

import (
	"encoding/json"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/contribution"
)

// SearchController SearchControler
type SearchController struct {
	controllers.BaseController
}

// SearchRequest Searchrequest
type SearchRequest struct {
	Search string `form:"search"`
	Order  int    `form:"order" validate:"min=1,max=2"`
	Page   int    `form:"page"`
	Limit  int    `form:"limit" validate:"min=1,max=50"`
}

// SearchResponse Searchresponse
type SearchResponse struct {
	List  []contributions.Contribution `json:"list"`
	Count int                          `json:"count"`
}

// Post Get search get
func (c *SearchController) Post() {
	request := SearchRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	orderMap := map[int]string{
		1: "ID desc",
		2: "follow_count desc",
	}

	offset := (request.Page - 1) * request.Limit

	searchValue, err := contributions.GetSearchValueListBySearch(request.Search, orderMap[request.Order], request.Limit, offset)
	if err != nil {
		c.ServerError(err, controllers.ErrContributionSearch, 0)
		return
	}

	if len(searchValue) == 0 {
		c.Data["json"] = SearchResponse{}
		c.ServeJSON()
		return
	}

	contributionlist, err := contributions.GetListBySearchValue(searchValue)
	if err != nil {
		c.ServerError(err, controllers.ErrContributionSearch, 0)
		return
	}

	count, err := contributions.GetCountBySearch(request.Search, orderMap[request.Order])
	if err != nil {
		c.ServerError(err, controllers.ErrContributionSearch, 0)
		return
	}

	c.Data["json"] = SearchResponse{
		List:  contributionlist,
		Count: count,
	}

	c.ServeJSON()
}
