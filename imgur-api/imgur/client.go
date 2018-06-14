// imgur api client
//
// api doc: https://apidocs.imgur.com

package imgur

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

// XMashpeKeyHeader for imgur commercial uaage
const XMashpeKeyHeader = "X-Mashape-Key"

// APIBaseV3 imgur API version 3 endpoint
const APIBaseV3 = "https://api.imgur.com/3/"

// Endpoint is imgur's OAuth 2.0 endpoint
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://api.imgur.com/oauth2/authorize",
	TokenURL: "https://api.imgur.com/oauth2/token",
}

// Client for Imgur API v3
type Client struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	MashpeKey    string `json:"mashpe_key"`
}

// New create ImgurClinet instance
func New(id, secret, accessToken, refreshToken string) *Client {
	return &Client{
		ClientID:     id,
		ClientSecret: secret,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

// NewFromFile create ImgurClinet instance from json file
func NewFromFile(filename string) (*Client, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := new(Client)
	err = json.Unmarshal(bytes, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// OAuth2Config create oauth2.Config
func (ic *Client) OAuth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     ic.ClientID,
		ClientSecret: ic.ClientSecret,
		Endpoint:     Endpoint,
	}
}

// HTTPClient create http client with oauth2 token
func (ic *Client) HTTPClient() (*http.Client, error) {
	ctx := context.Background()
	conf := ic.OAuth2Config()

	tok := conf.TokenSource(ctx, &oauth2.Token{
		AccessToken:  ic.AccessToken,
		RefreshToken: ic.RefreshToken,
	})

	tt, err := tok.Token()
	if err != nil {
		return nil, err
	}

	return conf.Client(ctx, tt), nil
}

func parseResponse(resp *http.Response, r interface{}) (err error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println("resp body:", string(body))

	fmt.Println(resp.Status)
	if resp.StatusCode != 200 {
		return parseError(body)
	}

	err = json.Unmarshal(body, r)
	return
}
