package dataflowgraph

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowGraphDestinationNodeSettings struct {
	DataDestination string                                  `json:"dataDestination"`
	EndpointRef     string                                  `json:"endpointRef"`
	Headers         *[]DataflowGraphDestinationHeaderAction `json:"headers,omitempty"`
}

var _ json.Unmarshaler = &DataflowGraphDestinationNodeSettings{}

func (s *DataflowGraphDestinationNodeSettings) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DataDestination string `json:"dataDestination"`
		EndpointRef     string `json:"endpointRef"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DataDestination = decoded.DataDestination
	s.EndpointRef = decoded.EndpointRef

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DataflowGraphDestinationNodeSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["headers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Headers into list []json.RawMessage: %+v", err)
		}

		output := make([]DataflowGraphDestinationHeaderAction, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDataflowGraphDestinationHeaderActionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Headers' for 'DataflowGraphDestinationNodeSettings': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Headers = &output
	}

	return nil
}
