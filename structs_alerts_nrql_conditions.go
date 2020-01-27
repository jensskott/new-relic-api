package new_relic_api

type AlertsNrqlConditions struct {
	NrqlConditions []struct {
		Type          string `json:"type"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Enabled       bool   `json:"enabled"`
		ValueFunction string `json:"value_function"`
		Terms         []struct {
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
	} `json:"nrql_conditions"`
}

type AlertsNrqlConditionsPayload struct {
	NrqlCondition NrqlCondition `json:"nrql_condition"`
}

type AlertsNrqlConditionsTerms struct {
	Duration     string `json:"duration"`
	Operator     string `json:"operator"`
	Priority     string `json:"priority"`
	Threshold    string `json:"threshold"`
	TimeFunction string `json:"time_function"`
}

type Nrql struct {
	Query      string `json:"query"`
	SinceValue string `json:"since_value"`
}

type NrqlCondition struct {
	Name                      string                      `json:"name"`
	RunbookURL                string                      `json:"runbook_url"`
	Enabled                   string                      `json:"enabled"`
	ExpectedGroups            string                      `json:"expected_groups"`
	IgnoreOverlap             string                      `json:"ignore_overlap"`
	ValueFunction             string                      `json:"value_function"`
	ViolationTimeLimitSeconds string                      `json:"violation_time_limit_seconds"`
	Terms                     []AlertsNrqlConditionsTerms `json:"terms"`
	Nrql                      Nrql                        `json:"nrql"`
}
