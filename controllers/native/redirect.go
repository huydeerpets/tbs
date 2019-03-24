package controllersNative

import (
	"github.com/huydeerpets/tbs/controllers"
)

// RedirectController CallbackControler
type RedirectController struct {
	controllers.BaseController
}

// RedirectRequest Callbackrequest
type RedirectRequest struct {
	AccessToken string `form:"access_token"`
}

// Get Callback
func (c *RedirectController) Get() {
	request := RedirectRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	url := "exp://exp.host/@wheatandcat/dotstamp_native/?access_token=" + request.AccessToken
	c.Redirect(url, 302)
}
