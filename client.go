package new_relic_api

import (
	"net/http"
	"net/url"
	"stash.massiveinteractive.com/to/new-relic-api/lib/convert"
	"time"
)

type client struct {
	baseUrl *url.URL
	apiKey *string

	httpClient *http.Client
}

func New (baseUrl, apiKey string) (*client, error) {
	h := &http.Client{
		Timeout:       20 * time.Second,
	}

	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	return &client{
		baseUrl:    u,
		apiKey:     convert.Str(apiKey),
		httpClient: h,
	}, nil
}
