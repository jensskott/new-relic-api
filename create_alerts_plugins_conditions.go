package new_relic_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (s *Client) CreateAlertsPluginsConditions(payload AlertsPluginsConditionsPayload, id string) (AlertsPluginsConditions, error) {
	var data AlertsPluginsConditions

	if err := payload.validate(); err != nil {
		return AlertsPluginsConditions{}, err
	}

	fullUrl := fmt.Sprintf("%s/v2/alerts_plugins_conditions/policies/%s.json", s.BaseUrl, id)

	body, err := json.Marshal(payload)
	if err != nil {
		return AlertsPluginsConditions{}, err
	}

	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(body))
	if err != nil {
		return data, err
	}

	resp, err := s.doRequest(req)
	if err != nil {
		return AlertsPluginsConditions{}, err
	}

	if err := json.Unmarshal(resp, &data); err != nil {
		return AlertsPluginsConditions{}, err
	}

	return data, nil
}

func (p *AlertsPluginsConditionsPayload) validate() error {
	if p.PluginsCondition.Name == "" {
		return errors.New("configuration name missing")
	}

	if p.PluginsCondition.Plugin.ID == "" {
		return errors.New("configuration plugin ID missing")
	}

	if p.PluginsCondition.Enabled == "" {
		return errors.New("configuration enabled missing")
	}

	if (len(p.PluginsCondition.Terms) == 1) || (len(p.PluginsCondition.Terms) == 2) {
		return errors.New("configuration terms need to have 1 or 2 values configured")
	}

	return nil
}
