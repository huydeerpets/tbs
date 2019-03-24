package controllersLogin

import (
	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/user"
)

// AuthController Authentication Controller
type AuthController struct {
	controllers.BaseController
}

// AuthResponse Authentication response
type AuthResponse struct {
	Login bool   `json:"login"`
	Name  string `json:"name"`
}

// Post Determine if you are logged in
func (c *AuthController) Post() {
	userID := c.GetUserID()

	var response AuthResponse
	if !c.IsNoLogin(userID) {
		response = AuthResponse{
			Login: false,
			Name:  "",
		}
	} else {

		u, err := user.GetByUserID(userID)
		if err != nil {
			c.ServerError(err, controllers.ErrCodeUserNotFound, userID)
		}

		response = AuthResponse{
			Login: true,
			Name:  u.Name,
		}
	}

	c.Data["json"] = response

	c.ServeJSON()
}
