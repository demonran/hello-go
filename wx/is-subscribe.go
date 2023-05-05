package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
	"github.com/silenceper/wechat/v2/officialaccount/user"
)

const (
	appid     = "wx65826fad306848bc"
	appsecret = "b103f1c39ac98524b86555d10ddd418c"
	token     = "VkZkd2JrNUZPVlZhTUhOTENnPT0K"
)

var (
	offAccount *officialaccount.OfficialAccount
	auth       *oauth.Oauth
)

func init() {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     appid,
		AppSecret: appsecret,
		Token:     token,
		Cache:     memory,
	}
	offAccount = wc.GetOfficialAccount(cfg)

	auth = oauth.NewOauth(offAccount.GetContext())
}

func getUserInfo(code string) (*user.Info, error) {
	if len(code) == 0 {
		return nil, errors.New("empty code")
	}

	token, err := auth.GetUserAccessToken(code)
	if err != nil {
		return nil, err
	}
	fmt.Printf("access token: %+v\n", token)

	_user := user.NewUser(offAccount.GetContext())
	userInfo, err := _user.GetUserInfo(token.OpenID)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user info: %+v\n", userInfo)

	return userInfo, nil
}

type (
	ErrResp struct {
		Message string `json:"message"`
	}
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/login", func(c echo.Context) error {
		redirectUrl := "https://a721-2409-8a62-e41-34f0-18cf-5b5c-12c1-9be8.ngrok-free.app/info"
		scope := "snsapi_base"
		state := string(rand.Intn(2 * 16))
		url, err := auth.GetRedirectURL(redirectUrl, scope, state)
		if err != nil {
			return err
		}
		return c.Redirect(http.StatusFound, url)
	})

	e.GET("/info", func(c echo.Context) error {
		code := c.QueryParam("code")

		info, err := getUserInfo(code)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrResp{Message: err.Error()})
		}

		return c.JSON(http.StatusOK, info)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
