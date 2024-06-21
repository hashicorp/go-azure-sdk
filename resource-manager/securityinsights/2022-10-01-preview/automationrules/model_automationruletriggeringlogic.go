package automationrules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationRuleTriggeringLogic struct {
	Conditions        *[]AutomationRuleCondition `json:"conditions,omitempty"`
	ExpirationTimeUtc *string                    `json:"expirationTimeUtc,omitempty"`
	IsEnabled         bool                       `json:"isEnabled"`
	TriggersOn        TriggersOn                 `json:"triggersOn"`
	TriggersWhen      TriggersWhen               `json:"triggersWhen"`
}

var _ json.Unmarshaler = &AutomationRuleTriggeringLogic{}

func (s *AutomationRuleTriggeringLogic) UnmarshalJSON(bytes []byte) error {
	type alias AutomationRuleTriggeringLogic
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into AutomationRuleTriggeringLogic: %+v", err)
	}

	s.ExpirationTimeUtc = decoded.ExpirationTimeUtc
	s.IsEnabled = decoded.IsEnabled
	s.TriggersOn = decoded.TriggersOn
	s.TriggersWhen = decoded.TriggersWhen

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AutomationRuleTriggeringLogic into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["conditions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Conditions into list []json.RawMessage: %+v", err)
		}

		output := make([]AutomationRuleCondition, 0)
		for i, val := range listTemp {
			impl, err := unmarshalAutomationRuleConditionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Conditions' for 'AutomationRuleTriggeringLogic': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Conditions = &output
	}
	return nil
}
