package schedule

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitorComputeIdentityBase = ManagedComputeIdentity{}

type ManagedComputeIdentity struct {
	Identity *identity.LegacySystemAndUserAssignedMap `json:"identity,omitempty"`

	// Fields inherited from MonitorComputeIdentityBase

	ComputeIdentityType MonitorComputeIdentityType `json:"computeIdentityType"`
}

func (s ManagedComputeIdentity) MonitorComputeIdentityBase() BaseMonitorComputeIdentityBaseImpl {
	return BaseMonitorComputeIdentityBaseImpl{
		ComputeIdentityType: s.ComputeIdentityType,
	}
}

var _ json.Marshaler = ManagedComputeIdentity{}

func (s ManagedComputeIdentity) MarshalJSON() ([]byte, error) {
	type wrapper ManagedComputeIdentity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedComputeIdentity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedComputeIdentity: %+v", err)
	}

	decoded["computeIdentityType"] = "ManagedIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedComputeIdentity: %+v", err)
	}

	return encoded, nil
}
