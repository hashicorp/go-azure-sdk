package regioninforesources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegionInfo struct {
	AvailabilityZoneMappings  *[]RegionInfoAvailabilityZoneMappingsItem `json:"availabilityZoneMappings,omitempty"`
	StorageToNetworkProximity *RegionStorageToNetworkProximity          `json:"storageToNetworkProximity,omitempty"`
}
