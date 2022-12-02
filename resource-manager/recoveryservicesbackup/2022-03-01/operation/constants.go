package operation

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopyOptions string

const (
	CopyOptionsCreateCopy     CopyOptions = "CreateCopy"
	CopyOptionsFailOnConflict CopyOptions = "FailOnConflict"
	CopyOptionsInvalid        CopyOptions = "Invalid"
	CopyOptionsOverwrite      CopyOptions = "Overwrite"
	CopyOptionsSkip           CopyOptions = "Skip"
)

func PossibleValuesForCopyOptions() []string {
	return []string{
		string(CopyOptionsCreateCopy),
		string(CopyOptionsFailOnConflict),
		string(CopyOptionsInvalid),
		string(CopyOptionsOverwrite),
		string(CopyOptionsSkip),
	}
}

func parseCopyOptions(input string) (*CopyOptions, error) {
	vals := map[string]CopyOptions{
		"createcopy":     CopyOptionsCreateCopy,
		"failonconflict": CopyOptionsFailOnConflict,
		"invalid":        CopyOptionsInvalid,
		"overwrite":      CopyOptionsOverwrite,
		"skip":           CopyOptionsSkip,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CopyOptions(input)
	return &out, nil
}

type OverwriteOptions string

const (
	OverwriteOptionsFailOnConflict OverwriteOptions = "FailOnConflict"
	OverwriteOptionsInvalid        OverwriteOptions = "Invalid"
	OverwriteOptionsOverwrite      OverwriteOptions = "Overwrite"
)

func PossibleValuesForOverwriteOptions() []string {
	return []string{
		string(OverwriteOptionsFailOnConflict),
		string(OverwriteOptionsInvalid),
		string(OverwriteOptionsOverwrite),
	}
}

func parseOverwriteOptions(input string) (*OverwriteOptions, error) {
	vals := map[string]OverwriteOptions{
		"failonconflict": OverwriteOptionsFailOnConflict,
		"invalid":        OverwriteOptionsInvalid,
		"overwrite":      OverwriteOptionsOverwrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OverwriteOptions(input)
	return &out, nil
}

type RecoveryMode string

const (
	RecoveryModeFileRecovery     RecoveryMode = "FileRecovery"
	RecoveryModeInvalid          RecoveryMode = "Invalid"
	RecoveryModeWorkloadRecovery RecoveryMode = "WorkloadRecovery"
)

func PossibleValuesForRecoveryMode() []string {
	return []string{
		string(RecoveryModeFileRecovery),
		string(RecoveryModeInvalid),
		string(RecoveryModeWorkloadRecovery),
	}
}

func parseRecoveryMode(input string) (*RecoveryMode, error) {
	vals := map[string]RecoveryMode{
		"filerecovery":     RecoveryModeFileRecovery,
		"invalid":          RecoveryModeInvalid,
		"workloadrecovery": RecoveryModeWorkloadRecovery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecoveryMode(input)
	return &out, nil
}

type RecoveryType string

const (
	RecoveryTypeAlternateLocation RecoveryType = "AlternateLocation"
	RecoveryTypeInvalid           RecoveryType = "Invalid"
	RecoveryTypeOffline           RecoveryType = "Offline"
	RecoveryTypeOriginalLocation  RecoveryType = "OriginalLocation"
	RecoveryTypeRestoreDisks      RecoveryType = "RestoreDisks"
)

func PossibleValuesForRecoveryType() []string {
	return []string{
		string(RecoveryTypeAlternateLocation),
		string(RecoveryTypeInvalid),
		string(RecoveryTypeOffline),
		string(RecoveryTypeOriginalLocation),
		string(RecoveryTypeRestoreDisks),
	}
}

func parseRecoveryType(input string) (*RecoveryType, error) {
	vals := map[string]RecoveryType{
		"alternatelocation": RecoveryTypeAlternateLocation,
		"invalid":           RecoveryTypeInvalid,
		"offline":           RecoveryTypeOffline,
		"originallocation":  RecoveryTypeOriginalLocation,
		"restoredisks":      RecoveryTypeRestoreDisks,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecoveryType(input)
	return &out, nil
}

type RehydrationPriority string

const (
	RehydrationPriorityHigh     RehydrationPriority = "High"
	RehydrationPriorityStandard RehydrationPriority = "Standard"
)

func PossibleValuesForRehydrationPriority() []string {
	return []string{
		string(RehydrationPriorityHigh),
		string(RehydrationPriorityStandard),
	}
}

func parseRehydrationPriority(input string) (*RehydrationPriority, error) {
	vals := map[string]RehydrationPriority{
		"high":     RehydrationPriorityHigh,
		"standard": RehydrationPriorityStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RehydrationPriority(input)
	return &out, nil
}

type RestoreRequestType string

const (
	RestoreRequestTypeFullShareRestore RestoreRequestType = "FullShareRestore"
	RestoreRequestTypeInvalid          RestoreRequestType = "Invalid"
	RestoreRequestTypeItemLevelRestore RestoreRequestType = "ItemLevelRestore"
)

func PossibleValuesForRestoreRequestType() []string {
	return []string{
		string(RestoreRequestTypeFullShareRestore),
		string(RestoreRequestTypeInvalid),
		string(RestoreRequestTypeItemLevelRestore),
	}
}

func parseRestoreRequestType(input string) (*RestoreRequestType, error) {
	vals := map[string]RestoreRequestType{
		"fullsharerestore": RestoreRequestTypeFullShareRestore,
		"invalid":          RestoreRequestTypeInvalid,
		"itemlevelrestore": RestoreRequestTypeItemLevelRestore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestoreRequestType(input)
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
