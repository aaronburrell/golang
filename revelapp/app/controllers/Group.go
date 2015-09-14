package controllers

import (
	"github.com/revel/revel"
)
type Group struct {
	App
}

func (c Group) Group() revel.Result {
	u := c.connected()
	if u != nil && u.AccessToken != "" {
	return c.Render()
	}
	c.Session["redirect"] = "/groups"
	return c.Redirect(App.Index)
}
