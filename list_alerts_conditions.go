package new_relic_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Client) ListAlertsConditions(id string) (AlertsConditions, error) {
	var data AlertsConditions

	fullUrl := fmt.Sprintf("%s/v2/alerts_conditions.json?policy_id=%s", s.BaseUrl, id)

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
