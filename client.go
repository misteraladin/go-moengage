package moengage

import (
	b64 "encoding/base64"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	Client  *http.Client
	BaseUrl string
	APIKey  string
	APPID   string
}

func NewClient(url, apiKey, appID string) Client {
	var HTTPTimeout = 60 * time.Second

	return Client{
		Client:  &http.Client{Timeout: HTTPTimeout},
		BaseUrl: url,
		APIKey:  apiKey,
		APPID:   appID,
	}
}

func (c *Client) Call(method string, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		log.Print("Cannot create http request : ", err)
		return nil, err
	}

	authStr := b64.StdEncoding.EncodeToString([]byte(c.APPID + ":" + c.APIKey))
	req.Header.Add("Authorization", "Basic "+authStr)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.Client.Do(req)
	if err != nil {
		log.Print("Cannot send request : ", err)
		return nil, err
	}

	return res, nil
}
