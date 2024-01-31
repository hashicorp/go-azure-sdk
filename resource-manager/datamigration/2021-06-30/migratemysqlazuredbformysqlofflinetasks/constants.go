package migratemysqlazuredbformysqlofflinetasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseMigrationStage string

const (
	DatabaseMigrationStageBackup     DatabaseMigrationStage = "Backup"
	DatabaseMigrationStageCompleted  DatabaseMigrationStage = "Completed"
	DatabaseMigrationStageFileCopy   DatabaseMigrationStage = "FileCopy"
	DatabaseMigrationStageInitialize DatabaseMigrationStage = "Initialize"
	DatabaseMigrationStageNone       DatabaseMigrationStage = "None"
	DatabaseMigrationStageRestore    DatabaseMigrationStage = "Restore"
)

func PossibleValuesForDatabaseMigrationStage() []string {
	return []string{
		string(DatabaseMigrationStageBackup),
		string(DatabaseMigrationStageCompleted),
		string(DatabaseMigrationStageFileCopy),
		string(DatabaseMigrationStageInitialize),
		string(DatabaseMigrationStageNone),
		string(DatabaseMigrationStageRestore),
	}
}

func (s *DatabaseMigrationStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseMigrationStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseMigrationStage(input string) (*DatabaseMigrationStage, error) {
	vals := map[string]DatabaseMigrationStage{
		"backup":     DatabaseMigrationStageBackup,
		"completed":  DatabaseMigrationStageCompleted,
		"filecopy":   DatabaseMigrationStageFileCopy,
		"initialize": DatabaseMigrationStageInitialize,
		"none":       DatabaseMigrationStageNone,
		"restore":    DatabaseMigrationStageRestore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseMigrationStage(input)
	return &out, nil
}

type MigrationState string

const (
	MigrationStateCompleted  MigrationState = "Completed"
	MigrationStateFailed     MigrationState = "Failed"
	MigrationStateInProgress MigrationState = "InProgress"
	MigrationStateNone       MigrationState = "None"
	MigrationStateSkipped    MigrationState = "Skipped"
	MigrationStateStopped    MigrationState = "Stopped"
	MigrationStateWarning    MigrationState = "Warning"
)

func PossibleValuesForMigrationState() []string {
	return []string{
		string(MigrationStateCompleted),
		string(MigrationStateFailed),
		string(MigrationStateInProgress),
		string(MigrationStateNone),
		string(MigrationStateSkipped),
		string(MigrationStateStopped),
		string(MigrationStateWarning),
	}
}

func (s *MigrationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMigrationState(input string) (*MigrationState, error) {
	vals := map[string]MigrationState{
		"completed":  MigrationStateCompleted,
		"failed":     MigrationStateFailed,
		"inprogress": MigrationStateInProgress,
		"none":       MigrationStateNone,
		"skipped":    MigrationStateSkipped,
		"stopped":    MigrationStateStopped,
		"warning":    MigrationStateWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationState(input)
	return &out, nil
}

type MigrationStatus string

const (
	MigrationStatusCompleted               MigrationStatus = "Completed"
	MigrationStatusCompletedWithWarnings   MigrationStatus = "CompletedWithWarnings"
	MigrationStatusConfigured              MigrationStatus = "Configured"
	MigrationStatusConnecting              MigrationStatus = "Connecting"
	MigrationStatusDefault                 MigrationStatus = "Default"
	MigrationStatusError                   MigrationStatus = "Error"
	MigrationStatusRunning                 MigrationStatus = "Running"
	MigrationStatusSelectLogins            MigrationStatus = "SelectLogins"
	MigrationStatusSourceAndTargetSelected MigrationStatus = "SourceAndTargetSelected"
	MigrationStatusStopped                 MigrationStatus = "Stopped"
)

func PossibleValuesForMigrationStatus() []string {
	return []string{
		string(MigrationStatusCompleted),
		string(MigrationStatusCompletedWithWarnings),
		string(MigrationStatusConfigured),
		string(MigrationStatusConnecting),
		string(MigrationStatusDefault),
		string(MigrationStatusError),
		string(MigrationStatusRunning),
		string(MigrationStatusSelectLogins),
		string(MigrationStatusSourceAndTargetSelected),
		string(MigrationStatusStopped),
	}
}

func (s *MigrationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMigrationStatus(input string) (*MigrationStatus, error) {
	vals := map[string]MigrationStatus{
		"completed":               MigrationStatusCompleted,
		"completedwithwarnings":   MigrationStatusCompletedWithWarnings,
		"configured":              MigrationStatusConfigured,
		"connecting":              MigrationStatusConnecting,
		"default":                 MigrationStatusDefault,
		"error":                   MigrationStatusError,
		"running":                 MigrationStatusRunning,
		"selectlogins":            MigrationStatusSelectLogins,
		"sourceandtargetselected": MigrationStatusSourceAndTargetSelected,
		"stopped":                 MigrationStatusStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationStatus(input)
	return &out, nil
}
