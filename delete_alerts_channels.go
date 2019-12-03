package new_relic_api

import (
	"fmt"
	"net/http"
)

func (s *Client) DeleteAlertsChannels(id string) error {
	fullUrl := fmt.Sprintf("%s/v2/alerts_channels/%s.json", s.BaseUrl, id)

	req, err := http.NewRequest("DELETE", fullUrl, nil)
	if err != nil {
		return err
	}

	if _, err := s.doRequest(req); err != nil {
		return err
	}

	return nil
}
