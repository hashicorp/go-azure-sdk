package virtualmachinescalesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetManagedDiskParameters struct {
	DiskEncryptionSet  *SubResource           `json:"diskEncryptionSet"`
	SecurityProfile    *VMDiskSecurityProfile `json:"securityProfile"`
	StorageAccountType *StorageAccountTypes   `json:"storageAccountType,omitempty"`
}
