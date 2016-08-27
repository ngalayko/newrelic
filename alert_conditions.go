package newrelic

type AlertCondition struct {
	Id          int                  `json:"id,omitempty"`
	Type        string               `json:"type,omitempty"`
	Name        string               `json:"name,omitempty"`
	Enabled     bool                 `json:"name,omitempty"`
	Entities    []string             `json:"entities,omitempty"`
	Metric      string               `json:"metric,omitempty"`
	RunbookUrl  string               `json:"runbook_url,omitempty"`
	Terms       []AlertConditionTerm `json:"terms,omitempty"`
	UserDefined AlertUserDefined     `json:"user_defined,omitempty"`
}

type AlertConditionTerm struct {
	Duration     string `json:"duration,omitempty"`
	Operator     string `json:"operator,omitempty"`
	Priority     string `json:"priority,omitempty"`
	Threshold    string `json:"threshold,omitempty"`
	TimeFunction string `json:"time_function,omitempty"`
}

type AlertUserDefined struct {
	Metric        string `json:"metric,omitempty"`
	ValueFunction string `json:"value_function,omitempty"`
}

type AlertConditionOptions struct {
	policyId int `json:"policy_id,omitempty"`
	Page     int `json:"page,omitempty"`
}

type AlertConditionResponse struct {
	Conditions []AlertCondition `json:"conditions,omitempty"`
}

func (c *Client) GetAlertConditions(policy int, options *AlertConditionOptions) (*AlertConditionResponse, error) {
	conditions := &AlertConditionResponse{}
	options.policyId = policy
	err := c.doGet("/alerts_conditions.json", options.encode(), conditions)
	if err != nil {
		return nil, err
	}
	return conditions, nil
}

func (o *AlertConditionOptions) encode() string {
	return encodeGetParams(map[string]interface{}{
		"policy_id": o.policyId,
		"page":      o.Page,
	})
}