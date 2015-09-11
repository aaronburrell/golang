package controllers

import (
	"golang.org/x/oauth2"
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"gitprojects/golang/revelapp/app/models"
	"net/http"
	"net/url"
	"strconv"
	"log"
)

type App struct {
	*revel.Controller
}

var GITHUB = &oauth2.Config{
	ClientID:     "316920733717ec5b4771",
	ClientSecret: "c4c9d450d1dceba04632aea69bd7c4c839d917f1",
	Endpoint: oauth2.Endpoint{
        AuthURL:  "https://github.com/login/oauth/authorize",
        TokenURL: "https://github.com/login/oauth/access_token",
    },
	RedirectURL:  "http://localhost:9000/App/Auth",
}

func (c App) Index() revel.Result {
	u := c.connected()
	me := map[string]interface{}{}
	if u != nil && u.AccessToken != "" {
		log.Println("This is the user: " + u.AccessToken)
		resp, _ := http.Get("https://api.github.com/user?access_token=" +
			url.QueryEscape(u.AccessToken))
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&me); err != nil {
			revel.ERROR.Println(err)
		}
		revel.INFO.Println(me)
	}

	authUrl := GITHUB.AuthCodeURL("foo")
	return c.Render(me, authUrl)
}

func (c App) Auth(code string) revel.Result {
	//t := &oauth2.Transport{Config: GITHUB}
	//tok, err := t.Exchange(code)
	tok, err := GITHUB.Exchange(oauth2.NoContext, code)
	if err != nil {
		revel.ERROR.Println(err)
		return c.Redirect(App.Index)
	}

	user := c.connected()
	user.AccessToken = tok.AccessToken
	return c.Redirect(App.Index)
}

func setuser(c *revel.Controller) revel.Result {
	var user *models.User
	if _, ok := c.Session["uid"]; ok {
		uid, _ := strconv.ParseInt(c.Session["uid"], 10, 0)
		user = models.GetUser(int(uid))
	}
	if user == nil {
		user = models.NewUser()
		c.Session["uid"] = fmt.Sprintf("%d", user.Uid)
	}
	c.RenderArgs["user"] = user
	return nil
}

func init() {
	revel.InterceptFunc(setuser, revel.BEFORE, &App{})
}

func (c App) connected() *models.User {
	return c.RenderArgs["user"].(*models.User)
}
