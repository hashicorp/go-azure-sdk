package namespaceassets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceEvent struct {
	DataSource         *string             `json:"dataSource,omitempty"`
	Destinations       *[]EventDestination `json:"destinations,omitempty"`
	EventConfiguration *string             `json:"eventConfiguration,omitempty"`
	Name               string              `json:"name"`
	TypeRef            *string             `json:"typeRef,omitempty"`
}

var _ json.Unmarshaler = &NamespaceEvent{}

func (s *NamespaceEvent) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DataSource         *string `json:"dataSource,omitempty"`
		EventConfiguration *string `json:"eventConfiguration,omitempty"`
		Name               string  `json:"name"`
		TypeRef            *string `json:"typeRef,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DataSource = decoded.DataSource
	s.EventConfiguration = decoded.EventConfiguration
	s.Name = decoded.Name
	s.TypeRef = decoded.TypeRef

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NamespaceEvent into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["destinations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Destinations into list []json.RawMessage: %+v", err)
		}

		output := make([]EventDestination, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalEventDestinationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Destinations' for 'NamespaceEvent': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Destinations = &output
	}

	return nil
}
