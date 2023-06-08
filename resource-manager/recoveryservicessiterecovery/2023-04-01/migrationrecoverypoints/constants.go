package migrationrecoverypoints

import "strings"

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

func parseMigrationRecoveryPointType(input string) (*MigrationRecoveryPointType, error) {
	vals := map[string]MigrationRecoveryPointType{
		"applicationconsistent": MigrationRecoveryPointTypeApplicationConsistent,
		"crashconsistent":       MigrationRecoveryPointTypeCrashConsistent,
		"notspecified":          MigrationRecoveryPointTypeNotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationRecoveryPointType(input)
	return &out, nil
}
