package new_relic_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (s *Client) CreateAlertsChannels(payload CreateChannelsPayload) (AlertsChannels, error) {
	var data AlertsChannels

	if err := payload.validate(); err != nil {
		return AlertsChannels{}, err
	}

	fullUrl := fmt.Sprintf("%s/v2/alerts_channels.json", *s.BaseUrl)

	body, err := json.Marshal(payload)
	if err != nil {
		return AlertsChannels{}, err
	}

	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(body))
	if err != nil {
		return data, err
	}

	resp, err := s.doRequest(req)
	if err != nil {
		return AlertsChannels{}, err
	}

	if err := json.Unmarshal(resp, &data); err != nil {
		return AlertsChannels{}, err
	}

	return data, nil
}

func (p *CreateChannelsPayload) validate() error {
	if p.Channel.Name == "" {
		return errors.New("configuration name missing")
	}

	if p.Channel.Type == "" {
		return errors.New("configuration type missing")
	}

	if p.Channel.Configuration == nil {
		return errors.New("configuration for type missing")
	}

	return nil
}
