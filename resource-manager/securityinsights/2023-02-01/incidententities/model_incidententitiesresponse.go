package incidententities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentEntitiesResponse struct {
	Entities *[]Entity                          `json:"entities,omitempty"`
	MetaData *[]IncidentEntitiesResultsMetadata `json:"metaData,omitempty"`
}

var _ json.Unmarshaler = &IncidentEntitiesResponse{}

func (s *IncidentEntitiesResponse) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		MetaData *[]IncidentEntitiesResultsMetadata `json:"metaData,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.MetaData = decoded.MetaData

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IncidentEntitiesResponse into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Entities' for 'IncidentEntitiesResponse': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Entities = &output
	}

	return nil
}
