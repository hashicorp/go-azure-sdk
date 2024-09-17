package placementpolicies

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlacementPolicyProperties = VMHostPlacementPolicyProperties{}

type VMHostPlacementPolicyProperties struct {
	AffinityStrength       *AffinityStrength       `json:"affinityStrength,omitempty"`
	AffinityType           AffinityType            `json:"affinityType"`
	AzureHybridBenefitType *AzureHybridBenefitType `json:"azureHybridBenefitType,omitempty"`
	HostMembers            []string                `json:"hostMembers"`
	VMMembers              []string                `json:"vmMembers"`

	// Fields inherited from PlacementPolicyProperties

	DisplayName       *string                           `json:"displayName,omitempty"`
	ProvisioningState *PlacementPolicyProvisioningState `json:"provisioningState,omitempty"`
	State             *PlacementPolicyState             `json:"state,omitempty"`
	Type              PlacementPolicyType               `json:"type"`
}

func (s VMHostPlacementPolicyProperties) PlacementPolicyProperties() BasePlacementPolicyPropertiesImpl {
	return BasePlacementPolicyPropertiesImpl{
		DisplayName:       s.DisplayName,
		ProvisioningState: s.ProvisioningState,
		State:             s.State,
		Type:              s.Type,
	}
}

var _ json.Marshaler = VMHostPlacementPolicyProperties{}

func (s VMHostPlacementPolicyProperties) MarshalJSON() ([]byte, error) {
	type wrapper VMHostPlacementPolicyProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VMHostPlacementPolicyProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VMHostPlacementPolicyProperties: %+v", err)
	}

	decoded["type"] = "VmHost"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VMHostPlacementPolicyProperties: %+v", err)
	}

	return encoded, nil
}
