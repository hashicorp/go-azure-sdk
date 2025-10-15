package namespacediscoveredassets

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceDiscoveredStream struct {
	Destinations        *[]StreamDestination `json:"destinations,omitempty"`
	LastUpdatedOn       *string              `json:"lastUpdatedOn,omitempty"`
	Name                string               `json:"name"`
	StreamConfiguration *string              `json:"streamConfiguration,omitempty"`
	TypeRef             *string              `json:"typeRef,omitempty"`
}

func (o *NamespaceDiscoveredStream) GetLastUpdatedOnAsTime() (*time.Time, error) {
	if o.LastUpdatedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *NamespaceDiscoveredStream) SetLastUpdatedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedOn = &formatted
}

var _ json.Unmarshaler = &NamespaceDiscoveredStream{}

func (s *NamespaceDiscoveredStream) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		LastUpdatedOn       *string `json:"lastUpdatedOn,omitempty"`
		Name                string  `json:"name"`
		StreamConfiguration *string `json:"streamConfiguration,omitempty"`
		TypeRef             *string `json:"typeRef,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.LastUpdatedOn = decoded.LastUpdatedOn
	s.Name = decoded.Name
	s.StreamConfiguration = decoded.StreamConfiguration
	s.TypeRef = decoded.TypeRef

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NamespaceDiscoveredStream into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Destinations' for 'NamespaceDiscoveredStream': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Destinations = &output
	}

	return nil
}
