package new_relic_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Client) UpdateAlertsConditions(payload AlertsConditionsPayload) (AlertsConditions, error) {
	var data AlertsConditions

	if err := payload.validate(); err != nil {
		return AlertsConditions{}, err
	}

	fullUrl := fmt.Sprintf("%s/v2/alerts_conditions.json", *s.BaseUrl)

	body, err := json.Marshal(payload)
	if err != nil {
		return AlertsConditions{}, err
	}

	req, err := http.NewRequest("PUT", fullUrl, bytes.NewBuffer(body))
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
