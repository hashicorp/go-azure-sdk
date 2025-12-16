package devicecapacityinfos

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterStorageViewData struct {
	ClusterFreeStorageMb  *float64 `json:"clusterFreeStorageMb,omitempty"`
	ClusterTotalStorageMb *float64 `json:"clusterTotalStorageMb,omitempty"`
}
