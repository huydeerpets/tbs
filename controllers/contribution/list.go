package controllersContribution

import (
	"strconv"

	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/utils/contribution"
)

// ListController list controller
type ListController struct {
	controllers.BaseController
}

// Get GetList
func (c *ListController) Get() {
	order, err := strconv.Atoi(c.Ctx.Input.Param(":order"))
	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	contributionlist, err := contributions.GetListByTop(0, (order+1)*10)

	if err != nil {
		c.ServerError(err, controllers.ErrParameter, 0)
		return
	}

	c.Data["json"] = contributionlist
	c.ServeJSON()
}
