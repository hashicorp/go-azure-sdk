package dataflowgraph

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowGraphProperties struct {
	HealthState            *ResourceHealthState          `json:"healthState,omitempty"`
	Mode                   *OperationalMode              `json:"mode,omitempty"`
	NodeConnections        []DataflowGraphNodeConnection `json:"nodeConnections"`
	Nodes                  []DataflowGraphNode           `json:"nodes"`
	ProvisioningState      *ProvisioningState            `json:"provisioningState,omitempty"`
	RequestDiskPersistence *OperationalMode              `json:"requestDiskPersistence,omitempty"`
}

var _ json.Unmarshaler = &DataflowGraphProperties{}

func (s *DataflowGraphProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		HealthState            *ResourceHealthState          `json:"healthState,omitempty"`
		Mode                   *OperationalMode              `json:"mode,omitempty"`
		NodeConnections        []DataflowGraphNodeConnection `json:"nodeConnections"`
		ProvisioningState      *ProvisioningState            `json:"provisioningState,omitempty"`
		RequestDiskPersistence *OperationalMode              `json:"requestDiskPersistence,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.HealthState = decoded.HealthState
	s.Mode = decoded.Mode
	s.NodeConnections = decoded.NodeConnections
	s.ProvisioningState = decoded.ProvisioningState
	s.RequestDiskPersistence = decoded.RequestDiskPersistence

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DataflowGraphProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["nodes"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Nodes into list []json.RawMessage: %+v", err)
		}

		output := make([]DataflowGraphNode, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDataflowGraphNodeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Nodes' for 'DataflowGraphProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Nodes = output
	}

	return nil
}
