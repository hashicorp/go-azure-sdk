package serverendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerEndpointCloudTieringStatus struct {
	CachePerformance            *CloudTieringCachePerformance            `json:"cachePerformance,omitempty"`
	DatePolicyStatus            *CloudTieringDatePolicyStatus            `json:"datePolicyStatus,omitempty"`
	FilesNotTiering             *CloudTieringFilesNotTiering             `json:"filesNotTiering,omitempty"`
	Health                      *ServerEndpointHealthState               `json:"health,omitempty"`
	HealthLastUpdatedTimestamp  *string                                  `json:"healthLastUpdatedTimestamp,omitempty"`
	LastCloudTieringResult      *int64                                   `json:"lastCloudTieringResult,omitempty"`
	LastSuccessTimestamp        *string                                  `json:"lastSuccessTimestamp,omitempty"`
	LastUpdatedTimestamp        *string                                  `json:"lastUpdatedTimestamp,omitempty"`
	LowDiskMode                 *CloudTieringLowDiskMode                 `json:"lowDiskMode,omitempty"`
	SpaceSavings                *CloudTieringSpaceSavings                `json:"spaceSavings,omitempty"`
	VolumeFreeSpacePolicyStatus *CloudTieringVolumeFreeSpacePolicyStatus `json:"volumeFreeSpacePolicyStatus,omitempty"`
}
