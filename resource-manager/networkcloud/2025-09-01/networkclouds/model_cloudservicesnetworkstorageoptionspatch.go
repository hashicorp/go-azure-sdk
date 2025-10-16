package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudServicesNetworkStorageOptionsPatch struct {
	Mode               *CloudServicesNetworkStorageMode `json:"mode,omitempty"`
	SizeMiB            *int64                           `json:"sizeMiB,omitempty"`
	StorageApplianceId *string                          `json:"storageApplianceId,omitempty"`
}
