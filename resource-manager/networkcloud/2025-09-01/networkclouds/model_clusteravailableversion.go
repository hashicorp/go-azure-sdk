package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterAvailableVersion struct {
	SupportExpiryDate    *string `json:"supportExpiryDate,omitempty"`
	TargetClusterVersion *string `json:"targetClusterVersion,omitempty"`
}
