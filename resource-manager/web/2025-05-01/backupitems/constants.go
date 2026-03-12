package backupitems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupItemStatus string

const (
	BackupItemStatusCreated            BackupItemStatus = "Created"
	BackupItemStatusDeleteFailed       BackupItemStatus = "DeleteFailed"
	BackupItemStatusDeleteInProgress   BackupItemStatus = "DeleteInProgress"
	BackupItemStatusDeleted            BackupItemStatus = "Deleted"
	BackupItemStatusFailed             BackupItemStatus = "Failed"
	BackupItemStatusInProgress         BackupItemStatus = "InProgress"
	BackupItemStatusPartiallySucceeded BackupItemStatus = "PartiallySucceeded"
	BackupItemStatusSkipped            BackupItemStatus = "Skipped"
	BackupItemStatusSucceeded          BackupItemStatus = "Succeeded"
	BackupItemStatusTimedOut           BackupItemStatus = "TimedOut"
)

func PossibleValuesForBackupItemStatus() []string {
	return []string{
		string(BackupItemStatusCreated),
		string(BackupItemStatusDeleteFailed),
		string(BackupItemStatusDeleteInProgress),
		string(BackupItemStatusDeleted),
		string(BackupItemStatusFailed),
		string(BackupItemStatusInProgress),
		string(BackupItemStatusPartiallySucceeded),
		string(BackupItemStatusSkipped),
		string(BackupItemStatusSucceeded),
		string(BackupItemStatusTimedOut),
	}
}

func (s *BackupItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBackupItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBackupItemStatus(input string) (*BackupItemStatus, error) {
	vals := map[string]BackupItemStatus{
		"created":            BackupItemStatusCreated,
		"deletefailed":       BackupItemStatusDeleteFailed,
		"deleteinprogress":   BackupItemStatusDeleteInProgress,
		"deleted":            BackupItemStatusDeleted,
		"failed":             BackupItemStatusFailed,
		"inprogress":         BackupItemStatusInProgress,
		"partiallysucceeded": BackupItemStatusPartiallySucceeded,
		"skipped":            BackupItemStatusSkipped,
		"succeeded":          BackupItemStatusSucceeded,
		"timedout":           BackupItemStatusTimedOut,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupItemStatus(input)
	return &out, nil
}

type BackupRestoreOperationType string

const (
	BackupRestoreOperationTypeClone      BackupRestoreOperationType = "Clone"
	BackupRestoreOperationTypeCloudFS    BackupRestoreOperationType = "CloudFS"
	BackupRestoreOperationTypeDefault    BackupRestoreOperationType = "Default"
	BackupRestoreOperationTypeRelocation BackupRestoreOperationType = "Relocation"
	BackupRestoreOperationTypeSnapshot   BackupRestoreOperationType = "Snapshot"
)

func PossibleValuesForBackupRestoreOperationType() []string {
	return []string{
		string(BackupRestoreOperationTypeClone),
		string(BackupRestoreOperationTypeCloudFS),
		string(BackupRestoreOperationTypeDefault),
		string(BackupRestoreOperationTypeRelocation),
		string(BackupRestoreOperationTypeSnapshot),
	}
}

func (s *BackupRestoreOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBackupRestoreOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBackupRestoreOperationType(input string) (*BackupRestoreOperationType, error) {
	vals := map[string]BackupRestoreOperationType{
		"clone":      BackupRestoreOperationTypeClone,
		"cloudfs":    BackupRestoreOperationTypeCloudFS,
		"default":    BackupRestoreOperationTypeDefault,
		"relocation": BackupRestoreOperationTypeRelocation,
		"snapshot":   BackupRestoreOperationTypeSnapshot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupRestoreOperationType(input)
	return &out, nil
}

type DatabaseType string

const (
	DatabaseTypeLocalMySql DatabaseType = "LocalMySql"
	DatabaseTypeMySql      DatabaseType = "MySql"
	DatabaseTypePostgreSql DatabaseType = "PostgreSql"
	DatabaseTypeSqlAzure   DatabaseType = "SqlAzure"
)

func PossibleValuesForDatabaseType() []string {
	return []string{
		string(DatabaseTypeLocalMySql),
		string(DatabaseTypeMySql),
		string(DatabaseTypePostgreSql),
		string(DatabaseTypeSqlAzure),
	}
}

func (s *DatabaseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseType(input string) (*DatabaseType, error) {
	vals := map[string]DatabaseType{
		"localmysql": DatabaseTypeLocalMySql,
		"mysql":      DatabaseTypeMySql,
		"postgresql": DatabaseTypePostgreSql,
		"sqlazure":   DatabaseTypeSqlAzure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseType(input)
	return &out, nil
}

type FrequencyUnit string

const (
	FrequencyUnitDay  FrequencyUnit = "Day"
	FrequencyUnitHour FrequencyUnit = "Hour"
)

func PossibleValuesForFrequencyUnit() []string {
	return []string{
		string(FrequencyUnitDay),
		string(FrequencyUnitHour),
	}
}

func (s *FrequencyUnit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFrequencyUnit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFrequencyUnit(input string) (*FrequencyUnit, error) {
	vals := map[string]FrequencyUnit{
		"day":  FrequencyUnitDay,
		"hour": FrequencyUnitHour,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FrequencyUnit(input)
	return &out, nil
}
