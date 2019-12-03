package new_relic_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (s *Client) CreateAlertsConditions(payload AlertsConditionsPayload) (AlertsConditions, error) {
	var data AlertsConditions

	if err := payload.validate(); err != nil {
		return AlertsConditions{}, err
	}

	fullUrl := fmt.Sprintf("%s/v2/alerts_conditions.json", s.BaseUrl)

	body, err := json.Marshal(payload)
	if err != nil {
		return AlertsConditions{}, err
	}

	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(body))
	if err != nil {
		return data, err
	}

	resp, err := s.doRequest(req)
	if err != nil {
		return AlertsConditions{}, err
	}

	if err := json.Unmarshal(resp, &data); err != nil {
		return AlertsConditions{}, err
	}

	return data, nil
}

func (p *AlertsConditionsPayload) validate() error {
	if p.Condition.Name == "" {
		return errors.New("configuration name missing")
	}

	if p.Condition.Type == "" {
		return errors.New("configuration type missing")
	}

	if p.Condition.Enabled == "" {
		return errors.New("configuration enabled missing")
	}

	if (len(p.Condition.Terms) == 1) || (len(p.Condition.Terms) == 2) {
		return errors.New("configuration terms need to have 1 or 2 values configured")
	}

	if p.Condition.Metric == "" {
		return errors.New("configuration metric missing")
	}

	return nil
}
