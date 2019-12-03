package new_relic_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Client) UpdateAlertsPluginsConditions(payload AlertsPluginsConditionsPayload, id string) (AlertsPluginsConditions, error) {
	var data AlertsPluginsConditions

	if err := payload.validate(); err != nil {
		return AlertsPluginsConditions{}, err
	}

	fullUrl := fmt.Sprintf("%s/v2/alerts_plugins_conditions/%s.json", s.BaseUrl, id)

	body, err := json.Marshal(payload)
	if err != nil {
		return AlertsPluginsConditions{}, err
	}

	req, err := http.NewRequest("PUT", fullUrl, bytes.NewBuffer(body))
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
