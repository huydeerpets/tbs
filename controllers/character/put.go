package controllersCharacter

import (
	"encoding/json"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils/character"

	validator "gopkg.in/go-playground/validator.v9"
)

// PutRequest Request
type PutRequest struct {
	ID        int `form:"id" validate:"min=1"`
	VoiceType int `form:"voiceType" validate:"min=1"`
}

// PutResponse Save response
type PutResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
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

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	tx := models.Begin()

	if err := characters.SaveToVoiceType(request.ID, request.VoiceType, int(userID)); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeUserNotFound, userID)
		return
	}

	models.Commit(tx)

	c.Data["json"] = PutResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
