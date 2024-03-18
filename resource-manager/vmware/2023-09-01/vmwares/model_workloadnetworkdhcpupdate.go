package vmwares

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkDhcpUpdate struct {
	Properties WorkloadNetworkDhcpEntityUpdate `json:"properties"`
}

var _ json.Unmarshaler = &WorkloadNetworkDhcpUpdate{}

func (s *WorkloadNetworkDhcpUpdate) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WorkloadNetworkDhcpUpdate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := unmarshalWorkloadNetworkDhcpEntityUpdateImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'WorkloadNetworkDhcpUpdate': %+v", err)
		}
		s.Properties = impl
	}
	return nil
}
