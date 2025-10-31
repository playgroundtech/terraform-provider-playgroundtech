package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// HostURL - Default URL
const HostURL string = "https://jobapi.aws.playgroundtech.cloud"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// AuthStruct -
type AuthStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	UserID int    `json:"id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

// NewClient -
func NewClient(host, password, email *string) (*Client, error) {

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if (password != nil) && (email != nil) {
		// form request body
		rb, err := json.Marshal(AuthStruct{
			Password: *password,
			Email:    *email,
		})
		if err != nil {
			return nil, err
		}

		// authenticate
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/login", c.HostURL), strings.NewReader(string(rb)))
		if err != nil {
			return nil, err
		}

		body, err := c.doRequest(req)
		if err != nil {
			return nil, err
		}

		// parse response body
		ar := AuthResponse{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return nil, err
		}
		if ar.Token == "" {
			return nil, fmt.Errorf("body is: %s", body)
		}
		c.Token = ar.Token
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "Bearer "+c.Token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
