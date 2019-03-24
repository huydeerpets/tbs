package controllersNative

import (
	"github.com/huydeerpets/tbs/controllers"
)

// DevRedirectController CallbackControler
type DevRedirectController struct {
	controllers.BaseController
}

// DevRedirectRequest Callbackrequest
type DevRedirectRequest struct {
	AccessToken string `form:"access_token"`
}

// Get Callback
func (c *DevRedirectController) Get() {
	request := DevRedirectRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	url := "exp://ts-xwe.wheatandcat.dotstamp-native.exp.direct:80/?access_token=" + request.AccessToken
	c.Redirect(url, 302)
}
