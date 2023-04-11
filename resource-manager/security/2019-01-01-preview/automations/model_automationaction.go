package automations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationAction interface {
}

func unmarshalAutomationActionImplementation(input []byte) (AutomationAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AutomationAction into map[string]interface: %+v", err)
	}

	value, ok := temp["actionType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "EventHub") {
		var out AutomationActionEventHub
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AutomationActionEventHub: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "LogicApp") {
		var out AutomationActionLogicApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AutomationActionLogicApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Workspace") {
		var out AutomationActionWorkspace
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AutomationActionWorkspace: %+v", err)
		}
		return out, nil
	}

	type RawAutomationActionImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawAutomationActionImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
