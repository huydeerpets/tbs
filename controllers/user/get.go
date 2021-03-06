package controllersUser

import (
	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/user"
)

// GetResponse Detail response
type GetResponse struct {
	User    user.User      `json:"user"`
	Profile []user.Profile `json:"profiles"`
}

// Get User information
func (c *MainController) Get() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	u, err := user.GetByUserID(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeUserNotFound, userID)
	}

	p, err := user.GetProfileImageListByUserID(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeUserNotFound, userID)
	}

	c.Data["json"] = GetResponse{
		User:    u,
		Profile: p,
	}
	c.ServeJSON()
}
