package controllers

import (
	"github.com/revel/revel"
	"gitprojects/golang/revelapp/app/models"
	"log"
)
type Group struct {
	*revel.Controller
}

func (c Group) Group() revel.Result {
	u := c.connected()
	log.Println("HERE")
	if u != nil && u.AccessToken != "" {
	return c.Render()
	}
	c.Session["redirect"] = "/groups"
	return c.Redirect(App.Index)
}

func init() {
	revel.InterceptFunc(setuser, revel.BEFORE, &Group{})
}

func (c Group) connected() *models.User {
	return c.RenderArgs["user"].(*models.User)
}
