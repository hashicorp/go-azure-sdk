package integrationruntimes

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSetupBase interface {
}

// RawCustomSetupBaseImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCustomSetupBaseImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalCustomSetupBaseImplementation(input []byte) (CustomSetupBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomSetupBase into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzPowerShellSetup") {
		var out AzPowerShellSetup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzPowerShellSetup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "CmdkeySetup") {
		var out CmdkeySetup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CmdkeySetup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ComponentSetup") {
		var out ComponentSetup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ComponentSetup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "EnvironmentVariableSetup") {
		var out EnvironmentVariableSetup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnvironmentVariableSetup: %+v", err)
		}
		return out, nil
	}

	out := RawCustomSetupBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
