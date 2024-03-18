package vmwares

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkloadNetworkDhcpEntityUpdate = WorkloadNetworkDhcpServerUpdate{}

type WorkloadNetworkDhcpServerUpdate struct {
	LeaseTime     *int64  `json:"leaseTime,omitempty"`
	ServerAddress *string `json:"serverAddress,omitempty"`

	// Fields inherited from WorkloadNetworkDhcpEntityUpdate
	DisplayName *string `json:"displayName,omitempty"`
	Revision    *int64  `json:"revision,omitempty"`
}

var _ json.Marshaler = WorkloadNetworkDhcpServerUpdate{}

func (s WorkloadNetworkDhcpServerUpdate) MarshalJSON() ([]byte, error) {
	type wrapper WorkloadNetworkDhcpServerUpdate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkloadNetworkDhcpServerUpdate: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkloadNetworkDhcpServerUpdate: %+v", err)
	}
	decoded["dhcpType"] = "SERVER"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkloadNetworkDhcpServerUpdate: %+v", err)
	}

	return encoded, nil
}
