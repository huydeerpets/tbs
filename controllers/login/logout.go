package controllersLogin

import "github.com/huydeerpets/tbs/controllers"

// LogoutControllerLogoutControler
type LogoutController struct {
	controllers.BaseController
}

// PostLogout
func (c *LogoutController) Post() {

	c.DelSession("user_id")

	c.Data["json"] = true

	c.ServeJSON()
}
