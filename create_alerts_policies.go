package new_relic_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (s *Client) CreateAlertsPolicies(payload AlertsPoliciesPayload) (AlertsPolicies, error) {
	var data AlertsPolicies

	if err := payload.validate(); err != nil {
		return AlertsPolicies{}, err
	}

	fullUrl := fmt.Sprintf("%s/v2/alerts_policies.json", s.BaseUrl)

	body, err := json.Marshal(payload)
	if err != nil {
		return AlertsPolicies{}, err
	}

	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(body))
	if err != nil {
		return data, err
	}

	resp, err := s.doRequest(req)
	if err != nil {
		return AlertsPolicies{}, err
	}

	if err := json.Unmarshal(resp, &data); err != nil {
		return AlertsPolicies{}, err
	}

	return data, nil
}

func (p *AlertsPoliciesPayload) validate() error {
	if p.Policy.Name == "" {
		return errors.New("configuration name missing")
	}

	if p.Policy.IncidentPreference == "" {
		return errors.New("configuration incident preference missing")
	}

	return nil
}
