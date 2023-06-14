package databases

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CatalogCollationType string

const (
	CatalogCollationTypeDATABASEDEFAULT             CatalogCollationType = "DATABASE_DEFAULT"
	CatalogCollationTypeSQLLatinOneGeneralCPOneCIAS CatalogCollationType = "SQL_Latin1_General_CP1_CI_AS"
)

func PossibleValuesForCatalogCollationType() []string {
	return []string{
		string(CatalogCollationTypeDATABASEDEFAULT),
		string(CatalogCollationTypeSQLLatinOneGeneralCPOneCIAS),
	}
}

func (s *CatalogCollationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCatalogCollationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCatalogCollationType(input string) (*CatalogCollationType, error) {
	vals := map[string]CatalogCollationType{
		"database_default":             CatalogCollationTypeDATABASEDEFAULT,
		"sql_latin1_general_cp1_ci_as": CatalogCollationTypeSQLLatinOneGeneralCPOneCIAS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CatalogCollationType(input)
	return &out, nil
}

type CreateMode string

const (
	CreateModeCopy                           CreateMode = "Copy"
	CreateModeDefault                        CreateMode = "Default"
	CreateModeOnlineSecondary                CreateMode = "OnlineSecondary"
	CreateModePointInTimeRestore             CreateMode = "PointInTimeRestore"
	CreateModeRecovery                       CreateMode = "Recovery"
	CreateModeRestore                        CreateMode = "Restore"
	CreateModeRestoreExternalBackup          CreateMode = "RestoreExternalBackup"
	CreateModeRestoreExternalBackupSecondary CreateMode = "RestoreExternalBackupSecondary"
	CreateModeRestoreLongTermRetentionBackup CreateMode = "RestoreLongTermRetentionBackup"
	CreateModeSecondary                      CreateMode = "Secondary"
)

func PossibleValuesForCreateMode() []string {
	return []string{
		string(CreateModeCopy),
		string(CreateModeDefault),
		string(CreateModeOnlineSecondary),
		string(CreateModePointInTimeRestore),
		string(CreateModeRecovery),
		string(CreateModeRestore),
		string(CreateModeRestoreExternalBackup),
		string(CreateModeRestoreExternalBackupSecondary),
		string(CreateModeRestoreLongTermRetentionBackup),
		string(CreateModeSecondary),
	}
}

func (s *CreateMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMode(input string) (*CreateMode, error) {
	vals := map[string]CreateMode{
		"copy":                           CreateModeCopy,
		"default":                        CreateModeDefault,
		"onlinesecondary":                CreateModeOnlineSecondary,
		"pointintimerestore":             CreateModePointInTimeRestore,
		"recovery":                       CreateModeRecovery,
		"restore":                        CreateModeRestore,
		"restoreexternalbackup":          CreateModeRestoreExternalBackup,
		"restoreexternalbackupsecondary": CreateModeRestoreExternalBackupSecondary,
		"restorelongtermretentionbackup": CreateModeRestoreLongTermRetentionBackup,
		"secondary":                      CreateModeSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMode(input)
	return &out, nil
}

type DatabaseStatus string

const (
	DatabaseStatusAutoClosed       DatabaseStatus = "AutoClosed"
	DatabaseStatusCopying          DatabaseStatus = "Copying"
	DatabaseStatusCreating         DatabaseStatus = "Creating"
	DatabaseStatusEmergencyMode    DatabaseStatus = "EmergencyMode"
	DatabaseStatusInaccessible     DatabaseStatus = "Inaccessible"
	DatabaseStatusOffline          DatabaseStatus = "Offline"
	DatabaseStatusOfflineSecondary DatabaseStatus = "OfflineSecondary"
	DatabaseStatusOnline           DatabaseStatus = "Online"
	DatabaseStatusPaused           DatabaseStatus = "Paused"
	DatabaseStatusPausing          DatabaseStatus = "Pausing"
	DatabaseStatusRecovering       DatabaseStatus = "Recovering"
	DatabaseStatusRecoveryPending  DatabaseStatus = "RecoveryPending"
	DatabaseStatusRestoring        DatabaseStatus = "Restoring"
	DatabaseStatusResuming         DatabaseStatus = "Resuming"
	DatabaseStatusScaling          DatabaseStatus = "Scaling"
	DatabaseStatusShutdown         DatabaseStatus = "Shutdown"
	DatabaseStatusStandby          DatabaseStatus = "Standby"
	DatabaseStatusSuspect          DatabaseStatus = "Suspect"
)

func PossibleValuesForDatabaseStatus() []string {
	return []string{
		string(DatabaseStatusAutoClosed),
		string(DatabaseStatusCopying),
		string(DatabaseStatusCreating),
		string(DatabaseStatusEmergencyMode),
		string(DatabaseStatusInaccessible),
		string(DatabaseStatusOffline),
		string(DatabaseStatusOfflineSecondary),
		string(DatabaseStatusOnline),
		string(DatabaseStatusPaused),
		string(DatabaseStatusPausing),
		string(DatabaseStatusRecovering),
		string(DatabaseStatusRecoveryPending),
		string(DatabaseStatusRestoring),
		string(DatabaseStatusResuming),
		string(DatabaseStatusScaling),
		string(DatabaseStatusShutdown),
		string(DatabaseStatusStandby),
		string(DatabaseStatusSuspect),
	}
}

func (s *DatabaseStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseStatus(input string) (*DatabaseStatus, error) {
	vals := map[string]DatabaseStatus{
		"autoclosed":       DatabaseStatusAutoClosed,
		"copying":          DatabaseStatusCopying,
		"creating":         DatabaseStatusCreating,
		"emergencymode":    DatabaseStatusEmergencyMode,
		"inaccessible":     DatabaseStatusInaccessible,
		"offline":          DatabaseStatusOffline,
		"offlinesecondary": DatabaseStatusOfflineSecondary,
		"online":           DatabaseStatusOnline,
		"paused":           DatabaseStatusPaused,
		"pausing":          DatabaseStatusPausing,
		"recovering":       DatabaseStatusRecovering,
		"recoverypending":  DatabaseStatusRecoveryPending,
		"restoring":        DatabaseStatusRestoring,
		"resuming":         DatabaseStatusResuming,
		"scaling":          DatabaseStatusScaling,
		"shutdown":         DatabaseStatusShutdown,
		"standby":          DatabaseStatusStandby,
		"suspect":          DatabaseStatusSuspect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseStatus(input)
	return &out, nil
}

type ManagementOperationState string

const (
	ManagementOperationStateCancelInProgress ManagementOperationState = "CancelInProgress"
	ManagementOperationStateCancelled        ManagementOperationState = "Cancelled"
	ManagementOperationStateFailed           ManagementOperationState = "Failed"
	ManagementOperationStateInProgress       ManagementOperationState = "InProgress"
	ManagementOperationStatePending          ManagementOperationState = "Pending"
	ManagementOperationStateSucceeded        ManagementOperationState = "Succeeded"
)

func PossibleValuesForManagementOperationState() []string {
	return []string{
		string(ManagementOperationStateCancelInProgress),
		string(ManagementOperationStateCancelled),
		string(ManagementOperationStateFailed),
		string(ManagementOperationStateInProgress),
		string(ManagementOperationStatePending),
		string(ManagementOperationStateSucceeded),
	}
}

func (s *ManagementOperationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagementOperationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagementOperationState(input string) (*ManagementOperationState, error) {
	vals := map[string]ManagementOperationState{
		"cancelinprogress": ManagementOperationStateCancelInProgress,
		"cancelled":        ManagementOperationStateCancelled,
		"failed":           ManagementOperationStateFailed,
		"inprogress":       ManagementOperationStateInProgress,
		"pending":          ManagementOperationStatePending,
		"succeeded":        ManagementOperationStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagementOperationState(input)
	return &out, nil
}

type SampleName string

const (
	SampleNameAdventureWorksLT       SampleName = "AdventureWorksLT"
	SampleNameWideWorldImportersFull SampleName = "WideWorldImportersFull"
	SampleNameWideWorldImportersStd  SampleName = "WideWorldImportersStd"
)

func PossibleValuesForSampleName() []string {
	return []string{
		string(SampleNameAdventureWorksLT),
		string(SampleNameWideWorldImportersFull),
		string(SampleNameWideWorldImportersStd),
	}
}

func (s *SampleName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSampleName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSampleName(input string) (*SampleName, error) {
	vals := map[string]SampleName{
		"adventureworkslt":       SampleNameAdventureWorksLT,
		"wideworldimportersfull": SampleNameWideWorldImportersFull,
		"wideworldimportersstd":  SampleNameWideWorldImportersStd,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SampleName(input)
	return &out, nil
}
