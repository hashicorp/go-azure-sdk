package vmwares

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkloadNetworkDhcpEntityUpdate = WorkloadNetworkDhcpRelayUpdate{}

type WorkloadNetworkDhcpRelayUpdate struct {
	ServerAddresses *[]string `json:"serverAddresses,omitempty"`

	// Fields inherited from WorkloadNetworkDhcpEntityUpdate
	DisplayName *string `json:"displayName,omitempty"`
	Revision    *int64  `json:"revision,omitempty"`
}

var _ json.Marshaler = WorkloadNetworkDhcpRelayUpdate{}

func (s WorkloadNetworkDhcpRelayUpdate) MarshalJSON() ([]byte, error) {
	type wrapper WorkloadNetworkDhcpRelayUpdate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkloadNetworkDhcpRelayUpdate: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkloadNetworkDhcpRelayUpdate: %+v", err)
	}
	decoded["dhcpType"] = "RELAY"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkloadNetworkDhcpRelayUpdate: %+v", err)
	}

	return encoded, nil
}
