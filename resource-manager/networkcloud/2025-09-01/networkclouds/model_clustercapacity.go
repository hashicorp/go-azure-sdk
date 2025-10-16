package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterCapacity struct {
	AvailableApplianceStorageGB *int64 `json:"availableApplianceStorageGB,omitempty"`
	AvailableCoreCount          *int64 `json:"availableCoreCount,omitempty"`
	AvailableHostStorageGB      *int64 `json:"availableHostStorageGB,omitempty"`
	AvailableMemoryGB           *int64 `json:"availableMemoryGB,omitempty"`
	TotalApplianceStorageGB     *int64 `json:"totalApplianceStorageGB,omitempty"`
	TotalCoreCount              *int64 `json:"totalCoreCount,omitempty"`
	TotalHostStorageGB          *int64 `json:"totalHostStorageGB,omitempty"`
	TotalMemoryGB               *int64 `json:"totalMemoryGB,omitempty"`
}
