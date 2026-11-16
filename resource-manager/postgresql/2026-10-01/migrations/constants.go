package migrations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Cancel string

const (
	CancelFalse Cancel = "False"
	CancelTrue  Cancel = "True"
)

func PossibleValuesForCancel() []string {
	return []string{
		string(CancelFalse),
		string(CancelTrue),
	}
}

func (s *Cancel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCancel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCancel(input string) (*Cancel, error) {
	vals := map[string]Cancel{
		"false": CancelFalse,
		"true":  CancelTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Cancel(input)
	return &out, nil
}

type LogicalReplicationOnSourceServer string

const (
	LogicalReplicationOnSourceServerFalse LogicalReplicationOnSourceServer = "False"
	LogicalReplicationOnSourceServerTrue  LogicalReplicationOnSourceServer = "True"
)

func PossibleValuesForLogicalReplicationOnSourceServer() []string {
	return []string{
		string(LogicalReplicationOnSourceServerFalse),
		string(LogicalReplicationOnSourceServerTrue),
	}
}

func (s *LogicalReplicationOnSourceServer) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLogicalReplicationOnSourceServer(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLogicalReplicationOnSourceServer(input string) (*LogicalReplicationOnSourceServer, error) {
	vals := map[string]LogicalReplicationOnSourceServer{
		"false": LogicalReplicationOnSourceServerFalse,
		"true":  LogicalReplicationOnSourceServerTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogicalReplicationOnSourceServer(input)
	return &out, nil
}

type MigrateRolesAndPermissions string

const (
	MigrateRolesAndPermissionsFalse MigrateRolesAndPermissions = "False"
	MigrateRolesAndPermissionsTrue  MigrateRolesAndPermissions = "True"
)

func PossibleValuesForMigrateRolesAndPermissions() []string {
	return []string{
		string(MigrateRolesAndPermissionsFalse),
		string(MigrateRolesAndPermissionsTrue),
	}
}

func (s *MigrateRolesAndPermissions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrateRolesAndPermissions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMigrateRolesAndPermissions(input string) (*MigrateRolesAndPermissions, error) {
	vals := map[string]MigrateRolesAndPermissions{
		"false": MigrateRolesAndPermissionsFalse,
		"true":  MigrateRolesAndPermissionsTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrateRolesAndPermissions(input)
	return &out, nil
}

type MigrationDatabaseState string

const (
	MigrationDatabaseStateCanceled                 MigrationDatabaseState = "Canceled"
	MigrationDatabaseStateCanceling                MigrationDatabaseState = "Canceling"
	MigrationDatabaseStateFailed                   MigrationDatabaseState = "Failed"
	MigrationDatabaseStateInProgress               MigrationDatabaseState = "InProgress"
	MigrationDatabaseStateSucceeded                MigrationDatabaseState = "Succeeded"
	MigrationDatabaseStateWaitingForCutoverTrigger MigrationDatabaseState = "WaitingForCutoverTrigger"
)

func PossibleValuesForMigrationDatabaseState() []string {
	return []string{
		string(MigrationDatabaseStateCanceled),
		string(MigrationDatabaseStateCanceling),
		string(MigrationDatabaseStateFailed),
		string(MigrationDatabaseStateInProgress),
		string(MigrationDatabaseStateSucceeded),
		string(MigrationDatabaseStateWaitingForCutoverTrigger),
	}
}

func (s *MigrationDatabaseState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrationDatabaseState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMigrationDatabaseState(input string) (*MigrationDatabaseState, error) {
	vals := map[string]MigrationDatabaseState{
		"canceled":                 MigrationDatabaseStateCanceled,
		"canceling":                MigrationDatabaseStateCanceling,
		"failed":                   MigrationDatabaseStateFailed,
		"inprogress":               MigrationDatabaseStateInProgress,
		"succeeded":                MigrationDatabaseStateSucceeded,
		"waitingforcutovertrigger": MigrationDatabaseStateWaitingForCutoverTrigger,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationDatabaseState(input)
	return &out, nil
}

type MigrationListFilter string

const (
	MigrationListFilterActive MigrationListFilter = "Active"
	MigrationListFilterAll    MigrationListFilter = "All"
)

func PossibleValuesForMigrationListFilter() []string {
	return []string{
		string(MigrationListFilterActive),
		string(MigrationListFilterAll),
	}
}

func (s *MigrationListFilter) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrationListFilter(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMigrationListFilter(input string) (*MigrationListFilter, error) {
	vals := map[string]MigrationListFilter{
		"active": MigrationListFilterActive,
		"all":    MigrationListFilterAll,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationListFilter(input)
	return &out, nil
}

type MigrationMode string

const (
	MigrationModeOffline MigrationMode = "Offline"
	MigrationModeOnline  MigrationMode = "Online"
)

func PossibleValuesForMigrationMode() []string {
	return []string{
		string(MigrationModeOffline),
		string(MigrationModeOnline),
	}
}

func (s *MigrationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMigrationMode(input string) (*MigrationMode, error) {
	vals := map[string]MigrationMode{
		"offline": MigrationModeOffline,
		"online":  MigrationModeOnline,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationMode(input)
	return &out, nil
}

type MigrationOption string

const (
	MigrationOptionMigrate            MigrationOption = "Migrate"
	MigrationOptionValidate           MigrationOption = "Validate"
	MigrationOptionValidateAndMigrate MigrationOption = "ValidateAndMigrate"
)

func PossibleValuesForMigrationOption() []string {
	return []string{
		string(MigrationOptionMigrate),
		string(MigrationOptionValidate),
		string(MigrationOptionValidateAndMigrate),
	}
}

func (s *MigrationOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrationOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMigrationOption(input string) (*MigrationOption, error) {
	vals := map[string]MigrationOption{
		"migrate":            MigrationOptionMigrate,
		"validate":           MigrationOptionValidate,
		"validateandmigrate": MigrationOptionValidateAndMigrate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationOption(input)
	return &out, nil
}

type MigrationState string

const (
	MigrationStateCanceled             MigrationState = "Canceled"
	MigrationStateCleaningUp           MigrationState = "CleaningUp"
	MigrationStateFailed               MigrationState = "Failed"
	MigrationStateInProgress           MigrationState = "InProgress"
	MigrationStateSucceeded            MigrationState = "Succeeded"
	MigrationStateValidationFailed     MigrationState = "ValidationFailed"
	MigrationStateWaitingForUserAction MigrationState = "WaitingForUserAction"
)

func PossibleValuesForMigrationState() []string {
	return []string{
		string(MigrationStateCanceled),
		string(MigrationStateCleaningUp),
		string(MigrationStateFailed),
		string(MigrationStateInProgress),
		string(MigrationStateSucceeded),
		string(MigrationStateValidationFailed),
		string(MigrationStateWaitingForUserAction),
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
		"canceled":             MigrationStateCanceled,
		"cleaningup":           MigrationStateCleaningUp,
		"failed":               MigrationStateFailed,
		"inprogress":           MigrationStateInProgress,
		"succeeded":            MigrationStateSucceeded,
		"validationfailed":     MigrationStateValidationFailed,
		"waitingforuseraction": MigrationStateWaitingForUserAction,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationState(input)
	return &out, nil
}

type MigrationSubstate string

const (
	MigrationSubstateCancelingRequestedDBMigrations                     MigrationSubstate = "CancelingRequestedDBMigrations"
	MigrationSubstateCompleted                                          MigrationSubstate = "Completed"
	MigrationSubstateCompletingMigration                                MigrationSubstate = "CompletingMigration"
	MigrationSubstateMigratingData                                      MigrationSubstate = "MigratingData"
	MigrationSubstatePerformingPreRequisiteSteps                        MigrationSubstate = "PerformingPreRequisiteSteps"
	MigrationSubstateValidationInProgress                               MigrationSubstate = "ValidationInProgress"
	MigrationSubstateWaitingForCutoverTrigger                           MigrationSubstate = "WaitingForCutoverTrigger"
	MigrationSubstateWaitingForDBsToMigrateSpecification                MigrationSubstate = "WaitingForDBsToMigrateSpecification"
	MigrationSubstateWaitingForDataMigrationScheduling                  MigrationSubstate = "WaitingForDataMigrationScheduling"
	MigrationSubstateWaitingForDataMigrationWindow                      MigrationSubstate = "WaitingForDataMigrationWindow"
	MigrationSubstateWaitingForLogicalReplicationSetupRequestOnSourceDB MigrationSubstate = "WaitingForLogicalReplicationSetupRequestOnSourceDB"
	MigrationSubstateWaitingForTargetDBOverwriteConfirmation            MigrationSubstate = "WaitingForTargetDBOverwriteConfirmation"
)

func PossibleValuesForMigrationSubstate() []string {
	return []string{
		string(MigrationSubstateCancelingRequestedDBMigrations),
		string(MigrationSubstateCompleted),
		string(MigrationSubstateCompletingMigration),
		string(MigrationSubstateMigratingData),
		string(MigrationSubstatePerformingPreRequisiteSteps),
		string(MigrationSubstateValidationInProgress),
		string(MigrationSubstateWaitingForCutoverTrigger),
		string(MigrationSubstateWaitingForDBsToMigrateSpecification),
		string(MigrationSubstateWaitingForDataMigrationScheduling),
		string(MigrationSubstateWaitingForDataMigrationWindow),
		string(MigrationSubstateWaitingForLogicalReplicationSetupRequestOnSourceDB),
		string(MigrationSubstateWaitingForTargetDBOverwriteConfirmation),
	}
}

func (s *MigrationSubstate) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrationSubstate(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMigrationSubstate(input string) (*MigrationSubstate, error) {
	vals := map[string]MigrationSubstate{
		"cancelingrequesteddbmigrations":                     MigrationSubstateCancelingRequestedDBMigrations,
		"completed":                                          MigrationSubstateCompleted,
		"completingmigration":                                MigrationSubstateCompletingMigration,
		"migratingdata":                                      MigrationSubstateMigratingData,
		"performingprerequisitesteps":                        MigrationSubstatePerformingPreRequisiteSteps,
		"validationinprogress":                               MigrationSubstateValidationInProgress,
		"waitingforcutovertrigger":                           MigrationSubstateWaitingForCutoverTrigger,
		"waitingfordbstomigratespecification":                MigrationSubstateWaitingForDBsToMigrateSpecification,
		"waitingfordatamigrationscheduling":                  MigrationSubstateWaitingForDataMigrationScheduling,
		"waitingfordatamigrationwindow":                      MigrationSubstateWaitingForDataMigrationWindow,
		"waitingforlogicalreplicationsetuprequestonsourcedb": MigrationSubstateWaitingForLogicalReplicationSetupRequestOnSourceDB,
		"waitingfortargetdboverwriteconfirmation":            MigrationSubstateWaitingForTargetDBOverwriteConfirmation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationSubstate(input)
	return &out, nil
}

type OverwriteDatabasesOnTargetServer string

const (
	OverwriteDatabasesOnTargetServerFalse OverwriteDatabasesOnTargetServer = "False"
	OverwriteDatabasesOnTargetServerTrue  OverwriteDatabasesOnTargetServer = "True"
)

func PossibleValuesForOverwriteDatabasesOnTargetServer() []string {
	return []string{
		string(OverwriteDatabasesOnTargetServerFalse),
		string(OverwriteDatabasesOnTargetServerTrue),
	}
}

func (s *OverwriteDatabasesOnTargetServer) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOverwriteDatabasesOnTargetServer(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOverwriteDatabasesOnTargetServer(input string) (*OverwriteDatabasesOnTargetServer, error) {
	vals := map[string]OverwriteDatabasesOnTargetServer{
		"false": OverwriteDatabasesOnTargetServerFalse,
		"true":  OverwriteDatabasesOnTargetServerTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OverwriteDatabasesOnTargetServer(input)
	return &out, nil
}

type SkuTier string

const (
	SkuTierBurstable       SkuTier = "Burstable"
	SkuTierGeneralPurpose  SkuTier = "GeneralPurpose"
	SkuTierMemoryOptimized SkuTier = "MemoryOptimized"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBurstable),
		string(SkuTierGeneralPurpose),
		string(SkuTierMemoryOptimized),
	}
}

func (s *SkuTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"burstable":       SkuTierBurstable,
		"generalpurpose":  SkuTierGeneralPurpose,
		"memoryoptimized": SkuTierMemoryOptimized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}

type SourceType string

const (
	SourceTypeAWS                      SourceType = "AWS"
	SourceTypeAWSAURORA                SourceType = "AWS_AURORA"
	SourceTypeAWSECTwo                 SourceType = "AWS_EC2"
	SourceTypeAWSRDS                   SourceType = "AWS_RDS"
	SourceTypeApsaraDBRDS              SourceType = "ApsaraDB_RDS"
	SourceTypeAzureVM                  SourceType = "AzureVM"
	SourceTypeCrunchyPostgreSQL        SourceType = "Crunchy_PostgreSQL"
	SourceTypeDigitalOceanDroplets     SourceType = "Digital_Ocean_Droplets"
	SourceTypeDigitalOceanPostgreSQL   SourceType = "Digital_Ocean_PostgreSQL"
	SourceTypeEDB                      SourceType = "EDB"
	SourceTypeEDBOracleServer          SourceType = "EDB_Oracle_Server"
	SourceTypeEDBPostgreSQL            SourceType = "EDB_PostgreSQL"
	SourceTypeGCP                      SourceType = "GCP"
	SourceTypeGCPAlloyDB               SourceType = "GCP_AlloyDB"
	SourceTypeGCPCloudSQL              SourceType = "GCP_CloudSQL"
	SourceTypeGCPCompute               SourceType = "GCP_Compute"
	SourceTypeHerokuPostgreSQL         SourceType = "Heroku_PostgreSQL"
	SourceTypeHuaweiCompute            SourceType = "Huawei_Compute"
	SourceTypeHuaweiRDS                SourceType = "Huawei_RDS"
	SourceTypeOnPremises               SourceType = "OnPremises"
	SourceTypePostgreSQLCosmosDB       SourceType = "PostgreSQLCosmosDB"
	SourceTypePostgreSQLFlexibleServer SourceType = "PostgreSQLFlexibleServer"
	SourceTypePostgreSQLSingleServer   SourceType = "PostgreSQLSingleServer"
	SourceTypeSupabasePostgreSQL       SourceType = "Supabase_PostgreSQL"
)

func PossibleValuesForSourceType() []string {
	return []string{
		string(SourceTypeAWS),
		string(SourceTypeAWSAURORA),
		string(SourceTypeAWSECTwo),
		string(SourceTypeAWSRDS),
		string(SourceTypeApsaraDBRDS),
		string(SourceTypeAzureVM),
		string(SourceTypeCrunchyPostgreSQL),
		string(SourceTypeDigitalOceanDroplets),
		string(SourceTypeDigitalOceanPostgreSQL),
		string(SourceTypeEDB),
		string(SourceTypeEDBOracleServer),
		string(SourceTypeEDBPostgreSQL),
		string(SourceTypeGCP),
		string(SourceTypeGCPAlloyDB),
		string(SourceTypeGCPCloudSQL),
		string(SourceTypeGCPCompute),
		string(SourceTypeHerokuPostgreSQL),
		string(SourceTypeHuaweiCompute),
		string(SourceTypeHuaweiRDS),
		string(SourceTypeOnPremises),
		string(SourceTypePostgreSQLCosmosDB),
		string(SourceTypePostgreSQLFlexibleServer),
		string(SourceTypePostgreSQLSingleServer),
		string(SourceTypeSupabasePostgreSQL),
	}
}

func (s *SourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSourceType(input string) (*SourceType, error) {
	vals := map[string]SourceType{
		"aws":                      SourceTypeAWS,
		"aws_aurora":               SourceTypeAWSAURORA,
		"aws_ec2":                  SourceTypeAWSECTwo,
		"aws_rds":                  SourceTypeAWSRDS,
		"apsaradb_rds":             SourceTypeApsaraDBRDS,
		"azurevm":                  SourceTypeAzureVM,
		"crunchy_postgresql":       SourceTypeCrunchyPostgreSQL,
		"digital_ocean_droplets":   SourceTypeDigitalOceanDroplets,
		"digital_ocean_postgresql": SourceTypeDigitalOceanPostgreSQL,
		"edb":                      SourceTypeEDB,
		"edb_oracle_server":        SourceTypeEDBOracleServer,
		"edb_postgresql":           SourceTypeEDBPostgreSQL,
		"gcp":                      SourceTypeGCP,
		"gcp_alloydb":              SourceTypeGCPAlloyDB,
		"gcp_cloudsql":             SourceTypeGCPCloudSQL,
		"gcp_compute":              SourceTypeGCPCompute,
		"heroku_postgresql":        SourceTypeHerokuPostgreSQL,
		"huawei_compute":           SourceTypeHuaweiCompute,
		"huawei_rds":               SourceTypeHuaweiRDS,
		"onpremises":               SourceTypeOnPremises,
		"postgresqlcosmosdb":       SourceTypePostgreSQLCosmosDB,
		"postgresqlflexibleserver": SourceTypePostgreSQLFlexibleServer,
		"postgresqlsingleserver":   SourceTypePostgreSQLSingleServer,
		"supabase_postgresql":      SourceTypeSupabasePostgreSQL,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SourceType(input)
	return &out, nil
}

type SslMode string

const (
	SslModePrefer     SslMode = "Prefer"
	SslModeRequire    SslMode = "Require"
	SslModeVerifyCA   SslMode = "VerifyCA"
	SslModeVerifyFull SslMode = "VerifyFull"
)

func PossibleValuesForSslMode() []string {
	return []string{
		string(SslModePrefer),
		string(SslModeRequire),
		string(SslModeVerifyCA),
		string(SslModeVerifyFull),
	}
}

func (s *SslMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSslMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSslMode(input string) (*SslMode, error) {
	vals := map[string]SslMode{
		"prefer":     SslModePrefer,
		"require":    SslModeRequire,
		"verifyca":   SslModeVerifyCA,
		"verifyfull": SslModeVerifyFull,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SslMode(input)
	return &out, nil
}

type StartDataMigration string

const (
	StartDataMigrationFalse StartDataMigration = "False"
	StartDataMigrationTrue  StartDataMigration = "True"
)

func PossibleValuesForStartDataMigration() []string {
	return []string{
		string(StartDataMigrationFalse),
		string(StartDataMigrationTrue),
	}
}

func (s *StartDataMigration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStartDataMigration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStartDataMigration(input string) (*StartDataMigration, error) {
	vals := map[string]StartDataMigration{
		"false": StartDataMigrationFalse,
		"true":  StartDataMigrationTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StartDataMigration(input)
	return &out, nil
}

type TriggerCutover string

const (
	TriggerCutoverFalse TriggerCutover = "False"
	TriggerCutoverTrue  TriggerCutover = "True"
)

func PossibleValuesForTriggerCutover() []string {
	return []string{
		string(TriggerCutoverFalse),
		string(TriggerCutoverTrue),
	}
}

func (s *TriggerCutover) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggerCutover(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggerCutover(input string) (*TriggerCutover, error) {
	vals := map[string]TriggerCutover{
		"false": TriggerCutoverFalse,
		"true":  TriggerCutoverTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerCutover(input)
	return &out, nil
}

type ValidationState string

const (
	ValidationStateFailed    ValidationState = "Failed"
	ValidationStateSucceeded ValidationState = "Succeeded"
	ValidationStateWarning   ValidationState = "Warning"
)

func PossibleValuesForValidationState() []string {
	return []string{
		string(ValidationStateFailed),
		string(ValidationStateSucceeded),
		string(ValidationStateWarning),
	}
}

func (s *ValidationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseValidationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseValidationState(input string) (*ValidationState, error) {
	vals := map[string]ValidationState{
		"failed":    ValidationStateFailed,
		"succeeded": ValidationStateSucceeded,
		"warning":   ValidationStateWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ValidationState(input)
	return &out, nil
}
