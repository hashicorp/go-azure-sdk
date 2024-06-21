package disks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskProperties struct {
	CreatedDate       *string      `json:"createdDate,omitempty"`
	DiskBlobName      *string      `json:"diskBlobName,omitempty"`
	DiskSizeGiB       *int64       `json:"diskSizeGiB,omitempty"`
	DiskType          *StorageType `json:"diskType,omitempty"`
	DiskUri           *string      `json:"diskUri,omitempty"`
	HostCaching       *string      `json:"hostCaching,omitempty"`
	LeasedByLabVMId   *string      `json:"leasedByLabVmId,omitempty"`
	ManagedDiskId     *string      `json:"managedDiskId,omitempty"`
	ProvisioningState *string      `json:"provisioningState,omitempty"`
	StorageAccountId  *string      `json:"storageAccountId,omitempty"`
	UniqueIdentifier  *string      `json:"uniqueIdentifier,omitempty"`
}
