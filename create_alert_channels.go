package new_relic_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type CreateAlertChannels struct {
	Channel struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Type          string `json:"type"`
		Configuration string `json:"configuration"`
		Links         struct {
			PolicyIds []string `json:"policy_ids"`
		} `json:"links"`
	} `json:"channel"`
}

type Payload struct {
	Channel Channel `json:"channel"`
}

type Channel struct {
	Name          string            `json:"name"`
	Type          string            `json:"type"`
	Configuration map[string]string `json:"configuration"`
}

func (s *Client) CreateAlertChannels(payload Payload) (CreateAlertChannels, error) {
	var data CreateAlertChannels

	if payload.Channel.Name == "" {
		return CreateAlertChannels{}, errors.New("channel configuration missing")
	}

	fullUrl := fmt.Sprintf("%s/v2/alerts_channels.json", *s.BaseUrl)

	body, err := json.Marshal(payload)
	if err != nil {
		return CreateAlertChannels{}, err
	}

	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(body))
	if err != nil {
		return data, err
	}

	resp, err := s.doRequest(req)
	if err != nil {
		return CreateAlertChannels{}, err
	}

	if err := json.Unmarshal(resp, &data); err != nil {
		return CreateAlertChannels{}, err
	}

	return data, nil
}
