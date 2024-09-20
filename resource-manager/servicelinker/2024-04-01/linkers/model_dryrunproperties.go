package linkers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DryrunProperties struct {
	OperationPreviews   *[]DryrunOperationPreview   `json:"operationPreviews,omitempty"`
	Parameters          DryrunParameters            `json:"parameters"`
	PrerequisiteResults *[]DryrunPrerequisiteResult `json:"prerequisiteResults,omitempty"`
	ProvisioningState   *string                     `json:"provisioningState,omitempty"`
}

var _ json.Unmarshaler = &DryrunProperties{}

func (s *DryrunProperties) UnmarshalJSON(bytes []byte) error {
	type alias DryrunProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into DryrunProperties: %+v", err)
	}

	s.OperationPreviews = decoded.OperationPreviews
	s.ProvisioningState = decoded.ProvisioningState

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DryrunProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["parameters"]; ok {
		impl, err := UnmarshalDryrunParametersImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Parameters' for 'DryrunProperties': %+v", err)
		}
		s.Parameters = impl
	}

	if v, ok := temp["prerequisiteResults"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling PrerequisiteResults into list []json.RawMessage: %+v", err)
		}

		output := make([]DryrunPrerequisiteResult, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDryrunPrerequisiteResultImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'PrerequisiteResults' for 'DryrunProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PrerequisiteResults = &output
	}
	return nil
}
