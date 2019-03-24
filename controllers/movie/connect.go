package controllersMovie

import (
	"errors"
	"strconv"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/contribution"
	"github.com/huydeerpets/tbs/utils/movie"
)

// ConnectController Connection Controller
type ConnectController struct {
	controllers.BaseController
}

// Get Connecting
func (c *ConnectController) Get() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.RedirectError(errors.New("login not found"), 0)
		return
	}

	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.RedirectError(err, userID)
		return
	}

	u, err := contributions.GetByUserContributionID(id)
	if err != nil {
		c.RedirectError(err, userID)
		return
	}

	if userID != u.UserID {
		c.RedirectError(errors.New("diff user_id"), userID)
		return
	}

	if !contributions.ExistsMovie(id) {
		c.RedirectError(errors.New("not found movie"), userID)
		return
	}

	config := movie.GetConnect()

	url := config.AuthCodeURL(strconv.Itoa(id))

	c.Redirect(url, 302)
}
