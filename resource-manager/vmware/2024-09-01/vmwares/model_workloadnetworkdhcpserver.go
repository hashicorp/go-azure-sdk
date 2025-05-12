package vmwares

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkloadNetworkDhcpEntity = WorkloadNetworkDhcpServer{}

type WorkloadNetworkDhcpServer struct {
	LeaseTime     *int64  `json:"leaseTime,omitempty"`
	ServerAddress *string `json:"serverAddress,omitempty"`

	// Fields inherited from WorkloadNetworkDhcpEntity

	DhcpType          DhcpTypeEnum                          `json:"dhcpType"`
	DisplayName       *string                               `json:"displayName,omitempty"`
	ProvisioningState *WorkloadNetworkDhcpProvisioningState `json:"provisioningState,omitempty"`
	Revision          *int64                                `json:"revision,omitempty"`
	Segments          *[]string                             `json:"segments,omitempty"`
}

func (s WorkloadNetworkDhcpServer) WorkloadNetworkDhcpEntity() BaseWorkloadNetworkDhcpEntityImpl {
	return BaseWorkloadNetworkDhcpEntityImpl{
		DhcpType:          s.DhcpType,
		DisplayName:       s.DisplayName,
		ProvisioningState: s.ProvisioningState,
		Revision:          s.Revision,
		Segments:          s.Segments,
	}
}

var _ json.Marshaler = WorkloadNetworkDhcpServer{}

func (s WorkloadNetworkDhcpServer) MarshalJSON() ([]byte, error) {
	type wrapper WorkloadNetworkDhcpServer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkloadNetworkDhcpServer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkloadNetworkDhcpServer: %+v", err)
	}

	decoded["dhcpType"] = "SERVER"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkloadNetworkDhcpServer: %+v", err)
	}

	return encoded, nil
}
