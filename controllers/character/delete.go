package controllersCharacter

import (
	"strconv"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils/character"
)

// DeleteController Delete - controller
type DeleteController struct {
	controllers.BaseController
}

// DeleteResponse - Delete response
type DeleteResponse struct {
	Image []characters.Image `json:"images"`
}

// Delete DeleteImage
func (c *DeleteController) Delete() {
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

	tx := models.Begin()

	if err = characters.DeleteByID(id, userID); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeUserNotFound, userID)
		return
	}

	models.Commit(tx)

	image, err := characters.GetImageListByUserID(userID)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	c.Data["json"] = DeleteResponse{
		Image: image,
	}

	c.ServeJSON()
}
