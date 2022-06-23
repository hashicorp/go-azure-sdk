package automationrules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationRuleCondition interface {
}

func unmarshalAutomationRuleConditionImplementation(input []byte) (AutomationRuleCondition, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AutomationRuleCondition into map[string]interface: %+v", err)
	}

	value, ok := temp["conditionType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Property") {
		var out AutomationRulePropertyValuesCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AutomationRulePropertyValuesCondition: %+v", err)
		}
		return out, nil
	}

	type RawAutomationRuleConditionImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawAutomationRuleConditionImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
