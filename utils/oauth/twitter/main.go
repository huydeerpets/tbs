package twitter

import (
	"encoding/json"
	"net/url"

	"errors"

	"github.com/astaxie/beego"
	"github.com/garyburd/go-oauth/oauth"
)

// AccountAccount
type Account struct {
	ID              string `json:"id_str"`
	ScreenName      string `json:"screen_name"`
	ProfileImageURL string `json:"profile_image_url"`
	Email           string `json:"email"`
}

// GetConnect
func GetConnect() *oauth.Client {
	return &oauth.Client{
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
		Credentials: oauth.Credentials{
			Token:  beego.AppConfig.String("twitterConsumerKey"),
			Secret: beego.AppConfig.String("twitterConsumerSecret"),
		},
	}
}

// GetAccessToken 
func GetAccessToken(rt *oauth.Credentials, oauthVerifier string) (*oauth.Credentials, error) {
	oc := GetConnect()
	at, _, err := oc.RequestToken(nil, rt, oauthVerifier)

	return at, err
}

// GetMe 
func GetMe(at *oauth.Credentials, user *Account) error {
	oc := GetConnect()

	v := url.Values{}
	v.Set("include_email", "true")

	resp, err := oc.Get(nil, at, "https://api.twitter.com/1.1/account/verify_credentials.json", v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return errors.New("Twitter is unavailable")
	}

	if resp.StatusCode >= 400 {
		return errors.New("Twitter request is invalid")
	}

	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return err
	}

	return nil

}

// PostTweet 
func PostTweet(at *oauth.Credentials) error {
	oc := GetConnect()

	v := url.Values{}
	v.Set("status", "Test post. Posted from API. \n Post source：https://github.com/wheatandcat/example-golang-oauth2")

	resp, err := oc.Post(nil, at, "https://api.twitter.com/1.1/statuses/update.json", v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}
