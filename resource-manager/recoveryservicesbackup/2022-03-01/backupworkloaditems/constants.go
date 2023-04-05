package backupworkloaditems

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
