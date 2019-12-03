package new_relic_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Client) UpdateAlertsNrqlConditions(payload AlertsNrqlConditionsPayload, id string) (AlertsNrqlConditions, error) {
	var data AlertsNrqlConditions

	if err := payload.validate(); err != nil {
		return AlertsNrqlConditions{}, err
	}

	fullUrl := fmt.Sprintf("%s/v2/alerts_nrql_conditions/%s.json", s.BaseUrl, id)

	body, err := json.Marshal(payload)
	if err != nil {
		return AlertsNrqlConditions{}, err
	}

	req, err := http.NewRequest("PUT", fullUrl, bytes.NewBuffer(body))
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
