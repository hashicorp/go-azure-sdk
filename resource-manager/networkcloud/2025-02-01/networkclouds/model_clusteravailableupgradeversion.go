package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterAvailableUpgradeVersion struct {
	ControlImpact        *ControlImpact  `json:"controlImpact,omitempty"`
	ExpectedDuration     *string         `json:"expectedDuration,omitempty"`
	ImpactDescription    *string         `json:"impactDescription,omitempty"`
	SupportExpiryDate    *string         `json:"supportExpiryDate,omitempty"`
	TargetClusterVersion *string         `json:"targetClusterVersion,omitempty"`
	WorkloadImpact       *WorkloadImpact `json:"workloadImpact,omitempty"`
}
