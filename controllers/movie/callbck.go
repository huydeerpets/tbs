package controllersMovie

import (
	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/contribution"

	"github.com/astaxie/beego"
)

// CallbackController CallbackControler
type CallbackController struct {
	controllers.BaseController
}

// CallbackRequest Callbackrequest
type CallbackRequest struct {
	Code               string `form:"code"`
	UserContributionID int    `form:"state"`
}

// Get Callbackする
func (c *CallbackController) Get() {
	request := CallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	u, err := contributions.GetUploadByUserContributionID(request.UserContributionID)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	if u.ID == uint(0) {
		err = contributions.AddUpload(request.UserContributionID, request.Code)
	} else {
		u.Token = request.Code
		err = u.Save()
	}

	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	c.Redirect(beego.AppConfig.String("topurl")+"static/html/success.html", 302)
}
