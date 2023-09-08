package usermanageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdManagedDeviceBulkReprovisionCloudPCRequest struct {
	ManagedDeviceIds *[]string `json:"managedDeviceIds,omitempty"`
}
