package controllersSound

import (
	"encoding/json"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils/contribution"

	validator "gopkg.in/go-playground/validator.v9"
)

// SaveBodyController Save textControler
type SaveBodyController struct {
	controllers.BaseController
}

// SaveBodydRequest Save textrequest
type SaveBodydRequest struct {
	ID   uint   `form:"id" validate:"min=1"`
	Body string `form:"body" validate:"max=256"`
}

// SaveBodyResponse Save textresponse
type SaveBodyResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
	ID      uint   `json:"id"`
}

// Put Save text
func (c *SaveBodyController) Put() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	request := SaveBodydRequest{}
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

	if err := contributions.SaveSoundDetailToBodySound(request.ID, request.Body, userID); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	models.Commit(tx)

	u := models.UserContributionSoundDetail{}
	r, _, err := u.GetByID(request.ID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if err := contributions.AddTmpSound(r); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = SaveBodyResponse{
		Warning: false,
		Message: "",
		ID:      request.ID,
	}

	c.ServeJSON()
}
