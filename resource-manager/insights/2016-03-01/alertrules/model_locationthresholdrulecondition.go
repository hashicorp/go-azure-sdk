package alertrules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RuleCondition = LocationThresholdRuleCondition{}

type LocationThresholdRuleCondition struct {
	FailedLocationCount int64   `json:"failedLocationCount"`
	WindowSize          *string `json:"windowSize,omitempty"`

	// Fields inherited from RuleCondition
	DataSource *RuleDataSource `json:"dataSource,omitempty"`
}

var _ json.Marshaler = LocationThresholdRuleCondition{}

func (s LocationThresholdRuleCondition) MarshalJSON() ([]byte, error) {
	type wrapper LocationThresholdRuleCondition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LocationThresholdRuleCondition: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LocationThresholdRuleCondition: %+v", err)
	}
	decoded["odata.type"] = "Microsoft.Azure.Management.Insights.Models.LocationThresholdRuleCondition"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LocationThresholdRuleCondition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &LocationThresholdRuleCondition{}

func (s *LocationThresholdRuleCondition) UnmarshalJSON(bytes []byte) error {
	type alias LocationThresholdRuleCondition
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into LocationThresholdRuleCondition: %+v", err)
	}

	s.FailedLocationCount = decoded.FailedLocationCount
	s.WindowSize = decoded.WindowSize

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling LocationThresholdRuleCondition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["dataSource"]; ok {
		impl, err := unmarshalRuleDataSourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DataSource' for 'LocationThresholdRuleCondition': %+v", err)
		}
		s.DataSource = &impl
	}
	return nil
}
