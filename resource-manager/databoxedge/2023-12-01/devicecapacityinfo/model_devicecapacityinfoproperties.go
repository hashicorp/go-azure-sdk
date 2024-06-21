package devicecapacityinfo

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCapacityInfoProperties struct {
	ClusterComputeCapacityInfo *ClusterCapacityViewData `json:"clusterComputeCapacityInfo,omitempty"`
	ClusterStorageCapacityInfo *ClusterStorageViewData  `json:"clusterStorageCapacityInfo,omitempty"`
	NodeCapacityInfos          *map[string]HostCapacity `json:"nodeCapacityInfos,omitempty"`
	TimeStamp                  *string                  `json:"timeStamp,omitempty"`
}
