package vmwares

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkloadNetworkDhcpEntity = WorkloadNetworkDhcpRelay{}

type WorkloadNetworkDhcpRelay struct {
	ServerAddresses *[]string `json:"serverAddresses,omitempty"`

	// Fields inherited from WorkloadNetworkDhcpEntity

	DhcpType          DhcpTypeEnum                          `json:"dhcpType"`
	DisplayName       *string                               `json:"displayName,omitempty"`
	ProvisioningState *WorkloadNetworkDhcpProvisioningState `json:"provisioningState,omitempty"`
	Revision          *int64                                `json:"revision,omitempty"`
	Segments          *[]string                             `json:"segments,omitempty"`
}

func (s WorkloadNetworkDhcpRelay) WorkloadNetworkDhcpEntity() BaseWorkloadNetworkDhcpEntityImpl {
	return BaseWorkloadNetworkDhcpEntityImpl{
		DhcpType:          s.DhcpType,
		DisplayName:       s.DisplayName,
		ProvisioningState: s.ProvisioningState,
		Revision:          s.Revision,
		Segments:          s.Segments,
	}
}

var _ json.Marshaler = WorkloadNetworkDhcpRelay{}

func (s WorkloadNetworkDhcpRelay) MarshalJSON() ([]byte, error) {
	type wrapper WorkloadNetworkDhcpRelay
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkloadNetworkDhcpRelay: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkloadNetworkDhcpRelay: %+v", err)
	}

	decoded["dhcpType"] = "RELAY"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkloadNetworkDhcpRelay: %+v", err)
	}

	return encoded, nil
}
