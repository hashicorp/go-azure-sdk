package migrateschemasqlserversqldbtasks

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

type SchemaMigrationOption string

const (
	SchemaMigrationOptionExtractFromSource SchemaMigrationOption = "ExtractFromSource"
	SchemaMigrationOptionNone              SchemaMigrationOption = "None"
	SchemaMigrationOptionUseStorageFile    SchemaMigrationOption = "UseStorageFile"
)

func PossibleValuesForSchemaMigrationOption() []string {
	return []string{
		string(SchemaMigrationOptionExtractFromSource),
		string(SchemaMigrationOptionNone),
		string(SchemaMigrationOptionUseStorageFile),
	}
}

func (s *SchemaMigrationOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSchemaMigrationOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSchemaMigrationOption(input string) (*SchemaMigrationOption, error) {
	vals := map[string]SchemaMigrationOption{
		"extractfromsource": SchemaMigrationOptionExtractFromSource,
		"none":              SchemaMigrationOptionNone,
		"usestoragefile":    SchemaMigrationOptionUseStorageFile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SchemaMigrationOption(input)
	return &out, nil
}

type SchemaMigrationStage string

const (
	SchemaMigrationStageCollectingObjects     SchemaMigrationStage = "CollectingObjects"
	SchemaMigrationStageCompleted             SchemaMigrationStage = "Completed"
	SchemaMigrationStageCompletedWithWarnings SchemaMigrationStage = "CompletedWithWarnings"
	SchemaMigrationStageDeployingSchema       SchemaMigrationStage = "DeployingSchema"
	SchemaMigrationStageDownloadingScript     SchemaMigrationStage = "DownloadingScript"
	SchemaMigrationStageFailed                SchemaMigrationStage = "Failed"
	SchemaMigrationStageGeneratingScript      SchemaMigrationStage = "GeneratingScript"
	SchemaMigrationStageNotStarted            SchemaMigrationStage = "NotStarted"
	SchemaMigrationStageUploadingScript       SchemaMigrationStage = "UploadingScript"
	SchemaMigrationStageValidatingInputs      SchemaMigrationStage = "ValidatingInputs"
)

func PossibleValuesForSchemaMigrationStage() []string {
	return []string{
		string(SchemaMigrationStageCollectingObjects),
		string(SchemaMigrationStageCompleted),
		string(SchemaMigrationStageCompletedWithWarnings),
		string(SchemaMigrationStageDeployingSchema),
		string(SchemaMigrationStageDownloadingScript),
		string(SchemaMigrationStageFailed),
		string(SchemaMigrationStageGeneratingScript),
		string(SchemaMigrationStageNotStarted),
		string(SchemaMigrationStageUploadingScript),
		string(SchemaMigrationStageValidatingInputs),
	}
}

func (s *SchemaMigrationStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSchemaMigrationStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSchemaMigrationStage(input string) (*SchemaMigrationStage, error) {
	vals := map[string]SchemaMigrationStage{
		"collectingobjects":     SchemaMigrationStageCollectingObjects,
		"completed":             SchemaMigrationStageCompleted,
		"completedwithwarnings": SchemaMigrationStageCompletedWithWarnings,
		"deployingschema":       SchemaMigrationStageDeployingSchema,
		"downloadingscript":     SchemaMigrationStageDownloadingScript,
		"failed":                SchemaMigrationStageFailed,
		"generatingscript":      SchemaMigrationStageGeneratingScript,
		"notstarted":            SchemaMigrationStageNotStarted,
		"uploadingscript":       SchemaMigrationStageUploadingScript,
		"validatinginputs":      SchemaMigrationStageValidatingInputs,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SchemaMigrationStage(input)
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
