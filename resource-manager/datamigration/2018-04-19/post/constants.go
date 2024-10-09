package post

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

type BackupMode string

const (
	BackupModeCreateBackup   BackupMode = "CreateBackup"
	BackupModeExistingBackup BackupMode = "ExistingBackup"
)

func PossibleValuesForBackupMode() []string {
	return []string{
		string(BackupModeCreateBackup),
		string(BackupModeExistingBackup),
	}
}

func (s *BackupMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBackupMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBackupMode(input string) (*BackupMode, error) {
	vals := map[string]BackupMode{
		"createbackup":   BackupModeCreateBackup,
		"existingbackup": BackupModeExistingBackup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupMode(input)
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

type CommandState string

const (
	CommandStateAccepted  CommandState = "Accepted"
	CommandStateFailed    CommandState = "Failed"
	CommandStateRunning   CommandState = "Running"
	CommandStateSucceeded CommandState = "Succeeded"
	CommandStateUnknown   CommandState = "Unknown"
)

func PossibleValuesForCommandState() []string {
	return []string{
		string(CommandStateAccepted),
		string(CommandStateFailed),
		string(CommandStateRunning),
		string(CommandStateSucceeded),
		string(CommandStateUnknown),
	}
}

func (s *CommandState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCommandState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCommandState(input string) (*CommandState, error) {
	vals := map[string]CommandState{
		"accepted":  CommandStateAccepted,
		"failed":    CommandStateFailed,
		"running":   CommandStateRunning,
		"succeeded": CommandStateSucceeded,
		"unknown":   CommandStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommandState(input)
	return &out, nil
}

type DatabaseCompatLevel string

const (
	DatabaseCompatLevelCompatLevelEightZero    DatabaseCompatLevel = "CompatLevel80"
	DatabaseCompatLevelCompatLevelNineZero     DatabaseCompatLevel = "CompatLevel90"
	DatabaseCompatLevelCompatLevelOneFourZero  DatabaseCompatLevel = "CompatLevel140"
	DatabaseCompatLevelCompatLevelOneHundred   DatabaseCompatLevel = "CompatLevel100"
	DatabaseCompatLevelCompatLevelOneOneZero   DatabaseCompatLevel = "CompatLevel110"
	DatabaseCompatLevelCompatLevelOneThreeZero DatabaseCompatLevel = "CompatLevel130"
	DatabaseCompatLevelCompatLevelOneTwoZero   DatabaseCompatLevel = "CompatLevel120"
)

func PossibleValuesForDatabaseCompatLevel() []string {
	return []string{
		string(DatabaseCompatLevelCompatLevelEightZero),
		string(DatabaseCompatLevelCompatLevelNineZero),
		string(DatabaseCompatLevelCompatLevelOneFourZero),
		string(DatabaseCompatLevelCompatLevelOneHundred),
		string(DatabaseCompatLevelCompatLevelOneOneZero),
		string(DatabaseCompatLevelCompatLevelOneThreeZero),
		string(DatabaseCompatLevelCompatLevelOneTwoZero),
	}
}

func (s *DatabaseCompatLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseCompatLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseCompatLevel(input string) (*DatabaseCompatLevel, error) {
	vals := map[string]DatabaseCompatLevel{
		"compatlevel80":  DatabaseCompatLevelCompatLevelEightZero,
		"compatlevel90":  DatabaseCompatLevelCompatLevelNineZero,
		"compatlevel140": DatabaseCompatLevelCompatLevelOneFourZero,
		"compatlevel100": DatabaseCompatLevelCompatLevelOneHundred,
		"compatlevel110": DatabaseCompatLevelCompatLevelOneOneZero,
		"compatlevel130": DatabaseCompatLevelCompatLevelOneThreeZero,
		"compatlevel120": DatabaseCompatLevelCompatLevelOneTwoZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseCompatLevel(input)
	return &out, nil
}

type DatabaseFileType string

const (
	DatabaseFileTypeFilestream   DatabaseFileType = "Filestream"
	DatabaseFileTypeFulltext     DatabaseFileType = "Fulltext"
	DatabaseFileTypeLog          DatabaseFileType = "Log"
	DatabaseFileTypeNotSupported DatabaseFileType = "NotSupported"
	DatabaseFileTypeRows         DatabaseFileType = "Rows"
)

func PossibleValuesForDatabaseFileType() []string {
	return []string{
		string(DatabaseFileTypeFilestream),
		string(DatabaseFileTypeFulltext),
		string(DatabaseFileTypeLog),
		string(DatabaseFileTypeNotSupported),
		string(DatabaseFileTypeRows),
	}
}

func (s *DatabaseFileType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseFileType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseFileType(input string) (*DatabaseFileType, error) {
	vals := map[string]DatabaseFileType{
		"filestream":   DatabaseFileTypeFilestream,
		"fulltext":     DatabaseFileTypeFulltext,
		"log":          DatabaseFileTypeLog,
		"notsupported": DatabaseFileTypeNotSupported,
		"rows":         DatabaseFileTypeRows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseFileType(input)
	return &out, nil
}

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

type DatabaseState string

const (
	DatabaseStateCopying          DatabaseState = "Copying"
	DatabaseStateEmergency        DatabaseState = "Emergency"
	DatabaseStateOffline          DatabaseState = "Offline"
	DatabaseStateOfflineSecondary DatabaseState = "OfflineSecondary"
	DatabaseStateOnline           DatabaseState = "Online"
	DatabaseStateRecovering       DatabaseState = "Recovering"
	DatabaseStateRecoveryPending  DatabaseState = "RecoveryPending"
	DatabaseStateRestoring        DatabaseState = "Restoring"
	DatabaseStateSuspect          DatabaseState = "Suspect"
)

func PossibleValuesForDatabaseState() []string {
	return []string{
		string(DatabaseStateCopying),
		string(DatabaseStateEmergency),
		string(DatabaseStateOffline),
		string(DatabaseStateOfflineSecondary),
		string(DatabaseStateOnline),
		string(DatabaseStateRecovering),
		string(DatabaseStateRecoveryPending),
		string(DatabaseStateRestoring),
		string(DatabaseStateSuspect),
	}
}

func (s *DatabaseState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseState(input string) (*DatabaseState, error) {
	vals := map[string]DatabaseState{
		"copying":          DatabaseStateCopying,
		"emergency":        DatabaseStateEmergency,
		"offline":          DatabaseStateOffline,
		"offlinesecondary": DatabaseStateOfflineSecondary,
		"online":           DatabaseStateOnline,
		"recovering":       DatabaseStateRecovering,
		"recoverypending":  DatabaseStateRecoveryPending,
		"restoring":        DatabaseStateRestoring,
		"suspect":          DatabaseStateSuspect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseState(input)
	return &out, nil
}

type LoginMigrationStage string

const (
	LoginMigrationStageAssignRoleMembership       LoginMigrationStage = "AssignRoleMembership"
	LoginMigrationStageAssignRoleOwnership        LoginMigrationStage = "AssignRoleOwnership"
	LoginMigrationStageCompleted                  LoginMigrationStage = "Completed"
	LoginMigrationStageEstablishObjectPermissions LoginMigrationStage = "EstablishObjectPermissions"
	LoginMigrationStageEstablishServerPermissions LoginMigrationStage = "EstablishServerPermissions"
	LoginMigrationStageEstablishUserMapping       LoginMigrationStage = "EstablishUserMapping"
	LoginMigrationStageInitialize                 LoginMigrationStage = "Initialize"
	LoginMigrationStageLoginMigration             LoginMigrationStage = "LoginMigration"
	LoginMigrationStageNone                       LoginMigrationStage = "None"
)

func PossibleValuesForLoginMigrationStage() []string {
	return []string{
		string(LoginMigrationStageAssignRoleMembership),
		string(LoginMigrationStageAssignRoleOwnership),
		string(LoginMigrationStageCompleted),
		string(LoginMigrationStageEstablishObjectPermissions),
		string(LoginMigrationStageEstablishServerPermissions),
		string(LoginMigrationStageEstablishUserMapping),
		string(LoginMigrationStageInitialize),
		string(LoginMigrationStageLoginMigration),
		string(LoginMigrationStageNone),
	}
}

func (s *LoginMigrationStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLoginMigrationStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLoginMigrationStage(input string) (*LoginMigrationStage, error) {
	vals := map[string]LoginMigrationStage{
		"assignrolemembership":       LoginMigrationStageAssignRoleMembership,
		"assignroleownership":        LoginMigrationStageAssignRoleOwnership,
		"completed":                  LoginMigrationStageCompleted,
		"establishobjectpermissions": LoginMigrationStageEstablishObjectPermissions,
		"establishserverpermissions": LoginMigrationStageEstablishServerPermissions,
		"establishusermapping":       LoginMigrationStageEstablishUserMapping,
		"initialize":                 LoginMigrationStageInitialize,
		"loginmigration":             LoginMigrationStageLoginMigration,
		"none":                       LoginMigrationStageNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoginMigrationStage(input)
	return &out, nil
}

type LoginType string

const (
	LoginTypeAsymmetricKey LoginType = "AsymmetricKey"
	LoginTypeCertificate   LoginType = "Certificate"
	LoginTypeExternalGroup LoginType = "ExternalGroup"
	LoginTypeExternalUser  LoginType = "ExternalUser"
	LoginTypeSqlLogin      LoginType = "SqlLogin"
	LoginTypeWindowsGroup  LoginType = "WindowsGroup"
	LoginTypeWindowsUser   LoginType = "WindowsUser"
)

func PossibleValuesForLoginType() []string {
	return []string{
		string(LoginTypeAsymmetricKey),
		string(LoginTypeCertificate),
		string(LoginTypeExternalGroup),
		string(LoginTypeExternalUser),
		string(LoginTypeSqlLogin),
		string(LoginTypeWindowsGroup),
		string(LoginTypeWindowsUser),
	}
}

func (s *LoginType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLoginType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLoginType(input string) (*LoginType, error) {
	vals := map[string]LoginType{
		"asymmetrickey": LoginTypeAsymmetricKey,
		"certificate":   LoginTypeCertificate,
		"externalgroup": LoginTypeExternalGroup,
		"externaluser":  LoginTypeExternalUser,
		"sqllogin":      LoginTypeSqlLogin,
		"windowsgroup":  LoginTypeWindowsGroup,
		"windowsuser":   LoginTypeWindowsUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoginType(input)
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

type NameCheckFailureReason string

const (
	NameCheckFailureReasonAlreadyExists NameCheckFailureReason = "AlreadyExists"
	NameCheckFailureReasonInvalid       NameCheckFailureReason = "Invalid"
)

func PossibleValuesForNameCheckFailureReason() []string {
	return []string{
		string(NameCheckFailureReasonAlreadyExists),
		string(NameCheckFailureReasonInvalid),
	}
}

func (s *NameCheckFailureReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNameCheckFailureReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNameCheckFailureReason(input string) (*NameCheckFailureReason, error) {
	vals := map[string]NameCheckFailureReason{
		"alreadyexists": NameCheckFailureReasonAlreadyExists,
		"invalid":       NameCheckFailureReasonInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NameCheckFailureReason(input)
	return &out, nil
}

type ObjectType string

const (
	ObjectTypeFunction         ObjectType = "Function"
	ObjectTypeStoredProcedures ObjectType = "StoredProcedures"
	ObjectTypeTable            ObjectType = "Table"
	ObjectTypeUser             ObjectType = "User"
	ObjectTypeView             ObjectType = "View"
)

func PossibleValuesForObjectType() []string {
	return []string{
		string(ObjectTypeFunction),
		string(ObjectTypeStoredProcedures),
		string(ObjectTypeTable),
		string(ObjectTypeUser),
		string(ObjectTypeView),
	}
}

func (s *ObjectType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseObjectType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseObjectType(input string) (*ObjectType, error) {
	vals := map[string]ObjectType{
		"function":         ObjectTypeFunction,
		"storedprocedures": ObjectTypeStoredProcedures,
		"table":            ObjectTypeTable,
		"user":             ObjectTypeUser,
		"view":             ObjectTypeView,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectType(input)
	return &out, nil
}

type ServerLevelPermissionsGroup string

const (
	ServerLevelPermissionsGroupDefault                             ServerLevelPermissionsGroup = "Default"
	ServerLevelPermissionsGroupMigrationFromMySQLToAzureDBForMySQL ServerLevelPermissionsGroup = "MigrationFromMySQLToAzureDBForMySQL"
	ServerLevelPermissionsGroupMigrationFromSqlServerToAzureDB     ServerLevelPermissionsGroup = "MigrationFromSqlServerToAzureDB"
	ServerLevelPermissionsGroupMigrationFromSqlServerToAzureMI     ServerLevelPermissionsGroup = "MigrationFromSqlServerToAzureMI"
)

func PossibleValuesForServerLevelPermissionsGroup() []string {
	return []string{
		string(ServerLevelPermissionsGroupDefault),
		string(ServerLevelPermissionsGroupMigrationFromMySQLToAzureDBForMySQL),
		string(ServerLevelPermissionsGroupMigrationFromSqlServerToAzureDB),
		string(ServerLevelPermissionsGroupMigrationFromSqlServerToAzureMI),
	}
}

func (s *ServerLevelPermissionsGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServerLevelPermissionsGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServerLevelPermissionsGroup(input string) (*ServerLevelPermissionsGroup, error) {
	vals := map[string]ServerLevelPermissionsGroup{
		"default":                             ServerLevelPermissionsGroupDefault,
		"migrationfrommysqltoazuredbformysql": ServerLevelPermissionsGroupMigrationFromMySQLToAzureDBForMySQL,
		"migrationfromsqlservertoazuredb":     ServerLevelPermissionsGroupMigrationFromSqlServerToAzureDB,
		"migrationfromsqlservertoazuremi":     ServerLevelPermissionsGroupMigrationFromSqlServerToAzureMI,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerLevelPermissionsGroup(input)
	return &out, nil
}

type Severity string

const (
	SeverityError   Severity = "Error"
	SeverityMessage Severity = "Message"
	SeverityWarning Severity = "Warning"
)

func PossibleValuesForSeverity() []string {
	return []string{
		string(SeverityError),
		string(SeverityMessage),
		string(SeverityWarning),
	}
}

func (s *Severity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSeverity(input string) (*Severity, error) {
	vals := map[string]Severity{
		"error":   SeverityError,
		"message": SeverityMessage,
		"warning": SeverityWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Severity(input)
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

type SyncDatabaseMigrationReportingState string

const (
	SyncDatabaseMigrationReportingStateCANCELLED       SyncDatabaseMigrationReportingState = "CANCELLED"
	SyncDatabaseMigrationReportingStateCANCELLING      SyncDatabaseMigrationReportingState = "CANCELLING"
	SyncDatabaseMigrationReportingStateCOMPLETE        SyncDatabaseMigrationReportingState = "COMPLETE"
	SyncDatabaseMigrationReportingStateCOMPLETING      SyncDatabaseMigrationReportingState = "COMPLETING"
	SyncDatabaseMigrationReportingStateCONFIGURING     SyncDatabaseMigrationReportingState = "CONFIGURING"
	SyncDatabaseMigrationReportingStateFAILED          SyncDatabaseMigrationReportingState = "FAILED"
	SyncDatabaseMigrationReportingStateINITIALIAZING   SyncDatabaseMigrationReportingState = "INITIALIAZING"
	SyncDatabaseMigrationReportingStateREADYTOCOMPLETE SyncDatabaseMigrationReportingState = "READY_TO_COMPLETE"
	SyncDatabaseMigrationReportingStateRUNNING         SyncDatabaseMigrationReportingState = "RUNNING"
	SyncDatabaseMigrationReportingStateSTARTING        SyncDatabaseMigrationReportingState = "STARTING"
	SyncDatabaseMigrationReportingStateUNDEFINED       SyncDatabaseMigrationReportingState = "UNDEFINED"
)

func PossibleValuesForSyncDatabaseMigrationReportingState() []string {
	return []string{
		string(SyncDatabaseMigrationReportingStateCANCELLED),
		string(SyncDatabaseMigrationReportingStateCANCELLING),
		string(SyncDatabaseMigrationReportingStateCOMPLETE),
		string(SyncDatabaseMigrationReportingStateCOMPLETING),
		string(SyncDatabaseMigrationReportingStateCONFIGURING),
		string(SyncDatabaseMigrationReportingStateFAILED),
		string(SyncDatabaseMigrationReportingStateINITIALIAZING),
		string(SyncDatabaseMigrationReportingStateREADYTOCOMPLETE),
		string(SyncDatabaseMigrationReportingStateRUNNING),
		string(SyncDatabaseMigrationReportingStateSTARTING),
		string(SyncDatabaseMigrationReportingStateUNDEFINED),
	}
}

func (s *SyncDatabaseMigrationReportingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSyncDatabaseMigrationReportingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSyncDatabaseMigrationReportingState(input string) (*SyncDatabaseMigrationReportingState, error) {
	vals := map[string]SyncDatabaseMigrationReportingState{
		"cancelled":         SyncDatabaseMigrationReportingStateCANCELLED,
		"cancelling":        SyncDatabaseMigrationReportingStateCANCELLING,
		"complete":          SyncDatabaseMigrationReportingStateCOMPLETE,
		"completing":        SyncDatabaseMigrationReportingStateCOMPLETING,
		"configuring":       SyncDatabaseMigrationReportingStateCONFIGURING,
		"failed":            SyncDatabaseMigrationReportingStateFAILED,
		"initialiazing":     SyncDatabaseMigrationReportingStateINITIALIAZING,
		"ready_to_complete": SyncDatabaseMigrationReportingStateREADYTOCOMPLETE,
		"running":           SyncDatabaseMigrationReportingStateRUNNING,
		"starting":          SyncDatabaseMigrationReportingStateSTARTING,
		"undefined":         SyncDatabaseMigrationReportingStateUNDEFINED,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SyncDatabaseMigrationReportingState(input)
	return &out, nil
}

type SyncTableMigrationState string

const (
	SyncTableMigrationStateBEFORELOAD SyncTableMigrationState = "BEFORE_LOAD"
	SyncTableMigrationStateCANCELED   SyncTableMigrationState = "CANCELED"
	SyncTableMigrationStateCOMPLETED  SyncTableMigrationState = "COMPLETED"
	SyncTableMigrationStateERROR      SyncTableMigrationState = "ERROR"
	SyncTableMigrationStateFAILED     SyncTableMigrationState = "FAILED"
	SyncTableMigrationStateFULLLOAD   SyncTableMigrationState = "FULL_LOAD"
)

func PossibleValuesForSyncTableMigrationState() []string {
	return []string{
		string(SyncTableMigrationStateBEFORELOAD),
		string(SyncTableMigrationStateCANCELED),
		string(SyncTableMigrationStateCOMPLETED),
		string(SyncTableMigrationStateERROR),
		string(SyncTableMigrationStateFAILED),
		string(SyncTableMigrationStateFULLLOAD),
	}
}

func (s *SyncTableMigrationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSyncTableMigrationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSyncTableMigrationState(input string) (*SyncTableMigrationState, error) {
	vals := map[string]SyncTableMigrationState{
		"before_load": SyncTableMigrationStateBEFORELOAD,
		"canceled":    SyncTableMigrationStateCANCELED,
		"completed":   SyncTableMigrationStateCOMPLETED,
		"error":       SyncTableMigrationStateERROR,
		"failed":      SyncTableMigrationStateFAILED,
		"full_load":   SyncTableMigrationStateFULLLOAD,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SyncTableMigrationState(input)
	return &out, nil
}

type TaskState string

const (
	TaskStateCanceled              TaskState = "Canceled"
	TaskStateFailed                TaskState = "Failed"
	TaskStateFailedInputValidation TaskState = "FailedInputValidation"
	TaskStateFaulted               TaskState = "Faulted"
	TaskStateQueued                TaskState = "Queued"
	TaskStateRunning               TaskState = "Running"
	TaskStateSucceeded             TaskState = "Succeeded"
	TaskStateUnknown               TaskState = "Unknown"
)

func PossibleValuesForTaskState() []string {
	return []string{
		string(TaskStateCanceled),
		string(TaskStateFailed),
		string(TaskStateFailedInputValidation),
		string(TaskStateFaulted),
		string(TaskStateQueued),
		string(TaskStateRunning),
		string(TaskStateSucceeded),
		string(TaskStateUnknown),
	}
}

func (s *TaskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTaskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTaskState(input string) (*TaskState, error) {
	vals := map[string]TaskState{
		"canceled":              TaskStateCanceled,
		"failed":                TaskStateFailed,
		"failedinputvalidation": TaskStateFailedInputValidation,
		"faulted":               TaskStateFaulted,
		"queued":                TaskStateQueued,
		"running":               TaskStateRunning,
		"succeeded":             TaskStateSucceeded,
		"unknown":               TaskStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TaskState(input)
	return &out, nil
}

type UpdateActionType string

const (
	UpdateActionTypeAddedOnTarget   UpdateActionType = "AddedOnTarget"
	UpdateActionTypeChangedOnTarget UpdateActionType = "ChangedOnTarget"
	UpdateActionTypeDeletedOnTarget UpdateActionType = "DeletedOnTarget"
)

func PossibleValuesForUpdateActionType() []string {
	return []string{
		string(UpdateActionTypeAddedOnTarget),
		string(UpdateActionTypeChangedOnTarget),
		string(UpdateActionTypeDeletedOnTarget),
	}
}

func (s *UpdateActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdateActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdateActionType(input string) (*UpdateActionType, error) {
	vals := map[string]UpdateActionType{
		"addedontarget":   UpdateActionTypeAddedOnTarget,
		"changedontarget": UpdateActionTypeChangedOnTarget,
		"deletedontarget": UpdateActionTypeDeletedOnTarget,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateActionType(input)
	return &out, nil
}

type ValidationStatus string

const (
	ValidationStatusCompleted           ValidationStatus = "Completed"
	ValidationStatusCompletedWithIssues ValidationStatus = "CompletedWithIssues"
	ValidationStatusDefault             ValidationStatus = "Default"
	ValidationStatusFailed              ValidationStatus = "Failed"
	ValidationStatusInProgress          ValidationStatus = "InProgress"
	ValidationStatusInitialized         ValidationStatus = "Initialized"
	ValidationStatusNotStarted          ValidationStatus = "NotStarted"
	ValidationStatusStopped             ValidationStatus = "Stopped"
)

func PossibleValuesForValidationStatus() []string {
	return []string{
		string(ValidationStatusCompleted),
		string(ValidationStatusCompletedWithIssues),
		string(ValidationStatusDefault),
		string(ValidationStatusFailed),
		string(ValidationStatusInProgress),
		string(ValidationStatusInitialized),
		string(ValidationStatusNotStarted),
		string(ValidationStatusStopped),
	}
}

func (s *ValidationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseValidationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseValidationStatus(input string) (*ValidationStatus, error) {
	vals := map[string]ValidationStatus{
		"completed":           ValidationStatusCompleted,
		"completedwithissues": ValidationStatusCompletedWithIssues,
		"default":             ValidationStatusDefault,
		"failed":              ValidationStatusFailed,
		"inprogress":          ValidationStatusInProgress,
		"initialized":         ValidationStatusInitialized,
		"notstarted":          ValidationStatusNotStarted,
		"stopped":             ValidationStatusStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ValidationStatus(input)
	return &out, nil
}
