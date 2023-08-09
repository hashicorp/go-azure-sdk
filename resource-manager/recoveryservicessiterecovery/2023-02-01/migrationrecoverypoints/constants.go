package migrationrecoverypoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *MigrationRecoveryPointType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrationRecoveryPointType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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
