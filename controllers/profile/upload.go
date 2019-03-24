package controllersUserProfile

import (
	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/user"
	"strconv"
)

// UploadController upload
type UploadController struct {
	controllers.BaseController
}

// Post upload
func (c *UploadController) Post() {
	userID := c.GetUserID()
	if !c.IsNoLogin(userID) {
		c.ServerLoginNotFound()
		return
	}

	id, err := user.GetIDAndAddProfileImage(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
	}

	fileName := strconv.Itoa(int(id)) + ".jpg"

	var code int
	if code, err = c.SetImageFileResize(fileName, "icon", 60, 60); err != nil {
		c.ServerError(err, code, userID)
		return
	}

	if err = user.UpadateToProfileImageID(userID, int(id)); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
	}

	c.Data["json"] = fileName
	c.ServeJSON()
}
