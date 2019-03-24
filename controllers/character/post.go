package controllersCharacter

import (
	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/character"
)

// PostResponse Upload response
type PostResponse struct {
	Image []characters.Image `json:"images"`
}

// PostImage upload
func (c *MainController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	id, err := characters.AddImage(userID, 0, 0)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	filePath := characters.GetImageName(id)

	var code int
	if code, err = c.SetImageFileResize(filePath, "character", 180, 180); err != nil {
		characters.DeleteByID(int(id), userID)
		c.ServerError(err, code, userID)
		return
	}

	image, err := characters.GetImageListByUserID(userID)
	if err != nil {
		characters.DeleteByID(int(id), userID)
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = PostResponse{
		Image: image,
	}

	c.ServeJSON()
}
