package new_relic_api

type AlertsConditions struct {
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
	} `json:"conditions"`
}

type AlertsConditionsPayload struct {
	Condition Condition `json:"condition"`
}
type Terms struct {
	Duration     string `json:"duration"`
	Operator     string `json:"operator"`
	Priority     string `json:"priority"`
	Threshold    string `json:"threshold"`
	TimeFunction string `json:"time_function"`
}
type UserDefined struct {
	Metric        string `json:"metric"`
	ValueFunction string `json:"value_function"`
}
type Condition struct {
	Type                string      `json:"type"`
	Name                string      `json:"name"`
	Enabled             string      `json:"enabled"`
	Entities            []string    `json:"entities"`
	Metric              string      `json:"metric"`
	GcMetric            string      `json:"gc_metric"`
	ConditionScope      string      `json:"condition_scope"`
	ViolationCloseTimer string      `json:"violation_close_timer"`
	Terms               []Terms     `json:"terms"`
	UserDefined         UserDefined `json:"user_defined"`
}
