package alertrules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuleCondition interface {
}

// RawRuleConditionImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRuleConditionImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalRuleConditionImplementation(input []byte) (RuleCondition, error) {
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

	out := RawRuleConditionImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
