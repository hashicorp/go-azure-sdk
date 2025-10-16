package recoverypointsrecommendedformove

import (
	"strings"
)

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

func parseRecoveryPointTierStatus(input string) (*RecoveryPointTierStatus, error) {
	vals := map[string]RecoveryPointTierStatus{
		"deleted":    RecoveryPointTierStatusDeleted,
		"disabled":   RecoveryPointTierStatusDisabled,
		"invalid":    RecoveryPointTierStatusInvalid,
		"rehydrated": RecoveryPointTierStatusRehydrated,
		"valid":      RecoveryPointTierStatusValid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecoveryPointTierStatus(input)
	return &out, nil
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

func parseRecoveryPointTierType(input string) (*RecoveryPointTierType, error) {
	vals := map[string]RecoveryPointTierType{
		"archivedrp": RecoveryPointTierTypeArchivedRP,
		"hardenedrp": RecoveryPointTierTypeHardenedRP,
		"instantrp":  RecoveryPointTierTypeInstantRP,
		"invalid":    RecoveryPointTierTypeInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecoveryPointTierType(input)
	return &out, nil
}

type RestorePointType string

const (
	RestorePointTypeDifferential         RestorePointType = "Differential"
	RestorePointTypeFull                 RestorePointType = "Full"
	RestorePointTypeIncremental          RestorePointType = "Incremental"
	RestorePointTypeInvalid              RestorePointType = "Invalid"
	RestorePointTypeLog                  RestorePointType = "Log"
	RestorePointTypeSnapshotCopyOnlyFull RestorePointType = "SnapshotCopyOnlyFull"
	RestorePointTypeSnapshotFull         RestorePointType = "SnapshotFull"
)

func PossibleValuesForRestorePointType() []string {
	return []string{
		string(RestorePointTypeDifferential),
		string(RestorePointTypeFull),
		string(RestorePointTypeIncremental),
		string(RestorePointTypeInvalid),
		string(RestorePointTypeLog),
		string(RestorePointTypeSnapshotCopyOnlyFull),
		string(RestorePointTypeSnapshotFull),
	}
}

func parseRestorePointType(input string) (*RestorePointType, error) {
	vals := map[string]RestorePointType{
		"differential":         RestorePointTypeDifferential,
		"full":                 RestorePointTypeFull,
		"incremental":          RestorePointTypeIncremental,
		"invalid":              RestorePointTypeInvalid,
		"log":                  RestorePointTypeLog,
		"snapshotcopyonlyfull": RestorePointTypeSnapshotCopyOnlyFull,
		"snapshotfull":         RestorePointTypeSnapshotFull,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestorePointType(input)
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
