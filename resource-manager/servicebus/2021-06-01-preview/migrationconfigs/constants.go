package migrationconfigs

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationConfigurationName string

const (
	MigrationConfigurationNameDefault MigrationConfigurationName = "$default"
)

func PossibleValuesForMigrationConfigurationName() []string {
	return []string{
		string(MigrationConfigurationNameDefault),
	}
}

func parseMigrationConfigurationName(input string) (*MigrationConfigurationName, error) {
	vals := map[string]MigrationConfigurationName{
		"$default": MigrationConfigurationNameDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationConfigurationName(input)
	return &out, nil
}
