package namespacediscoveredassets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceDiscoveredEventGroup struct {
	DataSource              *string                     `json:"dataSource,omitempty"`
	DefaultDestinations     *[]EventDestination         `json:"defaultDestinations,omitempty"`
	EventGroupConfiguration *string                     `json:"eventGroupConfiguration,omitempty"`
	Events                  *[]NamespaceDiscoveredEvent `json:"events,omitempty"`
	Name                    string                      `json:"name"`
	TypeRef                 *string                     `json:"typeRef,omitempty"`
}

var _ json.Unmarshaler = &NamespaceDiscoveredEventGroup{}

func (s *NamespaceDiscoveredEventGroup) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DataSource              *string                     `json:"dataSource,omitempty"`
		EventGroupConfiguration *string                     `json:"eventGroupConfiguration,omitempty"`
		Events                  *[]NamespaceDiscoveredEvent `json:"events,omitempty"`
		Name                    string                      `json:"name"`
		TypeRef                 *string                     `json:"typeRef,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DataSource = decoded.DataSource
	s.EventGroupConfiguration = decoded.EventGroupConfiguration
	s.Events = decoded.Events
	s.Name = decoded.Name
	s.TypeRef = decoded.TypeRef

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NamespaceDiscoveredEventGroup into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["defaultDestinations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DefaultDestinations into list []json.RawMessage: %+v", err)
		}

		output := make([]EventDestination, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalEventDestinationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DefaultDestinations' for 'NamespaceDiscoveredEventGroup': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DefaultDestinations = &output
	}

	return nil
}
