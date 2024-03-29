package restorepoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointSourceVMOSDisk struct {
	Caching                 *CachingTypes               `json:"caching,omitempty"`
	DiskRestorePoint        *DiskRestorePointAttributes `json:"diskRestorePoint,omitempty"`
	DiskSizeGB              *int64                      `json:"diskSizeGB,omitempty"`
	EncryptionSettings      *DiskEncryptionSettings     `json:"encryptionSettings,omitempty"`
	ManagedDisk             *ManagedDiskParameters      `json:"managedDisk,omitempty"`
	Name                    *string                     `json:"name,omitempty"`
	OsType                  *OperatingSystemType        `json:"osType,omitempty"`
	WriteAcceleratorEnabled *bool                       `json:"writeAcceleratorEnabled,omitempty"`
}
