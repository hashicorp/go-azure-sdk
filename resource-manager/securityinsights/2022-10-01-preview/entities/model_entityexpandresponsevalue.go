package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityExpandResponseValue struct {
	Edges    *[]EntityEdges `json:"edges,omitempty"`
	Entities *[]Entity      `json:"entities,omitempty"`
}

var _ json.Unmarshaler = &EntityExpandResponseValue{}

func (s *EntityExpandResponseValue) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Edges *[]EntityEdges `json:"edges,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Edges = decoded.Edges

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EntityExpandResponseValue into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["entities"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Entities into list []json.RawMessage: %+v", err)
		}

		output := make([]Entity, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalEntityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Entities' for 'EntityExpandResponseValue': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Entities = &output
	}

	return nil
}
