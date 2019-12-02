package new_relic_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"stash.massiveinteractive.com/to/new-relic-api/lib/convert"
)

type Client struct {
	BaseUrl *string
	ApiKey  *string

	httpClient *http.Client
}

func New(apiKey string) (*Client, error) {
	h := &http.Client{
		Timeout: 20 * time.Second,
	}

	return &Client{
		BaseUrl:    convert.Str("https://api.newrelic.com"),
		ApiKey:     convert.Str(apiKey),
		httpClient: h,
	}, nil
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("X-Api-Key", *s.ApiKey)
	req.Header.Set("content-type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
