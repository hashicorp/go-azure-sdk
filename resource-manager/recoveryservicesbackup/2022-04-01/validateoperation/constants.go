package validateoperation

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
