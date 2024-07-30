package comanageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateComanagedDeviceBulkRestoreCloudPCRequest struct {
	ManagedDeviceIds     *[]string                                                `json:"managedDeviceIds,omitempty"`
	RestorePointDateTime *string                                                  `json:"restorePointDateTime,omitempty"`
	TimeRange            *CreateComanagedDeviceBulkRestoreCloudPCRequestTimeRange `json:"timeRange,omitempty"`
}
