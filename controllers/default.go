package controllers

import (
	"github.com/astaxie/beego"
)

// MainController MainControler
type MainController struct {
	beego.Controller
}

// Get Get
func (c *MainController) Get() {
	c.Data["StaticUrl"] = beego.AppConfig.String("staticUrl")
	c.Data["Version"] = beego.AppConfig.String("staticVersion")

	c.TplName = "index.tpl"
}
