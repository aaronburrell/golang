package controllers

import (
	"github.com/revel/revel"
)
type Group struct {
	App
}

func (c Group) Group() revel.Result {

	if c.validateUser("/groups") {
	return c.Render()
	}
	return c.Redirect(App.Index)

}
