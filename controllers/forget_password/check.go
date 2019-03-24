package controllersForgetPassword

import (
	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils"
	"github.com/huydeerpets/tbs/utils/user"
)

// CheckController Forgotten password confirmation Controller
type CheckController struct {
	controllers.BaseController
}

// CheckResponse Forgot password confirmation response
type CheckResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

// Get Confirm password forgotten
func (c *CheckController) Get() {
	email, err := utils.Urldecode(c.Ctx.Input.Param(":email"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}
	e, err := utils.Decrypter([]byte(email))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	keyword, err := utils.Urldecode(c.Ctx.Input.Param(":keyword"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}
	k, err := utils.Decrypter([]byte(keyword))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	r, err := user.IsUpdatePassword(e, k)
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	if r == false {
		c.Data["json"] = CheckResponse{
			Warning: true,
			Message: "Invalid URL",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = CheckResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
