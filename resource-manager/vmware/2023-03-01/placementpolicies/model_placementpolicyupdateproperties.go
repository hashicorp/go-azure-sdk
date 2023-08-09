package placementpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlacementPolicyUpdateProperties struct {
	AffinityStrength       *AffinityStrength       `json:"affinityStrength,omitempty"`
	AzureHybridBenefitType *AzureHybridBenefitType `json:"azureHybridBenefitType,omitempty"`
	HostMembers            *[]string               `json:"hostMembers,omitempty"`
	State                  *PlacementPolicyState   `json:"state,omitempty"`
	VMMembers              *[]string               `json:"vmMembers,omitempty"`
}
