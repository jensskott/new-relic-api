package new_relic_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Client) ListAlertsChannels() (AlertsChannels, error) {
	var data AlertsChannels

	fullUrl := fmt.Sprintf("%s/v2/alerts_channels.json", *s.BaseUrl)

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return data, err
	}

	resp, err := s.doRequest(req)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(resp, &data); err != nil {
		return data, err
	}

	return data, nil
}
