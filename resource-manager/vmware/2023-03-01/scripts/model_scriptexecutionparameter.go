package scripts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptExecutionParameter interface {
	ScriptExecutionParameter() BaseScriptExecutionParameterImpl
}

var _ ScriptExecutionParameter = BaseScriptExecutionParameterImpl{}

type BaseScriptExecutionParameterImpl struct {
	Name string                       `json:"name"`
	Type ScriptExecutionParameterType `json:"type"`
}

func (s BaseScriptExecutionParameterImpl) ScriptExecutionParameter() BaseScriptExecutionParameterImpl {
	return s
}

var _ ScriptExecutionParameter = RawScriptExecutionParameterImpl{}

// RawScriptExecutionParameterImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawScriptExecutionParameterImpl struct {
	scriptExecutionParameter BaseScriptExecutionParameterImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawScriptExecutionParameterImpl) ScriptExecutionParameter() BaseScriptExecutionParameterImpl {
	return s.scriptExecutionParameter
}

func UnmarshalScriptExecutionParameterImplementation(input []byte) (ScriptExecutionParameter, error) {
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

	var parent BaseScriptExecutionParameterImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseScriptExecutionParameterImpl: %+v", err)
	}

	return RawScriptExecutionParameterImpl{
		scriptExecutionParameter: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
