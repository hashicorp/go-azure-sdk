package recoverypoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointTierStatus string

const (
	RecoveryPointTierStatusDeleted    RecoveryPointTierStatus = "Deleted"
	RecoveryPointTierStatusDisabled   RecoveryPointTierStatus = "Disabled"
	RecoveryPointTierStatusInvalid    RecoveryPointTierStatus = "Invalid"
	RecoveryPointTierStatusRehydrated RecoveryPointTierStatus = "Rehydrated"
	RecoveryPointTierStatusValid      RecoveryPointTierStatus = "Valid"
)

func PossibleValuesForRecoveryPointTierStatus() []string {
	return []string{
		string(RecoveryPointTierStatusDeleted),
		string(RecoveryPointTierStatusDisabled),
		string(RecoveryPointTierStatusInvalid),
		string(RecoveryPointTierStatusRehydrated),
		string(RecoveryPointTierStatusValid),
	}
}

type RecoveryPointTierType string

const (
	RecoveryPointTierTypeArchivedRP RecoveryPointTierType = "ArchivedRP"
	RecoveryPointTierTypeHardenedRP RecoveryPointTierType = "HardenedRP"
	RecoveryPointTierTypeInstantRP  RecoveryPointTierType = "InstantRP"
	RecoveryPointTierTypeInvalid    RecoveryPointTierType = "Invalid"
)

func PossibleValuesForRecoveryPointTierType() []string {
	return []string{
		string(RecoveryPointTierTypeArchivedRP),
		string(RecoveryPointTierTypeHardenedRP),
		string(RecoveryPointTierTypeInstantRP),
		string(RecoveryPointTierTypeInvalid),
	}
}

type RestorePointType string

const (
	RestorePointTypeDifferential RestorePointType = "Differential"
	RestorePointTypeFull         RestorePointType = "Full"
	RestorePointTypeIncremental  RestorePointType = "Incremental"
	RestorePointTypeInvalid      RestorePointType = "Invalid"
	RestorePointTypeLog          RestorePointType = "Log"
)

func PossibleValuesForRestorePointType() []string {
	return []string{
		string(RestorePointTypeDifferential),
		string(RestorePointTypeFull),
		string(RestorePointTypeIncremental),
		string(RestorePointTypeInvalid),
		string(RestorePointTypeLog),
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
