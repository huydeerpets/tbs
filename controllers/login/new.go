package controllersLogin

import (
	"encoding/json"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/user"

	"gopkg.in/go-playground/validator.v9"
)

// NewController New registration Controller
type NewController struct {
	controllers.BaseController
}

// NewRequest New request
type NewRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"min=8,max=100"`
}

// NewResponse New response
type NewResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
	UserID  uint   `json:"userID"`
}

// Post New login
func (c *NewController) Post() {
	request := NewRequest{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	u, err := user.GetByEmail(request.Email)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	if u.ID != uint(0) {
		c.Data["json"] = NewResponse{
			Warning: true,
			Message: "The entered email address has already been registered.",
			UserID:  0,
		}
		c.ServeJSON()
		return
	}

	userID, err := user.Add(request.Email, request.Email, request.Password)
	if err != nil {
		c.ServerError(err, controllers.ErrCreateUser, 0)
		return
	}

	c.SetSession("user_id", userID)

	c.Data["json"] = NewResponse{
		Warning: false,
		Message: "",
		UserID:  userID,
	}

	c.ServeJSON()
}
