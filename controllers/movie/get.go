package controllersMovie

import (
	"strconv"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils/contribution"
)

// GetResponse Confirmation response
type GetResponse struct {
	Warning     bool   `json:"warning"`
	Message     string `json:"message"`
	MovieStatus int    `json:"movieStatus"`
}

// Get Check
func (c *MainController) Get() {
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

	userMovie, err := contributions.GetMovie(id, models.MovieTypeYoutube)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	c.Data["json"] = GetResponse{
		Warning:     false,
		Message:     "",
		MovieStatus: userMovie.MovieStatus,
	}

	c.ServeJSON()
}
