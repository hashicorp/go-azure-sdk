package memanageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeManagedDeviceBulkRestoreCloudPCRequest struct {
	ManagedDeviceIds     *[]string                                                `json:"managedDeviceIds,omitempty"`
	RestorePointDateTime *string                                                  `json:"restorePointDateTime,omitempty"`
	TimeRange            *CreateMeManagedDeviceBulkRestoreCloudPCRequestTimeRange `json:"timeRange,omitempty"`
}
