package serverendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudTieringSpaceSavings struct {
	CachedSizeBytes      *int64  `json:"cachedSizeBytes,omitempty"`
	LastUpdatedTimestamp *string `json:"lastUpdatedTimestamp,omitempty"`
	SpaceSavingsBytes    *int64  `json:"spaceSavingsBytes,omitempty"`
	SpaceSavingsPercent  *int64  `json:"spaceSavingsPercent,omitempty"`
	TotalSizeCloudBytes  *int64  `json:"totalSizeCloudBytes,omitempty"`
	VolumeSizeBytes      *int64  `json:"volumeSizeBytes,omitempty"`
}
