package storageaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GeoReplicationStats struct {
	CanFailover                   *bool                          `json:"canFailover,omitempty"`
	CanPlannedFailover            *bool                          `json:"canPlannedFailover,omitempty"`
	LastSyncTime                  *string                        `json:"lastSyncTime,omitempty"`
	PostFailoverRedundancy        *PostFailoverRedundancy        `json:"postFailoverRedundancy,omitempty"`
	PostPlannedFailoverRedundancy *PostPlannedFailoverRedundancy `json:"postPlannedFailoverRedundancy,omitempty"`
	Status                        *GeoReplicationStatus          `json:"status,omitempty"`
}
