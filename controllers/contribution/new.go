package controllersContribution

import (
	"encoding/json"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils/contribution"
	"github.com/huydeerpets/tbs/utils/tag"

	validator "gopkg.in/go-playground/validator.v9"
)

// NewController NewControler
type NewController struct {
	controllers.BaseController
}

// NewRequest New request
type NewRequest struct {
	Title      string `form:"title" validate:"min=1,max=100"`
	Body       string `form:"body" validate:"min=1"`
	ViewStatus int    `form:"viewStatus"`
	Tag        string `form:"tag"`
}

// Post Register new
func (c *NewController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := NewRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	tx := models.Begin()

	userContributionID, err := contributions.Add(userID, request.Title, request.Body, request.ViewStatus)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrContributionNew, userID)
		return
	}

	tag := request.Tag
	if tag != "" {
		if err := tags.AddList(int(userContributionID), tag); err != nil {
			models.Rollback(tx)
			c.ServerError(err, controllers.ErrContributionNew, userID)
			return
		}
	}

	models.Commit(tx)

	if request.ViewStatus == models.ViewStatusPublic {
		b, err := contributions.GetSearchWordBody(request.Body)
		if err != nil {
			c.ServerError(err, controllers.ErrContributionNew, userID)
			return
		}

		searchWord := contributions.SearchWord{
			Title: request.Title,
			Body:  b,
			Tag:   request.Tag,
		}

		s := contributions.JoinSearchWord(searchWord)
		if err := contributions.AddSearch(int(userContributionID), s); err != nil {
			c.ServerError(err, controllers.ErrContributionNew, userID)
			return
		}
	}

	c.Data["json"] = userContributionID
	c.ServeJSON()
}
