package placementpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlacementPolicyUpdateProperties struct {
	HostMembers *[]string             `json:"hostMembers,omitempty"`
	State       *PlacementPolicyState `json:"state,omitempty"`
	VMMembers   *[]string             `json:"vmMembers,omitempty"`
}
