package protecteditems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DPMProtectedItemExtendedInfo struct {
	DiskStorageUsedInBytes       *string            `json:"diskStorageUsedInBytes,omitempty"`
	IsCollocated                 *bool              `json:"isCollocated,omitempty"`
	IsPresentOnCloud             *bool              `json:"isPresentOnCloud,omitempty"`
	LastBackupStatus             *string            `json:"lastBackupStatus,omitempty"`
	LastRefreshedAt              *string            `json:"lastRefreshedAt,omitempty"`
	OldestRecoveryPoint          *string            `json:"oldestRecoveryPoint,omitempty"`
	OnPremiseLatestRecoveryPoint *string            `json:"onPremiseLatestRecoveryPoint,omitempty"`
	OnPremiseOldestRecoveryPoint *string            `json:"onPremiseOldestRecoveryPoint,omitempty"`
	OnPremiseRecoveryPointCount  *int64             `json:"onPremiseRecoveryPointCount,omitempty"`
	ProtectableObjectLoadPath    *map[string]string `json:"protectableObjectLoadPath,omitempty"`
	Protected                    *bool              `json:"protected,omitempty"`
	ProtectionGroupName          *string            `json:"protectionGroupName,omitempty"`
	RecoveryPointCount           *int64             `json:"recoveryPointCount,omitempty"`
	TotalDiskStorageSizeInBytes  *string            `json:"totalDiskStorageSizeInBytes,omitempty"`
}
