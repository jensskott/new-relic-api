package new_relic_api

type AlertsChannels struct {
	Channels []struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Type          string `json:"type"`
		Configuration struct {
			IncludeJSONAttachment string `json:"include_json_attachment"`
			Recipients            string `json:"recipients"`
		} `json:"configuration"`
		Links struct {
			PolicyIds []interface{} `json:"policy_ids"`
		} `json:"links"`
	} `json:"channels"`
	Links struct {
		ChannelPolicyIds string `json:"channel.policy_ids"`
	} `json:"links"`
}

type CreateChannelsPayload struct {
	Channel Channel `json:"channel"`
}

// Channel config struct for the request
type Channel struct {
	Name          string            `json:"name"`
	Type          string            `json:"type"`
	Configuration map[string]string `json:"configuration"`
}
