package backupresourcestorageconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupResourceConfig struct {
	CrossRegionRestoreFlag *bool             `json:"crossRegionRestoreFlag,omitempty"`
	StorageModelType       *StorageType      `json:"storageModelType,omitempty"`
	StorageType            *StorageType      `json:"storageType,omitempty"`
	StorageTypeState       *StorageTypeState `json:"storageTypeState,omitempty"`
}
