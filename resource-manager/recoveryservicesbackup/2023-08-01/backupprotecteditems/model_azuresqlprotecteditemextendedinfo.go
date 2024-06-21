package backupprotecteditems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureSqlProtectedItemExtendedInfo struct {
	OldestRecoveryPoint *string `json:"oldestRecoveryPoint,omitempty"`
	PolicyState         *string `json:"policyState,omitempty"`
	RecoveryPointCount  *int64  `json:"recoveryPointCount,omitempty"`
}
