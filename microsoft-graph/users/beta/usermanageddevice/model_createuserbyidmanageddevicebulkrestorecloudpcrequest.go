package usermanageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdManagedDeviceBulkRestoreCloudPCRequest struct {
	ManagedDeviceIds     *[]string                                                      `json:"managedDeviceIds,omitempty"`
	RestorePointDateTime *string                                                        `json:"restorePointDateTime,omitempty"`
	TimeRange            *CreateUserByIdManagedDeviceBulkRestoreCloudPCRequestTimeRange `json:"timeRange,omitempty"`
}
