package controllersSound

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils/contribution"

	validator "gopkg.in/go-playground/validator.v9"
)

// SaveVoiceListController SaveVoiceListControler
type SaveVoiceListController struct {
	controllers.BaseController
}

// SaveVoiceListRequest SaveVoiceListrequest
type SaveVoiceListRequest struct {
	VoiceType int `form:"voiceType" validate:"min=1"`
}

// SaveVoiceListResponse SaveVoiceListresponse
type SaveVoiceListResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

// Put SaveVoiceList
func (c *SaveVoiceListController) Put() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, userID)
		return
	}

	request := SaveVoiceListRequest{}
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	validate := validator.New()
	if err = validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	u, err := contributions.GetByUserContributionID(id)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if userID != u.UserID {
		c.ServerError(errors.New("diff UserID"), controllers.ErrCodeCommon, userID)
		return
	}

	if err := contributions.UpdatesSoundToMakeStatusAndVoiceTypeByUserContributionID(id, models.MakeStatusUncreated, request.VoiceType); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = SaveVoiceListResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
