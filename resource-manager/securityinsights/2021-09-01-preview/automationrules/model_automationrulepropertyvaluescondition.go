package automationrules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AutomationRuleCondition = AutomationRulePropertyValuesCondition{}

type AutomationRulePropertyValuesCondition struct {
	ConditionProperties AutomationRulePropertyValuesConditionConditionProperties `json:"conditionProperties"`

	// Fields inherited from AutomationRuleCondition
}

var _ json.Marshaler = AutomationRulePropertyValuesCondition{}

func (s AutomationRulePropertyValuesCondition) MarshalJSON() ([]byte, error) {
	type wrapper AutomationRulePropertyValuesCondition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AutomationRulePropertyValuesCondition: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AutomationRulePropertyValuesCondition: %+v", err)
	}
	decoded["conditionType"] = "Property"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AutomationRulePropertyValuesCondition: %+v", err)
	}

	return encoded, nil
}
