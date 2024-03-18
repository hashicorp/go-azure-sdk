package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryVirtualNetworkCustomDetails = ExistingRecoveryVirtualNetwork{}

type ExistingRecoveryVirtualNetwork struct {
	RecoverySubnetName       *string `json:"recoverySubnetName,omitempty"`
	RecoveryVirtualNetworkId string  `json:"recoveryVirtualNetworkId"`

	// Fields inherited from RecoveryVirtualNetworkCustomDetails
}

var _ json.Marshaler = ExistingRecoveryVirtualNetwork{}

func (s ExistingRecoveryVirtualNetwork) MarshalJSON() ([]byte, error) {
	type wrapper ExistingRecoveryVirtualNetwork
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExistingRecoveryVirtualNetwork: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExistingRecoveryVirtualNetwork: %+v", err)
	}
	decoded["resourceType"] = "Existing"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExistingRecoveryVirtualNetwork: %+v", err)
	}

	return encoded, nil
}
