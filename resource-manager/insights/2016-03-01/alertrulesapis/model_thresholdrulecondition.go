package alertrulesapis

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RuleCondition = ThresholdRuleCondition{}

type ThresholdRuleCondition struct {
	Operator        ConditionOperator        `json:"operator"`
	Threshold       float64                  `json:"threshold"`
	TimeAggregation *TimeAggregationOperator `json:"timeAggregation,omitempty"`
	WindowSize      *string                  `json:"windowSize,omitempty"`

	// Fields inherited from RuleCondition
	DataSource RuleDataSource `json:"dataSource"`
}

var _ json.Marshaler = ThresholdRuleCondition{}

func (s ThresholdRuleCondition) MarshalJSON() ([]byte, error) {
	type wrapper ThresholdRuleCondition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ThresholdRuleCondition: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ThresholdRuleCondition: %+v", err)
	}
	decoded["odata.type"] = "Microsoft.Azure.Management.Insights.Models.ThresholdRuleCondition"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ThresholdRuleCondition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ThresholdRuleCondition{}

func (s *ThresholdRuleCondition) UnmarshalJSON(bytes []byte) error {
	type alias ThresholdRuleCondition
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ThresholdRuleCondition: %+v", err)
	}

	s.Operator = decoded.Operator
	s.Threshold = decoded.Threshold
	s.TimeAggregation = decoded.TimeAggregation
	s.WindowSize = decoded.WindowSize

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ThresholdRuleCondition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["dataSource"]; ok {
		impl, err := unmarshalRuleDataSourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DataSource' for 'ThresholdRuleCondition': %+v", err)
		}
		s.DataSource = impl
	}
	return nil
}
