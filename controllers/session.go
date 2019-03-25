package controllers

import (
	"github.com/astaxie/beego"
)

var session = make(map[string]interface{})

// isSession Determine if you want to use a session
func isSession() bool {
	if beego.AppConfig.String("runmode") == "test" {
		return false
	}

	return true
}

// GetUserID Get user ID
func (c *BaseController) GetUserID() int {
	if !isSession() {
		r, _ := c.GetInt("user_id", 1)
		return r
	}

	uID := c.GetSession("user_id")
	if uID, ok := uID.(int); ok {
		return uID
	}
	if uintID, ok := uID.(uint); ok {
		return int(uintID)
	}

	return noUserID
}

// GetSession Get a session
func (c *BaseController) GetSession(name string) interface{} {
	if !isSession() {
		return session[name]
	}

	if c.CruSession == nil {
		c.StartSession()
	}
	return c.CruSession.Get(name)
}

// SetSession Set the session
func (c *BaseController) SetSession(name string, value interface{}) {
	if !isSession() {
		session[name] = value
		return
	}

	if c.CruSession == nil {
		c.StartSession()
	}
	c.CruSession.Set(name, value)
}

// DelSession Delete session
func (c *BaseController) DelSession(name string) {
	if !isSession() {
		delete(session, name)
		return
	}

	if c.CruSession == nil {
		c.StartSession()
	}
	c.CruSession.Delete(name)
}
