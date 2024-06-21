package migrationrecoverypoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationRecoveryPointProperties struct {
	RecoveryPointTime *string                     `json:"recoveryPointTime,omitempty"`
	RecoveryPointType *MigrationRecoveryPointType `json:"recoveryPointType,omitempty"`
}
