package replicationprotecteditems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InMageProtectedDiskDetails struct {
	DiskCapacityInBytes                 *int64   `json:"diskCapacityInBytes,omitempty"`
	DiskId                              *string  `json:"diskId,omitempty"`
	DiskName                            *string  `json:"diskName,omitempty"`
	DiskResized                         *string  `json:"diskResized,omitempty"`
	FileSystemCapacityInBytes           *int64   `json:"fileSystemCapacityInBytes,omitempty"`
	HealthErrorCode                     *string  `json:"healthErrorCode,omitempty"`
	LastRpoCalculatedTime               *string  `json:"lastRpoCalculatedTime,omitempty"`
	ProgressHealth                      *string  `json:"progressHealth,omitempty"`
	ProgressStatus                      *string  `json:"progressStatus,omitempty"`
	ProtectionStage                     *string  `json:"protectionStage,omitempty"`
	PsDataInMB                          *float64 `json:"psDataInMB,omitempty"`
	ResyncDurationInSeconds             *int64   `json:"resyncDurationInSeconds,omitempty"`
	ResyncLast15MinutesTransferredBytes *int64   `json:"resyncLast15MinutesTransferredBytes,omitempty"`
	ResyncLastDataTransferTimeUTC       *string  `json:"resyncLastDataTransferTimeUTC,omitempty"`
	ResyncProcessedBytes                *int64   `json:"resyncProcessedBytes,omitempty"`
	ResyncProgressPercentage            *int64   `json:"resyncProgressPercentage,omitempty"`
	ResyncRequired                      *string  `json:"resyncRequired,omitempty"`
	ResyncStartTime                     *string  `json:"resyncStartTime,omitempty"`
	ResyncTotalTransferredBytes         *int64   `json:"resyncTotalTransferredBytes,omitempty"`
	RpoInSeconds                        *int64   `json:"rpoInSeconds,omitempty"`
	SourceDataInMB                      *float64 `json:"sourceDataInMB,omitempty"`
	TargetDataInMB                      *float64 `json:"targetDataInMB,omitempty"`
}
