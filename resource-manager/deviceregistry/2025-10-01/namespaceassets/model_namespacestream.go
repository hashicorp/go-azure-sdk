package namespaceassets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceStream struct {
	Destinations        *[]StreamDestination `json:"destinations,omitempty"`
	Name                string               `json:"name"`
	StreamConfiguration *string              `json:"streamConfiguration,omitempty"`
	TypeRef             *string              `json:"typeRef,omitempty"`
}

var _ json.Unmarshaler = &NamespaceStream{}

func (s *NamespaceStream) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Name                string  `json:"name"`
		StreamConfiguration *string `json:"streamConfiguration,omitempty"`
		TypeRef             *string `json:"typeRef,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Name = decoded.Name
	s.StreamConfiguration = decoded.StreamConfiguration
	s.TypeRef = decoded.TypeRef

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NamespaceStream into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["destinations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Destinations into list []json.RawMessage: %+v", err)
		}

		output := make([]StreamDestination, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalStreamDestinationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Destinations' for 'NamespaceStream': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Destinations = &output
	}

	return nil
}
