package externalsecuritysolutions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalSecuritySolution interface {
	ExternalSecuritySolution() BaseExternalSecuritySolutionImpl
}

var _ ExternalSecuritySolution = BaseExternalSecuritySolutionImpl{}

type BaseExternalSecuritySolutionImpl struct {
	Id       *string                       `json:"id,omitempty"`
	Kind     *ExternalSecuritySolutionKind `json:"kind,omitempty"`
	Location *string                       `json:"location,omitempty"`
	Name     *string                       `json:"name,omitempty"`
	Type     *string                       `json:"type,omitempty"`
}

func (s BaseExternalSecuritySolutionImpl) ExternalSecuritySolution() BaseExternalSecuritySolutionImpl {
	return s
}

var _ ExternalSecuritySolution = RawExternalSecuritySolutionImpl{}

// RawExternalSecuritySolutionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawExternalSecuritySolutionImpl struct {
	externalSecuritySolution BaseExternalSecuritySolutionImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawExternalSecuritySolutionImpl) ExternalSecuritySolution() BaseExternalSecuritySolutionImpl {
	return s.externalSecuritySolution
}

func UnmarshalExternalSecuritySolutionImplementation(input []byte) (ExternalSecuritySolution, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalSecuritySolution into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AAD") {
		var out AadExternalSecuritySolution
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AadExternalSecuritySolution: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ATA") {
		var out AtaExternalSecuritySolution
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AtaExternalSecuritySolution: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "CEF") {
		var out CefExternalSecuritySolution
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CefExternalSecuritySolution: %+v", err)
		}
		return out, nil
	}

	var parent BaseExternalSecuritySolutionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseExternalSecuritySolutionImpl: %+v", err)
	}

	return RawExternalSecuritySolutionImpl{
		externalSecuritySolution: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
