package namespacediscoveredassets

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceDiscoveredEvent struct {
	DataSource         *string             `json:"dataSource,omitempty"`
	Destinations       *[]EventDestination `json:"destinations,omitempty"`
	EventConfiguration *string             `json:"eventConfiguration,omitempty"`
	LastUpdatedOn      *string             `json:"lastUpdatedOn,omitempty"`
	Name               string              `json:"name"`
	TypeRef            *string             `json:"typeRef,omitempty"`
}

func (o *NamespaceDiscoveredEvent) GetLastUpdatedOnAsTime() (*time.Time, error) {
	if o.LastUpdatedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *NamespaceDiscoveredEvent) SetLastUpdatedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedOn = &formatted
}

var _ json.Unmarshaler = &NamespaceDiscoveredEvent{}

func (s *NamespaceDiscoveredEvent) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DataSource         *string `json:"dataSource,omitempty"`
		EventConfiguration *string `json:"eventConfiguration,omitempty"`
		LastUpdatedOn      *string `json:"lastUpdatedOn,omitempty"`
		Name               string  `json:"name"`
		TypeRef            *string `json:"typeRef,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DataSource = decoded.DataSource
	s.EventConfiguration = decoded.EventConfiguration
	s.LastUpdatedOn = decoded.LastUpdatedOn
	s.Name = decoded.Name
	s.TypeRef = decoded.TypeRef

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NamespaceDiscoveredEvent into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Destinations' for 'NamespaceDiscoveredEvent': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Destinations = &output
	}

	return nil
}
