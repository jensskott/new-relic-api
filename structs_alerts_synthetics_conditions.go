package new_relic_api

type AlertsSyntheticsConditions struct {
	SyntheticsConditions []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		MonitorID string `json:"monitor_id"`
		Enabled   bool   `json:"enabled"`
	} `json:"synthetics_conditions"`
}

type AlertsSyntheticsConditionsPayload struct {
	SyntheticsCondition Condition `json:"SynthecticsCondition"`
}

type SyntheticsCondition struct {
	Name       string `json:"name"`
	MonitorID  string `json:"monitor_id"`
	RunbookURL string `json:"runbook_url"`
	Enabled    string `json:"enabled"`
}
