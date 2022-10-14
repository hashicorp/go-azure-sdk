package placementpolicies

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlacementPolicyProperties = VmVmPlacementPolicyProperties{}

type VmVmPlacementPolicyProperties struct {
	AffinityType AffinityType `json:"affinityType"`
	VmMembers    []string     `json:"vmMembers"`

	// Fields inherited from PlacementPolicyProperties
	DisplayName       *string                           `json:"displayName,omitempty"`
	ProvisioningState *PlacementPolicyProvisioningState `json:"provisioningState,omitempty"`
	State             *PlacementPolicyState             `json:"state,omitempty"`
}

var _ json.Marshaler = VmVmPlacementPolicyProperties{}

func (s VmVmPlacementPolicyProperties) MarshalJSON() ([]byte, error) {
	type wrapper VmVmPlacementPolicyProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VmVmPlacementPolicyProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VmVmPlacementPolicyProperties: %+v", err)
	}
	decoded["type"] = "VmVm"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VmVmPlacementPolicyProperties: %+v", err)
	}

	return encoded, nil
}
