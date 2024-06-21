package serverendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudTieringCachePerformance struct {
	CacheHitBytes        *int64  `json:"cacheHitBytes,omitempty"`
	CacheHitBytesPercent *int64  `json:"cacheHitBytesPercent,omitempty"`
	CacheMissBytes       *int64  `json:"cacheMissBytes,omitempty"`
	LastUpdatedTimestamp *string `json:"lastUpdatedTimestamp,omitempty"`
}
