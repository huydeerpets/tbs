package controllers

import (
	"github.com/astaxie/beego"
)

// ErrorController Error Controller
type ErrorController struct {
	beego.Controller
}

// Error404 Error:404
func (c *ErrorController) Error404() {
	c.Redirect("http://206.189.90.165:8080/", 200)
}
