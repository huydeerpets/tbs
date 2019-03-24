package controllersForgetPassword

import (
	"encoding/json"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils"
	"github.com/huydeerpets/tbs/utils/user"

	validator "gopkg.in/go-playground/validator.v9"
)

// PutRequest Password save request
type PutRequest struct {
	Email    string `form:"email"`
	Keyword  string `form:"keyword"`
	Password string `form:"password" validate:"min=8,max=100"`
}

// PutResponse Password storage response
type PutResponse struct {
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

// Put Password storage
func (c *MainController) Put() {
	request := PutRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	email, err := utils.Urldecode(request.Email)
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}
	e, err := utils.Decrypter([]byte(email))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	keyword, err := utils.Urldecode(request.Keyword)
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
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	tx := models.Begin()

	if err := user.UpadateToPassword(e, request.Password); err != nil {
		models.Rollback(tx)
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	models.Commit(tx)

	c.Data["json"] = PutResponse{
		Warning: false,
		Message: "",
	}

	c.ServeJSON()
}
