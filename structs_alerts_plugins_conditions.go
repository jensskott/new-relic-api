package new_relic_api

type AlertsPluginsConditions struct {
	PluginsConditions []struct {
		ID                string   `json:"id"`
		Name              string   `json:"name"`
		Enabled           string   `json:"enabled"`
		Entities          []string `json:"entities"`
		MetricDescription string   `json:"metric_description"`
		Metric            string   `json:"metric"`
		ValueFunction     string   `json:"value_function"`
		RunbookURL        string   `json:"runbook_url"`
		Terms             []struct {
			Duration     string `json:"duration"`
			Operator     string `json:"operator"`
			Priority     string `json:"priority"`
			Threshold    string `json:"threshold"`
			TimeFunction string `json:"time_function"`
		} `json:"terms"`
		Plugin struct {
			ID   string `json:"id"`
			GUID string `json:"guid"`
		} `json:"plugin"`
	} `json:"plugins_condition"`
}

type AlertsPluginsConditionsPayload struct {
	PluginsCondition PluginsCondition `json:"plugins_condition"`
}
type AlertsPluginsConditionsTerms struct {
	Duration     string `json:"duration"`
	Operator     string `json:"operator"`
	Priority     string `json:"priority"`
	Threshold    string `json:"threshold"`
	TimeFunction string `json:"time_function"`
}
type Plugin struct {
	ID   string `json:"id"`
	GUID string `json:"guid"`
}
type PluginsCondition struct {
	Name              string                         `json:"name"`
	Enabled           string                         `json:"enabled"`
	Entities          []string                       `json:"entities"`
	MetricDescription string                         `json:"metric_description"`
	Metric            string                         `json:"metric"`
	ValueFunction     string                         `json:"value_function"`
	RunbookURL        string                         `json:"runbook_url"`
	Terms             []AlertsPluginsConditionsTerms `json:"terms"`
	Plugin            Plugin                         `json:"plugin"`
}
