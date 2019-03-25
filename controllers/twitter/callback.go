package controllersTwitter

import (
	"net/url"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/oauth/twitter"
	"github.com/huydeerpets/tbs/utils/user"

	"github.com/astaxie/beego"
	"github.com/garyburd/go-oauth/oauth"
)

// CallbackController CallbackControler
type CallbackController struct {
	controllers.BaseController
}

// CallbackRequest Callbackrequest
type CallbackRequest struct {
	Token    string `form:"oauth_token"`
	Verifier string `form:"oauth_verifier"`
}

// Get Callback
func (c *CallbackController) Get() {
	c.StartSession()

	request := CallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		c.RedirectError(err, 0)
		return
	}

	at, err := twitter.GetAccessToken(
		&oauth.Credentials{
			Token:  c.CruSession.Get("request_token").(string),
			Secret: c.CruSession.Get("request_token_secret").(string),
		},
		request.Verifier,
	)
	if err != nil {
		c.RedirectError(err, 0)
	}

	c.CruSession.Set("oauth_secret", at.Secret)
	c.CruSession.Set("oauth_token", at.Token)

	account := twitter.Account{}
	if err = twitter.GetMe(at, &account); err != nil {
		c.RedirectError(err, 0)
	}

	u, err := user.GetByEmail(account.Email)
	if err != nil {
		c.RedirectError(err, 0)
		return
	}

	if u.ID != 0 {
		c.SetSession("user_id", u.ID)
		c.Redirect(beego.AppConfig.String("topurl"), 302)
		return
	}

	url := beego.AppConfig.String("topurl") + "oauth/?email=" + url.QueryEscape(account.Email)
	c.Redirect(url, 302)
}
