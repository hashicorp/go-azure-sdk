package replicationmigrationitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMwareCbtProtectedDiskDetails struct {
	CapacityInBytes                *int64           `json:"capacityInBytes,omitempty"`
	DiskEncryptionSetId            *string          `json:"diskEncryptionSetId,omitempty"`
	DiskId                         *string          `json:"diskId,omitempty"`
	DiskName                       *string          `json:"diskName,omitempty"`
	DiskPath                       *string          `json:"diskPath,omitempty"`
	DiskType                       *DiskAccountType `json:"diskType,omitempty"`
	IsOSDisk                       *string          `json:"isOSDisk,omitempty"`
	LogStorageAccountId            *string          `json:"logStorageAccountId,omitempty"`
	LogStorageAccountSasSecretName *string          `json:"logStorageAccountSasSecretName,omitempty"`
	SeedBlobUri                    *string          `json:"seedBlobUri,omitempty"`
	SeedManagedDiskId              *string          `json:"seedManagedDiskId,omitempty"`
	TargetBlobUri                  *string          `json:"targetBlobUri,omitempty"`
	TargetDiskName                 *string          `json:"targetDiskName,omitempty"`
	TargetManagedDiskId            *string          `json:"targetManagedDiskId,omitempty"`
}
