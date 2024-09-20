package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryVirtualNetworkCustomDetails = NewRecoveryVirtualNetwork{}

type NewRecoveryVirtualNetwork struct {
	RecoveryVirtualNetworkName              *string `json:"recoveryVirtualNetworkName,omitempty"`
	RecoveryVirtualNetworkResourceGroupName *string `json:"recoveryVirtualNetworkResourceGroupName,omitempty"`

	// Fields inherited from RecoveryVirtualNetworkCustomDetails

	ResourceType string `json:"resourceType"`
}

func (s NewRecoveryVirtualNetwork) RecoveryVirtualNetworkCustomDetails() BaseRecoveryVirtualNetworkCustomDetailsImpl {
	return BaseRecoveryVirtualNetworkCustomDetailsImpl{
		ResourceType: s.ResourceType,
	}
}

var _ json.Marshaler = NewRecoveryVirtualNetwork{}

func (s NewRecoveryVirtualNetwork) MarshalJSON() ([]byte, error) {
	type wrapper NewRecoveryVirtualNetwork
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NewRecoveryVirtualNetwork: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NewRecoveryVirtualNetwork: %+v", err)
	}

	decoded["resourceType"] = "New"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NewRecoveryVirtualNetwork: %+v", err)
	}

	return encoded, nil
}
