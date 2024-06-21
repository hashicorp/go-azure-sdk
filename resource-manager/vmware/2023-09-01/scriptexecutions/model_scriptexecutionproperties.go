package scriptexecutions

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptExecutionProperties struct {
	Errors            *[]string                         `json:"errors,omitempty"`
	FailureReason     *string                           `json:"failureReason,omitempty"`
	FinishedAt        *string                           `json:"finishedAt,omitempty"`
	HiddenParameters  *[]ScriptExecutionParameter       `json:"hiddenParameters,omitempty"`
	Information       *[]string                         `json:"information,omitempty"`
	NamedOutputs      *map[string]interface{}           `json:"namedOutputs,omitempty"`
	Output            *[]string                         `json:"output,omitempty"`
	Parameters        *[]ScriptExecutionParameter       `json:"parameters,omitempty"`
	ProvisioningState *ScriptExecutionProvisioningState `json:"provisioningState,omitempty"`
	Retention         *string                           `json:"retention,omitempty"`
	ScriptCmdletId    *string                           `json:"scriptCmdletId,omitempty"`
	StartedAt         *string                           `json:"startedAt,omitempty"`
	SubmittedAt       *string                           `json:"submittedAt,omitempty"`
	Timeout           string                            `json:"timeout"`
	Warnings          *[]string                         `json:"warnings,omitempty"`
}

var _ json.Unmarshaler = &ScriptExecutionProperties{}

func (s *ScriptExecutionProperties) UnmarshalJSON(bytes []byte) error {
	type alias ScriptExecutionProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ScriptExecutionProperties: %+v", err)
	}

	s.Errors = decoded.Errors
	s.FailureReason = decoded.FailureReason
	s.FinishedAt = decoded.FinishedAt
	s.Information = decoded.Information
	s.NamedOutputs = decoded.NamedOutputs
	s.Output = decoded.Output
	s.ProvisioningState = decoded.ProvisioningState
	s.Retention = decoded.Retention
	s.ScriptCmdletId = decoded.ScriptCmdletId
	s.StartedAt = decoded.StartedAt
	s.SubmittedAt = decoded.SubmittedAt
	s.Timeout = decoded.Timeout
	s.Warnings = decoded.Warnings

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ScriptExecutionProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["hiddenParameters"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling HiddenParameters into list []json.RawMessage: %+v", err)
		}

		output := make([]ScriptExecutionParameter, 0)
		for i, val := range listTemp {
			impl, err := unmarshalScriptExecutionParameterImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'HiddenParameters' for 'ScriptExecutionProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.HiddenParameters = &output
	}

	if v, ok := temp["parameters"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Parameters into list []json.RawMessage: %+v", err)
		}

		output := make([]ScriptExecutionParameter, 0)
		for i, val := range listTemp {
			impl, err := unmarshalScriptExecutionParameterImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Parameters' for 'ScriptExecutionProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Parameters = &output
	}
	return nil
}
