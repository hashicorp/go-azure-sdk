package vmwares

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkDhcpEntityUpdate interface {
}

// RawWorkloadNetworkDhcpEntityUpdateImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWorkloadNetworkDhcpEntityUpdateImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalWorkloadNetworkDhcpEntityUpdateImplementation(input []byte) (WorkloadNetworkDhcpEntityUpdate, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkloadNetworkDhcpEntityUpdate into map[string]interface: %+v", err)
	}

	value, ok := temp["dhcpType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "RELAY") {
		var out WorkloadNetworkDhcpRelayUpdate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkloadNetworkDhcpRelayUpdate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SERVER") {
		var out WorkloadNetworkDhcpServerUpdate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WorkloadNetworkDhcpServerUpdate: %+v", err)
		}
		return out, nil
	}

	out := RawWorkloadNetworkDhcpEntityUpdateImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
