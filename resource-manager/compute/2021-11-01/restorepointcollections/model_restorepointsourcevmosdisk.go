package restorepointcollections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointSourceVMOSDisk struct {
	Caching            *CachingTypes           `json:"caching,omitempty"`
	DiskRestorePoint   *ApiEntityReference     `json:"diskRestorePoint"`
	DiskSizeGB         *int64                  `json:"diskSizeGB,omitempty"`
	EncryptionSettings *DiskEncryptionSettings `json:"encryptionSettings"`
	ManagedDisk        *ManagedDiskParameters  `json:"managedDisk"`
	Name               *string                 `json:"name,omitempty"`
	OsType             *OperatingSystemType    `json:"osType,omitempty"`
}
