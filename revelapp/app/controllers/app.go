package controllers

import (
	"golang.org/x/oauth2"
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"github.com/aaronburrell/golang/revelapp/app/models"
	"net/http"
	"net/url"
	"strconv"
	"log"
)

type App struct {
	*revel.Controller
}

var GOOGLE = &oauth2.Config{
	ClientID:     "966343595637-da2m2ti69g48f0ct7jtml5vtto7ub7un.apps.googleusercontent.com",
	ClientSecret: "2JrU5VRNSrTVexzq7_6A5DSD",
	Scopes:       []string{"openid"},
	Endpoint: oauth2.Endpoint{
				AuthURL:  "https://accounts.google.com/o/oauth2/auth",
				TokenURL: "https://www.googleapis.com/oauth2/v3/token",
    },
	RedirectURL:  "http://localhost:9000/App/Auth",
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
		//resp, _ := http.Get("https://api.github.com/user?access_token=" +
		//	url.QueryEscape(u.AccessToken))
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "https://www.googleapis.com/plus/v1/people/me/openIdConnect", nil)
		req.Header.Add("Authorization", "Bearer " + url.QueryEscape(u.AccessToken))
		resp, _ := client.Do(req)

		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&me); err != nil {
			revel.ERROR.Println(err)
		}
		revel.INFO.Println(me)
	}

	authUrl := GOOGLE.AuthCodeURL("foo")
	return c.Render(me, authUrl)
}


func (c App) Auth(code string) revel.Result {
	tok, err := GOOGLE.Exchange(oauth2.NoContext, code)
	if err != nil {
		revel.ERROR.Println(err)
		return c.Redirect(App.Index)
	}
	user := c.connected()
	user.AccessToken = tok.AccessToken

	if _, ok := c.Session["redirect"]; ok {
		if c.Session["redirect"] != "" {
		redirectURL := c.Session["redirect"]
		c.Session["redirect"] = ""
		return c.Redirect(redirectURL)
	}
	}

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

func (c App) validateUser(redirectUrl string) bool {
		u := c.connected()
		if u != nil && u.AccessToken != "" {
			c.Session["redirect"] = redirectUrl
			return true
		}
		return false
}
