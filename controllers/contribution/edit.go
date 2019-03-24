package controllersContribution

import (
	"strconv"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/models"
	"github.com/huydeerpets/tbs/utils/contribution"
)

// EditController Edit Controller
type EditController struct {
	controllers.BaseController
}

// EditResponse Edit response
type EditResponse struct {
	contributions.Contribution
	Sound     bool `json:"sound"`
	SoundFile bool `json:"soundFile"`
}

// Get edit
func (t *EditController) Get() {
	userID := t.GetUserID()
	if !t.IsNoLogin(userID) {
		t.ServerLoginNotFound()
		return
	}

	id, err := strconv.Atoi(t.Ctx.Input.Param(":id"))
	if err != nil {
		t.ServerError(err, controllers.ErrParameter, userID)
		return
	}

	c, err := contributions.GetContributionByUserContributionID(id)
	if err != nil {
		t.ServerError(err, controllers.ErrCodeCommon, userID)
		return
	}

	if int(c.User.ID) != userID {
		t.ServerError(err, controllers.ErrCodeUserNotFound, userID)
		return
	}

	s, err := contributions.GetSoundByUserContributionID(id)
	if err != nil {
		t.ServerError(err, controllers.ErrParameter, userID)
		return
	}

	soundFile := false
	if s.SoundStatus == models.SoundStatusPublic {
		soundFile = contributions.ExistsSound(id)
	}

	t.Data["json"] = EditResponse{
		Contribution: c,
		Sound:        (s.ID != uint(0)),
		SoundFile:    soundFile,
	}

	t.ServeJSON()
}
