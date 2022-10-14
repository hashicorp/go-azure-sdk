package scripts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptExecutionParameter interface {
}

func unmarshalScriptExecutionParameterImplementation(input []byte) (ScriptExecutionParameter, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ScriptExecutionParameter into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Credential") {
		var out PSCredentialExecutionParameter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PSCredentialExecutionParameter: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SecureValue") {
		var out ScriptSecureStringExecutionParameter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ScriptSecureStringExecutionParameter: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Value") {
		var out ScriptStringExecutionParameter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ScriptStringExecutionParameter: %+v", err)
		}
		return out, nil
	}

	type RawScriptExecutionParameterImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawScriptExecutionParameterImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
