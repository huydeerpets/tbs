package controllersSound

import (
	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/models"
)

// LengthController Controler
type LengthController struct {
	controllers.BaseController
}

// LengthResponse response
type LengthResponse struct {
	Character float32 `json:"character"`
}

// Get Get the length per character
func (c *LengthController) Get() {
	u := models.UserContributionSoundLength{}
	list, _, err := u.GetByTop(0, 1000)
	if err != nil {
		c.ServerError(err, controllers.ErrCodeCommon, 0)
		return
	}

	s := 0
	l := 0
	for _, v := range list {
		if v.Second == 0 {
			continue
		}
		s += v.Second
		l += v.Length
	}

	c.Data["json"] = LengthResponse{
		Character: float32(s) / float32(l),
	}
	c.ServeJSON()
}
