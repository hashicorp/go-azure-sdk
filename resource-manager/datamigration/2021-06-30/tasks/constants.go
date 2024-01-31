package tasks

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

type MongoDbClusterType string

const (
	MongoDbClusterTypeBlobContainer MongoDbClusterType = "BlobContainer"
	MongoDbClusterTypeCosmosDb      MongoDbClusterType = "CosmosDb"
	MongoDbClusterTypeMongoDb       MongoDbClusterType = "MongoDb"
)

func PossibleValuesForMongoDbClusterType() []string {
	return []string{
		string(MongoDbClusterTypeBlobContainer),
		string(MongoDbClusterTypeCosmosDb),
		string(MongoDbClusterTypeMongoDb),
	}
}

func (s *MongoDbClusterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMongoDbClusterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMongoDbClusterType(input string) (*MongoDbClusterType, error) {
	vals := map[string]MongoDbClusterType{
		"blobcontainer": MongoDbClusterTypeBlobContainer,
		"cosmosdb":      MongoDbClusterTypeCosmosDb,
		"mongodb":       MongoDbClusterTypeMongoDb,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MongoDbClusterType(input)
	return &out, nil
}

type MongoDbErrorType string

const (
	MongoDbErrorTypeError           MongoDbErrorType = "Error"
	MongoDbErrorTypeValidationError MongoDbErrorType = "ValidationError"
	MongoDbErrorTypeWarning         MongoDbErrorType = "Warning"
)

func PossibleValuesForMongoDbErrorType() []string {
	return []string{
		string(MongoDbErrorTypeError),
		string(MongoDbErrorTypeValidationError),
		string(MongoDbErrorTypeWarning),
	}
}

func (s *MongoDbErrorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMongoDbErrorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMongoDbErrorType(input string) (*MongoDbErrorType, error) {
	vals := map[string]MongoDbErrorType{
		"error":           MongoDbErrorTypeError,
		"validationerror": MongoDbErrorTypeValidationError,
		"warning":         MongoDbErrorTypeWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MongoDbErrorType(input)
	return &out, nil
}

type MongoDbMigrationState string

const (
	MongoDbMigrationStateCanceled        MongoDbMigrationState = "Canceled"
	MongoDbMigrationStateComplete        MongoDbMigrationState = "Complete"
	MongoDbMigrationStateCopying         MongoDbMigrationState = "Copying"
	MongoDbMigrationStateFailed          MongoDbMigrationState = "Failed"
	MongoDbMigrationStateFinalizing      MongoDbMigrationState = "Finalizing"
	MongoDbMigrationStateInitialReplay   MongoDbMigrationState = "InitialReplay"
	MongoDbMigrationStateInitializing    MongoDbMigrationState = "Initializing"
	MongoDbMigrationStateNotStarted      MongoDbMigrationState = "NotStarted"
	MongoDbMigrationStateReplaying       MongoDbMigrationState = "Replaying"
	MongoDbMigrationStateRestarting      MongoDbMigrationState = "Restarting"
	MongoDbMigrationStateValidatingInput MongoDbMigrationState = "ValidatingInput"
)

func PossibleValuesForMongoDbMigrationState() []string {
	return []string{
		string(MongoDbMigrationStateCanceled),
		string(MongoDbMigrationStateComplete),
		string(MongoDbMigrationStateCopying),
		string(MongoDbMigrationStateFailed),
		string(MongoDbMigrationStateFinalizing),
		string(MongoDbMigrationStateInitialReplay),
		string(MongoDbMigrationStateInitializing),
		string(MongoDbMigrationStateNotStarted),
		string(MongoDbMigrationStateReplaying),
		string(MongoDbMigrationStateRestarting),
		string(MongoDbMigrationStateValidatingInput),
	}
}

func (s *MongoDbMigrationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMongoDbMigrationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMongoDbMigrationState(input string) (*MongoDbMigrationState, error) {
	vals := map[string]MongoDbMigrationState{
		"canceled":        MongoDbMigrationStateCanceled,
		"complete":        MongoDbMigrationStateComplete,
		"copying":         MongoDbMigrationStateCopying,
		"failed":          MongoDbMigrationStateFailed,
		"finalizing":      MongoDbMigrationStateFinalizing,
		"initialreplay":   MongoDbMigrationStateInitialReplay,
		"initializing":    MongoDbMigrationStateInitializing,
		"notstarted":      MongoDbMigrationStateNotStarted,
		"replaying":       MongoDbMigrationStateReplaying,
		"restarting":      MongoDbMigrationStateRestarting,
		"validatinginput": MongoDbMigrationStateValidatingInput,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MongoDbMigrationState(input)
	return &out, nil
}

type MongoDbReplication string

const (
	MongoDbReplicationContinuous MongoDbReplication = "Continuous"
	MongoDbReplicationDisabled   MongoDbReplication = "Disabled"
	MongoDbReplicationOneTime    MongoDbReplication = "OneTime"
)

func PossibleValuesForMongoDbReplication() []string {
	return []string{
		string(MongoDbReplicationContinuous),
		string(MongoDbReplicationDisabled),
		string(MongoDbReplicationOneTime),
	}
}

func (s *MongoDbReplication) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMongoDbReplication(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMongoDbReplication(input string) (*MongoDbReplication, error) {
	vals := map[string]MongoDbReplication{
		"continuous": MongoDbReplicationContinuous,
		"disabled":   MongoDbReplicationDisabled,
		"onetime":    MongoDbReplicationOneTime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MongoDbReplication(input)
	return &out, nil
}

type MongoDbShardKeyOrder string

const (
	MongoDbShardKeyOrderForward MongoDbShardKeyOrder = "Forward"
	MongoDbShardKeyOrderHashed  MongoDbShardKeyOrder = "Hashed"
	MongoDbShardKeyOrderReverse MongoDbShardKeyOrder = "Reverse"
)

func PossibleValuesForMongoDbShardKeyOrder() []string {
	return []string{
		string(MongoDbShardKeyOrderForward),
		string(MongoDbShardKeyOrderHashed),
		string(MongoDbShardKeyOrderReverse),
	}
}

func (s *MongoDbShardKeyOrder) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMongoDbShardKeyOrder(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMongoDbShardKeyOrder(input string) (*MongoDbShardKeyOrder, error) {
	vals := map[string]MongoDbShardKeyOrder{
		"forward": MongoDbShardKeyOrderForward,
		"hashed":  MongoDbShardKeyOrderHashed,
		"reverse": MongoDbShardKeyOrderReverse,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MongoDbShardKeyOrder(input)
	return &out, nil
}

type MySqlTargetPlatformType string

const (
	MySqlTargetPlatformTypeAzureDbForMySQL MySqlTargetPlatformType = "AzureDbForMySQL"
	MySqlTargetPlatformTypeSqlServer       MySqlTargetPlatformType = "SqlServer"
)

func PossibleValuesForMySqlTargetPlatformType() []string {
	return []string{
		string(MySqlTargetPlatformTypeAzureDbForMySQL),
		string(MySqlTargetPlatformTypeSqlServer),
	}
}

func (s *MySqlTargetPlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMySqlTargetPlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMySqlTargetPlatformType(input string) (*MySqlTargetPlatformType, error) {
	vals := map[string]MySqlTargetPlatformType{
		"azuredbformysql": MySqlTargetPlatformTypeAzureDbForMySQL,
		"sqlserver":       MySqlTargetPlatformTypeSqlServer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MySqlTargetPlatformType(input)
	return &out, nil
}

type ResultType string

const (
	ResultTypeCollection ResultType = "Collection"
	ResultTypeDatabase   ResultType = "Database"
	ResultTypeMigration  ResultType = "Migration"
)

func PossibleValuesForResultType() []string {
	return []string{
		string(ResultTypeCollection),
		string(ResultTypeDatabase),
		string(ResultTypeMigration),
	}
}

func (s *ResultType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResultType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResultType(input string) (*ResultType, error) {
	vals := map[string]ResultType{
		"collection": ResultTypeCollection,
		"database":   ResultTypeDatabase,
		"migration":  ResultTypeMigration,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResultType(input)
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

type SsisMigrationOverwriteOption string

const (
	SsisMigrationOverwriteOptionIgnore    SsisMigrationOverwriteOption = "Ignore"
	SsisMigrationOverwriteOptionOverwrite SsisMigrationOverwriteOption = "Overwrite"
)

func PossibleValuesForSsisMigrationOverwriteOption() []string {
	return []string{
		string(SsisMigrationOverwriteOptionIgnore),
		string(SsisMigrationOverwriteOptionOverwrite),
	}
}

func (s *SsisMigrationOverwriteOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSsisMigrationOverwriteOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSsisMigrationOverwriteOption(input string) (*SsisMigrationOverwriteOption, error) {
	vals := map[string]SsisMigrationOverwriteOption{
		"ignore":    SsisMigrationOverwriteOptionIgnore,
		"overwrite": SsisMigrationOverwriteOptionOverwrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SsisMigrationOverwriteOption(input)
	return &out, nil
}

type SsisStoreType string

const (
	SsisStoreTypeSsisCatalog SsisStoreType = "SsisCatalog"
)

func PossibleValuesForSsisStoreType() []string {
	return []string{
		string(SsisStoreTypeSsisCatalog),
	}
}

func (s *SsisStoreType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSsisStoreType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSsisStoreType(input string) (*SsisStoreType, error) {
	vals := map[string]SsisStoreType{
		"ssiscatalog": SsisStoreTypeSsisCatalog,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SsisStoreType(input)
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
