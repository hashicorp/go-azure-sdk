package pool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDisk struct {
	SecurityProfile    *VMDiskSecurityProfile `json:"securityProfile,omitempty"`
	StorageAccountType *StorageAccountType    `json:"storageAccountType,omitempty"`
}
