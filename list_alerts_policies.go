package new_relic_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Client) ListAlertsPolicies(name, exactMatch string) (AlertsPolicies, error) {
	var data AlertsPolicies
	var fullUrl string

	if name != "" && exactMatch != "" {
		fullUrl = fmt.Sprintf("%s/v2/alerts_policies.json?filter[name]=%s&filter[exact_match]=%s", s.BaseUrl, name, exactMatch)
	} else {
		fullUrl = fmt.Sprintf("%s/v2/alerts_policies.json", s.BaseUrl)
	}

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
