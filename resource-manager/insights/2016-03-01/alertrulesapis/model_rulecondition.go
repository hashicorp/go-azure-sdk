package alertrulesapis

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuleCondition interface {
	RuleCondition() BaseRuleConditionImpl
}

var _ RuleCondition = BaseRuleConditionImpl{}

type BaseRuleConditionImpl struct {
	DataSource RuleDataSource `json:"dataSource"`
	OdataType  string         `json:"odata.type"`
}

func (s BaseRuleConditionImpl) RuleCondition() BaseRuleConditionImpl {
	return s
}

var _ RuleCondition = RawRuleConditionImpl{}

// RawRuleConditionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRuleConditionImpl struct {
	ruleCondition BaseRuleConditionImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawRuleConditionImpl) RuleCondition() BaseRuleConditionImpl {
	return s.ruleCondition
}

var _ json.Unmarshaler = &BaseRuleConditionImpl{}

func (s *BaseRuleConditionImpl) UnmarshalJSON(bytes []byte) error {
	type alias BaseRuleConditionImpl
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into BaseRuleConditionImpl: %+v", err)
	}

	s.OdataType = decoded.OdataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseRuleConditionImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["dataSource"]; ok {
		impl, err := UnmarshalRuleDataSourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DataSource' for 'BaseRuleConditionImpl': %+v", err)
		}
		s.DataSource = impl
	}
	return nil
}

func UnmarshalRuleConditionImplementation(input []byte) (RuleCondition, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RuleCondition into map[string]interface: %+v", err)
	}

	value, ok := temp["odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Microsoft.Azure.Management.Insights.Models.LocationThresholdRuleCondition") {
		var out LocationThresholdRuleCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LocationThresholdRuleCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Microsoft.Azure.Management.Insights.Models.ManagementEventRuleCondition") {
		var out ManagementEventRuleCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagementEventRuleCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Microsoft.Azure.Management.Insights.Models.ThresholdRuleCondition") {
		var out ThresholdRuleCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ThresholdRuleCondition: %+v", err)
		}
		return out, nil
	}

	var parent BaseRuleConditionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRuleConditionImpl: %+v", err)
	}

	return RawRuleConditionImpl{
		ruleCondition: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
