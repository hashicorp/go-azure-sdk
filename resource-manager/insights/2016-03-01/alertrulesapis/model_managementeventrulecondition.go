package alertrulesapis

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RuleCondition = ManagementEventRuleCondition{}

type ManagementEventRuleCondition struct {
	Aggregation *ManagementEventAggregationCondition `json:"aggregation,omitempty"`

	// Fields inherited from RuleCondition
	DataSource RuleDataSource `json:"dataSource"`
}

var _ json.Marshaler = ManagementEventRuleCondition{}

func (s ManagementEventRuleCondition) MarshalJSON() ([]byte, error) {
	type wrapper ManagementEventRuleCondition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagementEventRuleCondition: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagementEventRuleCondition: %+v", err)
	}
	decoded["odata.type"] = "Microsoft.Azure.Management.Insights.Models.ManagementEventRuleCondition"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagementEventRuleCondition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ManagementEventRuleCondition{}

func (s *ManagementEventRuleCondition) UnmarshalJSON(bytes []byte) error {
	type alias ManagementEventRuleCondition
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ManagementEventRuleCondition: %+v", err)
	}

	s.Aggregation = decoded.Aggregation

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ManagementEventRuleCondition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["dataSource"]; ok {
		impl, err := unmarshalRuleDataSourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DataSource' for 'ManagementEventRuleCondition': %+v", err)
		}
		s.DataSource = impl
	}
	return nil
}
