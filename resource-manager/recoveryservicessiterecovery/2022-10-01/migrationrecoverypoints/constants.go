package migrationrecoverypoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationRecoveryPointType string

const (
	MigrationRecoveryPointTypeApplicationConsistent MigrationRecoveryPointType = "ApplicationConsistent"
	MigrationRecoveryPointTypeCrashConsistent       MigrationRecoveryPointType = "CrashConsistent"
	MigrationRecoveryPointTypeNotSpecified          MigrationRecoveryPointType = "NotSpecified"
)

func PossibleValuesForMigrationRecoveryPointType() []string {
	return []string{
		string(MigrationRecoveryPointTypeApplicationConsistent),
		string(MigrationRecoveryPointTypeCrashConsistent),
		string(MigrationRecoveryPointTypeNotSpecified),
	}
}
