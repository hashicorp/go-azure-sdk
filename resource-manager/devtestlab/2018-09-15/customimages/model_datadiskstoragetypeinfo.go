package customimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataDiskStorageTypeInfo struct {
	Lun         *string      `json:"lun,omitempty"`
	StorageType *StorageType `json:"storageType,omitempty"`
}
