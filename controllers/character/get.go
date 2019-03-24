package controllersCharacter

import (
	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/character"
)

// GettResponse response
type GettResponse struct {
	Image []characters.Image `json:"images"`
}

// Get GetList
func (c *MainController) Get() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	image, err := characters.GetImageListByUserID(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = GettResponse{
		Image: image,
	}

	c.ServeJSON()
}
