package controllers

import (
	"github.com/aaronburrell/golang/revelapp/app/models"
	"github.com/revel/revel"
)

type Group struct {
	App
}

func (c Group) Group() revel.Result {

	if c.validateUser("/groups") {
		results, err := c.Txn.Select(models.User{}, `select * from User`)
		if err != nil {
			panic(err)
		}

		var users []*models.User

		for _, r := range results {
			u := r.(*models.User)
			users = append(users, u)
		}
		return c.Render(users)
	}
	return c.Redirect(App.Index)

}
