package new_relic_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type CreateAlertChannels struct {
	Channels []struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Type          string `json:"type"`
		Configuration struct {
			IncludeJSONAttachment string `json:"include_json_attachment"`
			Recipients            string `json:"recipients"`
		} `json:"configuration"`
		Links struct {
			PolicyIds []interface{} `json:"policy_ids"`
		} `json:"links"`
	} `json:"channels"`
	Links struct {
		ChannelPolicyIds string `json:"channel.policy_ids"`
	} `json:"links"`
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
