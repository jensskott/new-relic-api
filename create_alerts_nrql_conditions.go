package new_relic_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (s *Client) CreateAlertsNrqlConditions(payload AlertsNrqlConditionsPayload, id string) (AlertsNrqlConditions, error) {
	var data AlertsNrqlConditions

	if err := payload.validate(); err != nil {
		return AlertsNrqlConditions{}, err
	}

	fullUrl := fmt.Sprintf("%s/v2/alerts_nrql_conditions/policies/%s.json", s.BaseUrl, id)

	body, err := json.Marshal(payload)
	if err != nil {
		return AlertsNrqlConditions{}, err
	}

	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(body))
	if err != nil {
		return data, err
	}

	resp, err := s.doRequest(req)
	if err != nil {
		return AlertsNrqlConditions{}, err
	}

	if err := json.Unmarshal(resp, &data); err != nil {
		return AlertsNrqlConditions{}, err
	}

	return data, nil
}

func (p *AlertsNrqlConditionsPayload) validate() error {
	if p.NrqlCondition.Name == "" {
		return errors.New("configuration name missing")
	}

	if p.NrqlCondition.Nrql.Query == "" {
		return errors.New("configuration nrql query missing")
	}

	if p.NrqlCondition.Enabled == "" {
		return errors.New("configuration enabled missing")
	}

	if (len(p.NrqlCondition.Terms) == 1) || (len(p.NrqlCondition.Terms) == 2) {
		return errors.New("configuration terms need to have 1 or 2 values configured")
	}

	return nil
}
