package encodings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InputDefinition interface {
	InputDefinition() BaseInputDefinitionImpl
}

var _ InputDefinition = BaseInputDefinitionImpl{}

type BaseInputDefinitionImpl struct {
	IncludedTracks *[]TrackDescriptor `json:"includedTracks,omitempty"`
	OdataType      string             `json:"@odata.type"`
}

func (s BaseInputDefinitionImpl) InputDefinition() BaseInputDefinitionImpl {
	return s
}

var _ InputDefinition = RawInputDefinitionImpl{}

// RawInputDefinitionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawInputDefinitionImpl struct {
	inputDefinition BaseInputDefinitionImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawInputDefinitionImpl) InputDefinition() BaseInputDefinitionImpl {
	return s.inputDefinition
}

var _ json.Unmarshaler = &BaseInputDefinitionImpl{}

func (s *BaseInputDefinitionImpl) UnmarshalJSON(bytes []byte) error {
	type alias BaseInputDefinitionImpl
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into BaseInputDefinitionImpl: %+v", err)
	}

	s.OdataType = decoded.OdataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseInputDefinitionImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["includedTracks"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling IncludedTracks into list []json.RawMessage: %+v", err)
		}

		output := make([]TrackDescriptor, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalTrackDescriptorImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'IncludedTracks' for 'BaseInputDefinitionImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IncludedTracks = &output
	}
	return nil
}

func UnmarshalInputDefinitionImplementation(input []byte) (InputDefinition, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling InputDefinition into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.FromAllInputFile") {
		var out FromAllInputFile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FromAllInputFile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.FromEachInputFile") {
		var out FromEachInputFile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FromEachInputFile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.InputFile") {
		var out InputFile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InputFile: %+v", err)
		}
		return out, nil
	}

	var parent BaseInputDefinitionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseInputDefinitionImpl: %+v", err)
	}

	return RawInputDefinitionImpl{
		inputDefinition: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
