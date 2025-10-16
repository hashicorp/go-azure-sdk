package dataflow

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowDestinationOperationSettings struct {
	DataDestination string                             `json:"dataDestination"`
	EndpointRef     string                             `json:"endpointRef"`
	Headers         *[]DataflowDestinationHeaderAction `json:"headers,omitempty"`
}

var _ json.Unmarshaler = &DataflowDestinationOperationSettings{}

func (s *DataflowDestinationOperationSettings) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling DataflowDestinationOperationSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["headers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Headers into list []json.RawMessage: %+v", err)
		}

		output := make([]DataflowDestinationHeaderAction, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDataflowDestinationHeaderActionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Headers' for 'DataflowDestinationOperationSettings': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Headers = &output
	}

	return nil
}
