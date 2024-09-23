package encodings

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InputDefinition = InputFile{}

type InputFile struct {
	Filename *string `json:"filename,omitempty"`

	// Fields inherited from InputDefinition

	IncludedTracks *[]TrackDescriptor `json:"includedTracks,omitempty"`
	OdataType      string             `json:"@odata.type"`
}

func (s InputFile) InputDefinition() BaseInputDefinitionImpl {
	return BaseInputDefinitionImpl{
		IncludedTracks: s.IncludedTracks,
		OdataType:      s.OdataType,
	}
}

var _ json.Marshaler = InputFile{}

func (s InputFile) MarshalJSON() ([]byte, error) {
	type wrapper InputFile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InputFile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InputFile: %+v", err)
	}

	decoded["@odata.type"] = "#Microsoft.Media.InputFile"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InputFile: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &InputFile{}

func (s *InputFile) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Filename  *string `json:"filename,omitempty"`
		OdataType string  `json:"@odata.type"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Filename = decoded.Filename
	s.OdataType = decoded.OdataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling InputFile into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'IncludedTracks' for 'InputFile': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IncludedTracks = &output
	}

	return nil
}
