package controllersFacebook

import (
	"github.com/huydeerpets/tbs/utils"
	"github.com/huydeerpets/tbs/utils/oauth/facebook"

	"github.com/astaxie/beego"
)

// OauthController Oauth2Controlerー
type OauthController struct {
	beego.Controller
}

// Get Certify
func (c *OauthController) Get() {
	config := facebook.GetConnect()

	state := utils.GetRandString(10)
	c.SetSession("facebookOauthState", state)

	url := config.AuthCodeURL(state)

	c.Redirect(url, 302)
}
