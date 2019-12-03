package new_relic_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	BaseUrl string
	ApiKey  string

	httpClient *http.Client
}

func New(apiKey string) *Client {
	h := &http.Client{
		Timeout: 20 * time.Second,
	}

	return &Client{
		BaseUrl:    "https://api.newrelic.com",
		ApiKey:     apiKey,
		httpClient: h,
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("X-Api-Key", s.ApiKey)
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

	if !(resp.StatusCode >= 200) && !(resp.StatusCode <= 299) {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}
