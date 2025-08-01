package computerps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResiliencyPolicy struct {
	AutomaticZoneRebalancingPolicy *AutomaticZoneRebalancingPolicy `json:"automaticZoneRebalancingPolicy,omitempty"`
	ResilientVMCreationPolicy      *ResilientVMCreationPolicy      `json:"resilientVMCreationPolicy,omitempty"`
	ResilientVMDeletionPolicy      *ResilientVMDeletionPolicy      `json:"resilientVMDeletionPolicy,omitempty"`
}
