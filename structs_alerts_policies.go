package new_relic_api

type AlertsPolicies struct {
	Policies []struct {
		ID                 int    `json:"id"`
		IncidentPreference string `json:"incident_preference"`
		Name               string `json:"name"`
		CreatedAt          int64  `json:"created_at"`
		UpdatedAt          int64  `json:"updated_at"`
	} `json:"policies"`
}

type AlertsPoliciesPayload struct {
	Policy Policy `json:"policy"`
}

type Policy struct {
	IncidentPreference string `json:"incident_preference"`
	Name               string `json:"name"`
}
