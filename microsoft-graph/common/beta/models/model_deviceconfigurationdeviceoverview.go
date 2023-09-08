package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationDeviceOverview struct {
	ConfigurationVersion       *int64  `json:"configurationVersion,omitempty"`
	ConflictCount              *int64  `json:"conflictCount,omitempty"`
	ErrorCount                 *int64  `json:"errorCount,omitempty"`
	FailedCount                *int64  `json:"failedCount,omitempty"`
	Id                         *string `json:"id,omitempty"`
	LastUpdateDateTime         *string `json:"lastUpdateDateTime,omitempty"`
	NotApplicableCount         *int64  `json:"notApplicableCount,omitempty"`
	NotApplicablePlatformCount *int64  `json:"notApplicablePlatformCount,omitempty"`
	ODataType                  *string `json:"@odata.type,omitempty"`
	PendingCount               *int64  `json:"pendingCount,omitempty"`
	SuccessCount               *int64  `json:"successCount,omitempty"`
}
