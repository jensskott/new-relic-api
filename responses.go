package new_relic_api

type NewRelicResponse struct {
	Policies []struct {
		ID                 int    `json:"id"`
		IncidentPreference string `json:"incident_preference"`
		Name               string `json:"name"`
		CreatedAt          int64  `json:"created_at"`
		UpdatedAt          int64  `json:"updated_at"`
	} `json:"policies"`
	Conditions []struct {
		ID             int      `json:"id"`
		Type           string   `json:"type"`
		Name           string   `json:"name"`
		Enabled        bool     `json:"enabled"`
		Entities       []string `json:"entities"`
		Metric         string   `json:"metric"`
		ConditionScope string   `json:"condition_scope"`
		Terms          []struct {
			Duration     string `json:"duration"`
			Operator     string `json:"operator"`
			Priority     string `json:"priority"`
			Threshold    string `json:"threshold"`
			TimeFunction string `json:"time_function"`
		} `json:"terms"`
		UserDefined struct {
			Metric        string `json:"metric"`
			ValueFunction string `json:"value_function"`
		} `json:"user_defined"`
	} `json:"conditions"`
	NrqlCondition []struct {
		Type                      string `json:"type"`
		ID                        string `json:"id"`
		Name                      string `json:"name"`
		RunbookURL                string `json:"runbook_url"`
		Enabled                   string `json:"enabled"`
		ExpectedGroups            string `json:"expected_groups"`
		IgnoreOverlap             string `json:"ignore_overlap"`
		ValueFunction             string `json:"value_function"`
		ViolationTimeLimitSeconds string `json:"violation_time_limit_seconds"`
		Terms                     []struct {
			Duration     string `json:"duration"`
			Operator     string `json:"operator"`
			Priority     string `json:"priority"`
			Threshold    string `json:"threshold"`
			TimeFunction string `json:"time_function"`
		} `json:"terms"`
		Nrql struct {
			Query      string `json:"query"`
			SinceValue string `json:"since_value"`
		} `json:"nrql"`
	}
}
