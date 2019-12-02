package new_relic_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ListAlertChannels struct {
	Channels []struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Type          string `json:"type"`
		Configuration struct {
			UserID string `json:"user_id"`
		} `json:"configuration"`
		Links struct {
			PolicyIds []interface{} `json:"policy_ids"`
		} `json:"links"`
	} `json:"channels"`
}

func (s *Client) ListAlertChannels() (ListAlertChannels, error) {
	var data ListAlertChannels

	fullUrl := fmt.Sprintf("%s/%s", *s.BaseUrl, "v2/alerts_channels.json")

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
