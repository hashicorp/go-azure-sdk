package migratesqlserversqlmisynctasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationType string

const (
	AuthenticationTypeActiveDirectoryIntegrated AuthenticationType = "ActiveDirectoryIntegrated"
	AuthenticationTypeActiveDirectoryPassword   AuthenticationType = "ActiveDirectoryPassword"
	AuthenticationTypeNone                      AuthenticationType = "None"
	AuthenticationTypeSqlAuthentication         AuthenticationType = "SqlAuthentication"
	AuthenticationTypeWindowsAuthentication     AuthenticationType = "WindowsAuthentication"
)

func PossibleValuesForAuthenticationType() []string {
	return []string{
		string(AuthenticationTypeActiveDirectoryIntegrated),
		string(AuthenticationTypeActiveDirectoryPassword),
		string(AuthenticationTypeNone),
		string(AuthenticationTypeSqlAuthentication),
		string(AuthenticationTypeWindowsAuthentication),
	}
}

func (s *AuthenticationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationType(input string) (*AuthenticationType, error) {
	vals := map[string]AuthenticationType{
		"activedirectoryintegrated": AuthenticationTypeActiveDirectoryIntegrated,
		"activedirectorypassword":   AuthenticationTypeActiveDirectoryPassword,
		"none":                      AuthenticationTypeNone,
		"sqlauthentication":         AuthenticationTypeSqlAuthentication,
		"windowsauthentication":     AuthenticationTypeWindowsAuthentication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationType(input)
	return &out, nil
}

type BackupFileStatus string

const (
	BackupFileStatusArrived   BackupFileStatus = "Arrived"
	BackupFileStatusCancelled BackupFileStatus = "Cancelled"
	BackupFileStatusQueued    BackupFileStatus = "Queued"
	BackupFileStatusRestored  BackupFileStatus = "Restored"
	BackupFileStatusRestoring BackupFileStatus = "Restoring"
	BackupFileStatusUploaded  BackupFileStatus = "Uploaded"
	BackupFileStatusUploading BackupFileStatus = "Uploading"
)

func PossibleValuesForBackupFileStatus() []string {
	return []string{
		string(BackupFileStatusArrived),
		string(BackupFileStatusCancelled),
		string(BackupFileStatusQueued),
		string(BackupFileStatusRestored),
		string(BackupFileStatusRestoring),
		string(BackupFileStatusUploaded),
		string(BackupFileStatusUploading),
	}
}

func (s *BackupFileStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBackupFileStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBackupFileStatus(input string) (*BackupFileStatus, error) {
	vals := map[string]BackupFileStatus{
		"arrived":   BackupFileStatusArrived,
		"cancelled": BackupFileStatusCancelled,
		"queued":    BackupFileStatusQueued,
		"restored":  BackupFileStatusRestored,
		"restoring": BackupFileStatusRestoring,
		"uploaded":  BackupFileStatusUploaded,
		"uploading": BackupFileStatusUploading,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupFileStatus(input)
	return &out, nil
}

type BackupType string

const (
	BackupTypeDatabase             BackupType = "Database"
	BackupTypeDifferentialDatabase BackupType = "DifferentialDatabase"
	BackupTypeDifferentialFile     BackupType = "DifferentialFile"
	BackupTypeDifferentialPartial  BackupType = "DifferentialPartial"
	BackupTypeFile                 BackupType = "File"
	BackupTypePartial              BackupType = "Partial"
	BackupTypeTransactionLog       BackupType = "TransactionLog"
)

func PossibleValuesForBackupType() []string {
	return []string{
		string(BackupTypeDatabase),
		string(BackupTypeDifferentialDatabase),
		string(BackupTypeDifferentialFile),
		string(BackupTypeDifferentialPartial),
		string(BackupTypeFile),
		string(BackupTypePartial),
		string(BackupTypeTransactionLog),
	}
}

func (s *BackupType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBackupType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBackupType(input string) (*BackupType, error) {
	vals := map[string]BackupType{
		"database":             BackupTypeDatabase,
		"differentialdatabase": BackupTypeDifferentialDatabase,
		"differentialfile":     BackupTypeDifferentialFile,
		"differentialpartial":  BackupTypeDifferentialPartial,
		"file":                 BackupTypeFile,
		"partial":              BackupTypePartial,
		"transactionlog":       BackupTypeTransactionLog,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupType(input)
	return &out, nil
}

type DatabaseMigrationState string

const (
	DatabaseMigrationStateCANCELLED             DatabaseMigrationState = "CANCELLED"
	DatabaseMigrationStateCOMPLETED             DatabaseMigrationState = "COMPLETED"
	DatabaseMigrationStateCUTOVERSTART          DatabaseMigrationState = "CUTOVER_START"
	DatabaseMigrationStateFAILED                DatabaseMigrationState = "FAILED"
	DatabaseMigrationStateFULLBACKUPUPLOADSTART DatabaseMigrationState = "FULL_BACKUP_UPLOAD_START"
	DatabaseMigrationStateINITIAL               DatabaseMigrationState = "INITIAL"
	DatabaseMigrationStateLOGSHIPPINGSTART      DatabaseMigrationState = "LOG_SHIPPING_START"
	DatabaseMigrationStatePOSTCUTOVERCOMPLETE   DatabaseMigrationState = "POST_CUTOVER_COMPLETE"
	DatabaseMigrationStateUNDEFINED             DatabaseMigrationState = "UNDEFINED"
	DatabaseMigrationStateUPLOADLOGFILESSTART   DatabaseMigrationState = "UPLOAD_LOG_FILES_START"
)

func PossibleValuesForDatabaseMigrationState() []string {
	return []string{
		string(DatabaseMigrationStateCANCELLED),
		string(DatabaseMigrationStateCOMPLETED),
		string(DatabaseMigrationStateCUTOVERSTART),
		string(DatabaseMigrationStateFAILED),
		string(DatabaseMigrationStateFULLBACKUPUPLOADSTART),
		string(DatabaseMigrationStateINITIAL),
		string(DatabaseMigrationStateLOGSHIPPINGSTART),
		string(DatabaseMigrationStatePOSTCUTOVERCOMPLETE),
		string(DatabaseMigrationStateUNDEFINED),
		string(DatabaseMigrationStateUPLOADLOGFILESSTART),
	}
}

func (s *DatabaseMigrationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseMigrationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseMigrationState(input string) (*DatabaseMigrationState, error) {
	vals := map[string]DatabaseMigrationState{
		"cancelled":                DatabaseMigrationStateCANCELLED,
		"completed":                DatabaseMigrationStateCOMPLETED,
		"cutover_start":            DatabaseMigrationStateCUTOVERSTART,
		"failed":                   DatabaseMigrationStateFAILED,
		"full_backup_upload_start": DatabaseMigrationStateFULLBACKUPUPLOADSTART,
		"initial":                  DatabaseMigrationStateINITIAL,
		"log_shipping_start":       DatabaseMigrationStateLOGSHIPPINGSTART,
		"post_cutover_complete":    DatabaseMigrationStatePOSTCUTOVERCOMPLETE,
		"undefined":                DatabaseMigrationStateUNDEFINED,
		"upload_log_files_start":   DatabaseMigrationStateUPLOADLOGFILESSTART,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseMigrationState(input)
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

type SqlSourcePlatform string

const (
	SqlSourcePlatformSqlOnPrem SqlSourcePlatform = "SqlOnPrem"
)

func PossibleValuesForSqlSourcePlatform() []string {
	return []string{
		string(SqlSourcePlatformSqlOnPrem),
	}
}

func (s *SqlSourcePlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSqlSourcePlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSqlSourcePlatform(input string) (*SqlSourcePlatform, error) {
	vals := map[string]SqlSourcePlatform{
		"sqlonprem": SqlSourcePlatformSqlOnPrem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SqlSourcePlatform(input)
	return &out, nil
}
