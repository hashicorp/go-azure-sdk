package placementpolicies

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlacementPolicyProperties = VMVMPlacementPolicyProperties{}

type VMVMPlacementPolicyProperties struct {
	AffinityType AffinityType `json:"affinityType"`
	VMMembers    []string     `json:"vmMembers"`

	// Fields inherited from PlacementPolicyProperties

	DisplayName       *string                           `json:"displayName,omitempty"`
	ProvisioningState *PlacementPolicyProvisioningState `json:"provisioningState,omitempty"`
	State             *PlacementPolicyState             `json:"state,omitempty"`
	Type              PlacementPolicyType               `json:"type"`
}

func (s VMVMPlacementPolicyProperties) PlacementPolicyProperties() BasePlacementPolicyPropertiesImpl {
	return BasePlacementPolicyPropertiesImpl{
		DisplayName:       s.DisplayName,
		ProvisioningState: s.ProvisioningState,
		State:             s.State,
		Type:              s.Type,
	}
}

var _ json.Marshaler = VMVMPlacementPolicyProperties{}

func (s VMVMPlacementPolicyProperties) MarshalJSON() ([]byte, error) {
	type wrapper VMVMPlacementPolicyProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VMVMPlacementPolicyProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VMVMPlacementPolicyProperties: %+v", err)
	}

	decoded["type"] = "VmVm"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VMVMPlacementPolicyProperties: %+v", err)
	}

	return encoded, nil
}
