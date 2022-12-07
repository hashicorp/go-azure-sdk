package backupworkloaditems

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionStatus string

const (
	ProtectionStatusInvalid          ProtectionStatus = "Invalid"
	ProtectionStatusNotProtected     ProtectionStatus = "NotProtected"
	ProtectionStatusProtected        ProtectionStatus = "Protected"
	ProtectionStatusProtecting       ProtectionStatus = "Protecting"
	ProtectionStatusProtectionFailed ProtectionStatus = "ProtectionFailed"
)

func PossibleValuesForProtectionStatus() []string {
	return []string{
		string(ProtectionStatusInvalid),
		string(ProtectionStatusNotProtected),
		string(ProtectionStatusProtected),
		string(ProtectionStatusProtecting),
		string(ProtectionStatusProtectionFailed),
	}
}

func parseProtectionStatus(input string) (*ProtectionStatus, error) {
	vals := map[string]ProtectionStatus{
		"invalid":          ProtectionStatusInvalid,
		"notprotected":     ProtectionStatusNotProtected,
		"protected":        ProtectionStatusProtected,
		"protecting":       ProtectionStatusProtecting,
		"protectionfailed": ProtectionStatusProtectionFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionStatus(input)
	return &out, nil
}

type SQLDataDirectoryType string

const (
	SQLDataDirectoryTypeData    SQLDataDirectoryType = "Data"
	SQLDataDirectoryTypeInvalid SQLDataDirectoryType = "Invalid"
	SQLDataDirectoryTypeLog     SQLDataDirectoryType = "Log"
)

func PossibleValuesForSQLDataDirectoryType() []string {
	return []string{
		string(SQLDataDirectoryTypeData),
		string(SQLDataDirectoryTypeInvalid),
		string(SQLDataDirectoryTypeLog),
	}
}

func parseSQLDataDirectoryType(input string) (*SQLDataDirectoryType, error) {
	vals := map[string]SQLDataDirectoryType{
		"data":    SQLDataDirectoryTypeData,
		"invalid": SQLDataDirectoryTypeInvalid,
		"log":     SQLDataDirectoryTypeLog,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SQLDataDirectoryType(input)
	return &out, nil
}
