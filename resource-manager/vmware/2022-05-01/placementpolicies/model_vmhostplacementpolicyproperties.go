package placementpolicies

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlacementPolicyProperties = VmHostPlacementPolicyProperties{}

type VmHostPlacementPolicyProperties struct {
	AffinityStrength       *AffinityStrength       `json:"affinityStrength,omitempty"`
	AffinityType           AffinityType            `json:"affinityType"`
	AzureHybridBenefitType *AzureHybridBenefitType `json:"azureHybridBenefitType,omitempty"`
	HostMembers            []string                `json:"hostMembers"`
	VmMembers              []string                `json:"vmMembers"`

	// Fields inherited from PlacementPolicyProperties
	DisplayName       *string                           `json:"displayName,omitempty"`
	ProvisioningState *PlacementPolicyProvisioningState `json:"provisioningState,omitempty"`
	State             *PlacementPolicyState             `json:"state,omitempty"`
}

var _ json.Marshaler = VmHostPlacementPolicyProperties{}

func (s VmHostPlacementPolicyProperties) MarshalJSON() ([]byte, error) {
	type wrapper VmHostPlacementPolicyProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VmHostPlacementPolicyProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VmHostPlacementPolicyProperties: %+v", err)
	}
	decoded["type"] = "VmHost"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VmHostPlacementPolicyProperties: %+v", err)
	}

	return encoded, nil
}
