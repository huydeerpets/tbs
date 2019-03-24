package controllersProblem

import (
	"encoding/json"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/problem"
)

// AddController Add Controller
type AddController struct {
	controllers.BaseController
}

// AddRequest Additional request
type AddRequest struct {
	ID   int `form:"id"`
	Type int `form:"type"`
}

// AddResponse Additional response
type AddResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

// Post to add
func (c *AddController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		userID = 0
	}

	request := AddRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if err := problem.Add(userID, request.ID, request.Type); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = AddResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
