package controllersUser

import (
	"encoding/json"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils/user"

	validator "gopkg.in/go-playground/validator.v9"
)

// PutRequest Putrequest
type PutRequest struct {
	Name string `form:"name" validate:"min=1,max=100"`
}

// PutResponse Putresponse
type PutResponse struct {
	Success bool `json:"success"`
}

// Put User information
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

	if err := user.Upadate(userID, request.Name); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrUserSave, userID)
		return
	}

	models.Commit(tx)

	c.Data["json"] = PutResponse{
		Success: true,
	}

	c.ServeJSON()
}
