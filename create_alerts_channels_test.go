package new_relic_api_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	new_relic_api "stash.massiveinteractive.com/to/new-relic-api"
)

func TestClient_CreateAlertsChannels(t *testing.T) {

	apiResponse := `
{
    "channels": [{
        "id": 1,
        "name": "Test",
        "type": "email",
        "configuration": {
            "include_json_attachment": "true",
            "recipients": "test@google.com"
        },
        "links": {
            "policy_ids": []
        }
    }],
    "links": {
        "channel.policy_ids": "/v2/policies/{policy_id}"
    }
}`

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "1234", r.Header.Get("X-Api-Key"))
		w.Write([]byte(apiResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c := new_relic_api.New("1234")

	c.HttpClient = httpClient

	resp, err := c.CreateAlertsChannels(new_relic_api.ChannelsPayload{Channel: new_relic_api.Channel{
		Name:          "Test",
		Type:          "email",
		Configuration: map[string]string{"include_json_attachment": "true", "recipients": "test@google.com"},
	}})

	assert.NoError(t, err)

	assert.Equal(t, "Test", resp.Channels[0].Name)
	assert.Equal(t, 1, resp.Channels[0].ID)
	assert.Equal(t, "test@google.com", resp.Channels[0].Configuration.Recipients)
	assert.Equal(t, "true", resp.Channels[0].Configuration.IncludeJSONAttachment)

}

func TestClient_CreateAlertsChannelsValidateName(t *testing.T) {
	c := new_relic_api.New("1234")

	_, err := c.CreateAlertsChannels(new_relic_api.ChannelsPayload{Channel: new_relic_api.Channel{}})

	assert.Error(t, err)
}

func TestClient_CreateAlertsChannelsValidateType(t *testing.T) {
	c := new_relic_api.New("1234")

	_, err := c.CreateAlertsChannels(new_relic_api.ChannelsPayload{Channel: new_relic_api.Channel{
		Name: "Test",
	}})

	assert.Error(t, err)
}

func TestClient_CreateAlertsChannelsValidateConfiguration(t *testing.T) {
	c := new_relic_api.New("1234")

	_, err := c.CreateAlertsChannels(new_relic_api.ChannelsPayload{Channel: new_relic_api.Channel{
		Name: "Test",
		Type: "email",
	}})

	assert.Error(t, err)
}
