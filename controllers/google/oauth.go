package controllersGoogle

import (
	"github.com/huydeerpets/tbs/utils"
	"github.com/huydeerpets/tbs/utils/oauth/google"

	"github.com/astaxie/beego"
)

// OauthController Oauth2Controlerー
type OauthController struct {
	beego.Controller
}

// Get Certify
func (c *OauthController) Get() {
	config := google.GetConnect()

	state := utils.GetRandString(10)
	c.SetSession("googleOauthState", state)

	url := config.AuthCodeURL(state)

	c.Redirect(url, 302)
}
