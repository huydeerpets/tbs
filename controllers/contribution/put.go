package controllersContribution

import (
	"encoding/json"
	"strconv"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils/contribution"
	"github.com/huydeerpets/tbs/utils/tag"

	validator "gopkg.in/go-playground/validator.v9"
)

// PutRequest PUTrequest
type PutRequest struct {
	Title      string `form:"title" validate:"min=1,max=100"`
	Body       string `form:"body" validate:"min=1"`
	ViewStatus int    `form:"viewStatus"`
}

// Put Save
func (c *MainController) Put() {

	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := PutRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	tx := models.Begin()

	if err := contributions.Save(id, userID, request.Title, request.ViewStatus); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrContributionSave, userID)
		return
	}

	if err := contributions.SaveDetail(id, request.Body); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrContributionSave, userID)
		return
	}

	models.Commit(tx)

	if request.ViewStatus == models.ViewStatusPublic {
		t, err := tags.GetTagNameJoin(id)
		if err != nil {
			c.ServerError(err, controllers.ErrContributionSave, userID)
			return
		}

		b, err := contributions.GetSearchWordBody(request.Body)
		if err != nil {
			c.ServerError(err, controllers.ErrContributionNew, userID)
			return
		}

		searchWord := contributions.SearchWord{
			Title: request.Title,
			Body:  b,
			Tag:   t,
		}

		s := contributions.JoinSearchWord(searchWord)
		if err := contributions.AddOrSaveSearch(id, s); err != nil {
			c.ServerError(err, controllers.ErrContributionSave, userID)
			return
		}

	} else {
		if err := contributions.DeleteSearchByUserContributionID(id); err != nil {
			c.ServerError(err, controllers.ErrContributionSave, userID)
			return
		}
	}

	c.Data["json"] = id
	c.ServeJSON()
}
