package controllers

import "github.com/revel/revel"

type Group struct {
	*revel.Controller
}

func (c Group) Group() revel.Result {
	return c.Render()
}
