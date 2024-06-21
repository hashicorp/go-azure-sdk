package backupprotecteditems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureVMWorkloadProtectedItemExtendedInfo struct {
	NewestRecoveryPointInArchive *string `json:"newestRecoveryPointInArchive,omitempty"`
	OldestRecoveryPoint          *string `json:"oldestRecoveryPoint,omitempty"`
	OldestRecoveryPointInArchive *string `json:"oldestRecoveryPointInArchive,omitempty"`
	OldestRecoveryPointInVault   *string `json:"oldestRecoveryPointInVault,omitempty"`
	PolicyState                  *string `json:"policyState,omitempty"`
	RecoveryModel                *string `json:"recoveryModel,omitempty"`
	RecoveryPointCount           *int64  `json:"recoveryPointCount,omitempty"`
}
