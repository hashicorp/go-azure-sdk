package pools

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentProfile interface {
}

// RawAgentProfileImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAgentProfileImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalAgentProfileImplementation(input []byte) (AgentProfile, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AgentProfile into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Stateful") {
		var out Stateful
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Stateful: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Stateless") {
		var out StatelessAgentProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StatelessAgentProfile: %+v", err)
		}
		return out, nil
	}

	out := RawAgentProfileImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
