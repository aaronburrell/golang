package controllers

import (
	"github.com/revel/revel"
	"gitprojects/golang/revelapp/app/models"
)
type Group struct {
	*revel.Controller
}

func (c Group) Group() revel.Result {
	u := c.connected()
	if u != nil && u.AccessToken != "" {
	return c.Render()
	}
	return c.Redirect(App.Index)
}

func init() {
	revel.InterceptFunc(setuser, revel.BEFORE, &Group{})
}

func (c Group) connected() *models.User {
	return c.RenderArgs["user"].(*models.User)
}
