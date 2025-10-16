package namespaceassets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceDataset struct {
	DataPoints           *[]NamespaceDatasetDataPoint `json:"dataPoints,omitempty"`
	DataSource           *string                      `json:"dataSource,omitempty"`
	DatasetConfiguration *string                      `json:"datasetConfiguration,omitempty"`
	Destinations         *[]DatasetDestination        `json:"destinations,omitempty"`
	Name                 string                       `json:"name"`
	TypeRef              *string                      `json:"typeRef,omitempty"`
}

var _ json.Unmarshaler = &NamespaceDataset{}

func (s *NamespaceDataset) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DataPoints           *[]NamespaceDatasetDataPoint `json:"dataPoints,omitempty"`
		DataSource           *string                      `json:"dataSource,omitempty"`
		DatasetConfiguration *string                      `json:"datasetConfiguration,omitempty"`
		Name                 string                       `json:"name"`
		TypeRef              *string                      `json:"typeRef,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DataPoints = decoded.DataPoints
	s.DataSource = decoded.DataSource
	s.DatasetConfiguration = decoded.DatasetConfiguration
	s.Name = decoded.Name
	s.TypeRef = decoded.TypeRef

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NamespaceDataset into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["destinations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Destinations into list []json.RawMessage: %+v", err)
		}

		output := make([]DatasetDestination, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDatasetDestinationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Destinations' for 'NamespaceDataset': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Destinations = &output
	}

	return nil
}
