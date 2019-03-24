package controllersNative

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/user"
)

// CallbackController CallbackControler
type CallbackController struct {
	controllers.BaseController
}

// CallbackRequest Callbackrequest
type CallbackRequest struct {
	AccessToken string `form:"access_token"`
}

// Response response
type Response struct {
	Login bool   `json:"login"`
	Email string `json:"email"`
}

// Get Callbackする
func (c *CallbackController) Get() {
	request := CallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	url := "https://graph.facebook.com/me?access_token=" + request.AccessToken + "&fields=email"
	r, _ := http.Get(url)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	res := Response{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	u, err := user.GetByEmail(res.Email)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	login := false

	if u.ID != 0 {
		c.SetSession("user_id", u.ID)
		login = true
	}

	c.Data["json"] = DevResponse{
		Email: res.Email,
		Login: login,
	}

	c.ServeJSON()
}
