package new_relic_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (s *Client) CreateAlertsSyntheticsConditions(payload AlertsSyntheticsConditionsPayload, id string) (SyntheticsCondition, error) {
	var data SyntheticsCondition

	if err := payload.validate(); err != nil {
		return SyntheticsCondition{}, err
	}

	fullUrl := fmt.Sprintf("%s/v2/alerts_conditions/policies/%s.json", s.BaseUrl, id)

	body, err := json.Marshal(payload)
	if err != nil {
		return SyntheticsCondition{}, err
	}

	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(body))
	if err != nil {
		return data, err
	}

	resp, err := s.doRequest(req)
	if err != nil {
		return SyntheticsCondition{}, err
	}

	if err := json.Unmarshal(resp, &data); err != nil {
		return SyntheticsCondition{}, err
	}

	return data, nil
}

func (p *AlertsSyntheticsConditionsPayload) validate() error {
	if p.SyntheticsCondition.Name == "" {
		return errors.New("configuration name missing")
	}

	if p.SyntheticsCondition.Type == "" {
		return errors.New("configuration type missing")
	}

	if p.SyntheticsCondition.Enabled == "" {
		return errors.New("configuration enabled missing")
	}

	return nil
}
