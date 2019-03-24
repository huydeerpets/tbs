package controllersForgetPassword

import (
	"encoding/json"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils"
	"github.com/huydeerpets/tbs/utils/mail"
	"github.com/huydeerpets/tbs/utils/user"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/astaxie/beego"
)

// PostRequest Forgot password add request
type PostRequest struct {
	Email string `form:"email" validate:"required,email"`
}

// PostResponse Forgot Password Add Response
type PostResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

// Post Get user post list
func (c *MainController) Post() {
	request := PostRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	u, err := user.GetByEmail(request.Email)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	if u.ID == uint(0) {
		c.Data["json"] = PostResponse{
			Warning: true,
			Message: "Email address not found",
		}
		c.ServeJSON()
		return
	}

	if err = user.DeleteByEmail(request.Email); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	keyword := utils.GetRandString(50)

	tx := models.Begin()

	if err = user.AddForgetPassword(request.Email, keyword); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	var url string
	url, err = mail.GetForgetpasswordURL(request.Email, keyword)
	if err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	models.Commit(tx)

	top := beego.AppConfig.String("topurl")

	f := mail.ForgetpasswordTemplate{
		URL:   top + "password/reset/" + url,
		Host:  top,
		Email: beego.AppConfig.String("email"),
	}
	m := mail.GetForgetpasswordBody(f)
	b := mail.Body{
		From:    beego.AppConfig.String("email"),
		To:      request.Email,
		Subject: "password re-setting",
		Message: string(m),
	}

	err = mail.Send(request.Email, mail.GetBody(b))
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	c.Data["json"] = PostResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
