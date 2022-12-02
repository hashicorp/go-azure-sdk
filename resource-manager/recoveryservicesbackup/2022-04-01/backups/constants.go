package backups

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupType string

const (
	BackupTypeCopyOnlyFull BackupType = "CopyOnlyFull"
	BackupTypeDifferential BackupType = "Differential"
	BackupTypeFull         BackupType = "Full"
	BackupTypeIncremental  BackupType = "Incremental"
	BackupTypeInvalid      BackupType = "Invalid"
	BackupTypeLog          BackupType = "Log"
)

func PossibleValuesForBackupType() []string {
	return []string{
		string(BackupTypeCopyOnlyFull),
		string(BackupTypeDifferential),
		string(BackupTypeFull),
		string(BackupTypeIncremental),
		string(BackupTypeInvalid),
		string(BackupTypeLog),
	}
}

func parseBackupType(input string) (*BackupType, error) {
	vals := map[string]BackupType{
		"copyonlyfull": BackupTypeCopyOnlyFull,
		"differential": BackupTypeDifferential,
		"full":         BackupTypeFull,
		"incremental":  BackupTypeIncremental,
		"invalid":      BackupTypeInvalid,
		"log":          BackupTypeLog,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupType(input)
	return &out, nil
}
