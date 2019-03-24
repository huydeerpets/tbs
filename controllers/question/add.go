package controllersQuestion

import (
	"encoding/json"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/question"

	validator "gopkg.in/go-playground/validator.v9"
)

// AddController Add Controller
type AddController struct {
	controllers.BaseController
}

// AddRequest Additional request
type AddRequest struct {
	Body  string `form:"body" validate:"min=1"`
	Email string `form:"email" validate:"required,email"`
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

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if err := question.Add(userID, request.Body, request.Email); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = AddResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
