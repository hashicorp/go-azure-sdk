package migrations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CancelEnum string

const (
	CancelEnumFalse CancelEnum = "False"
	CancelEnumTrue  CancelEnum = "True"
)

func PossibleValuesForCancelEnum() []string {
	return []string{
		string(CancelEnumFalse),
		string(CancelEnumTrue),
	}
}

func (s *CancelEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCancelEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCancelEnum(input string) (*CancelEnum, error) {
	vals := map[string]CancelEnum{
		"false": CancelEnumFalse,
		"true":  CancelEnumTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CancelEnum(input)
	return &out, nil
}

type LogicalReplicationOnSourceDbEnum string

const (
	LogicalReplicationOnSourceDbEnumFalse LogicalReplicationOnSourceDbEnum = "False"
	LogicalReplicationOnSourceDbEnumTrue  LogicalReplicationOnSourceDbEnum = "True"
)

func PossibleValuesForLogicalReplicationOnSourceDbEnum() []string {
	return []string{
		string(LogicalReplicationOnSourceDbEnumFalse),
		string(LogicalReplicationOnSourceDbEnumTrue),
	}
}

func (s *LogicalReplicationOnSourceDbEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLogicalReplicationOnSourceDbEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLogicalReplicationOnSourceDbEnum(input string) (*LogicalReplicationOnSourceDbEnum, error) {
	vals := map[string]LogicalReplicationOnSourceDbEnum{
		"false": LogicalReplicationOnSourceDbEnumFalse,
		"true":  LogicalReplicationOnSourceDbEnumTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogicalReplicationOnSourceDbEnum(input)
	return &out, nil
}

type MigrateRolesEnum string

const (
	MigrateRolesEnumFalse MigrateRolesEnum = "False"
	MigrateRolesEnumTrue  MigrateRolesEnum = "True"
)

func PossibleValuesForMigrateRolesEnum() []string {
	return []string{
		string(MigrateRolesEnumFalse),
		string(MigrateRolesEnumTrue),
	}
}

func (s *MigrateRolesEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrateRolesEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMigrateRolesEnum(input string) (*MigrateRolesEnum, error) {
	vals := map[string]MigrateRolesEnum{
		"false": MigrateRolesEnumFalse,
		"true":  MigrateRolesEnumTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrateRolesEnum(input)
	return &out, nil
}

type MigrationDbState string

const (
	MigrationDbStateCanceled                 MigrationDbState = "Canceled"
	MigrationDbStateCanceling                MigrationDbState = "Canceling"
	MigrationDbStateFailed                   MigrationDbState = "Failed"
	MigrationDbStateInProgress               MigrationDbState = "InProgress"
	MigrationDbStateSucceeded                MigrationDbState = "Succeeded"
	MigrationDbStateWaitingForCutoverTrigger MigrationDbState = "WaitingForCutoverTrigger"
)

func PossibleValuesForMigrationDbState() []string {
	return []string{
		string(MigrationDbStateCanceled),
		string(MigrationDbStateCanceling),
		string(MigrationDbStateFailed),
		string(MigrationDbStateInProgress),
		string(MigrationDbStateSucceeded),
		string(MigrationDbStateWaitingForCutoverTrigger),
	}
}

func (s *MigrationDbState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrationDbState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMigrationDbState(input string) (*MigrationDbState, error) {
	vals := map[string]MigrationDbState{
		"canceled":                 MigrationDbStateCanceled,
		"canceling":                MigrationDbStateCanceling,
		"failed":                   MigrationDbStateFailed,
		"inprogress":               MigrationDbStateInProgress,
		"succeeded":                MigrationDbStateSucceeded,
		"waitingforcutovertrigger": MigrationDbStateWaitingForCutoverTrigger,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationDbState(input)
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

type MigrationSubState string

const (
	MigrationSubStateCancelingRequestedDBMigrations                     MigrationSubState = "CancelingRequestedDBMigrations"
	MigrationSubStateCompleted                                          MigrationSubState = "Completed"
	MigrationSubStateCompletingMigration                                MigrationSubState = "CompletingMigration"
	MigrationSubStateMigratingData                                      MigrationSubState = "MigratingData"
	MigrationSubStatePerformingPreRequisiteSteps                        MigrationSubState = "PerformingPreRequisiteSteps"
	MigrationSubStateValidationInProgress                               MigrationSubState = "ValidationInProgress"
	MigrationSubStateWaitingForCutoverTrigger                           MigrationSubState = "WaitingForCutoverTrigger"
	MigrationSubStateWaitingForDBsToMigrateSpecification                MigrationSubState = "WaitingForDBsToMigrateSpecification"
	MigrationSubStateWaitingForDataMigrationScheduling                  MigrationSubState = "WaitingForDataMigrationScheduling"
	MigrationSubStateWaitingForDataMigrationWindow                      MigrationSubState = "WaitingForDataMigrationWindow"
	MigrationSubStateWaitingForLogicalReplicationSetupRequestOnSourceDB MigrationSubState = "WaitingForLogicalReplicationSetupRequestOnSourceDB"
	MigrationSubStateWaitingForTargetDBOverwriteConfirmation            MigrationSubState = "WaitingForTargetDBOverwriteConfirmation"
)

func PossibleValuesForMigrationSubState() []string {
	return []string{
		string(MigrationSubStateCancelingRequestedDBMigrations),
		string(MigrationSubStateCompleted),
		string(MigrationSubStateCompletingMigration),
		string(MigrationSubStateMigratingData),
		string(MigrationSubStatePerformingPreRequisiteSteps),
		string(MigrationSubStateValidationInProgress),
		string(MigrationSubStateWaitingForCutoverTrigger),
		string(MigrationSubStateWaitingForDBsToMigrateSpecification),
		string(MigrationSubStateWaitingForDataMigrationScheduling),
		string(MigrationSubStateWaitingForDataMigrationWindow),
		string(MigrationSubStateWaitingForLogicalReplicationSetupRequestOnSourceDB),
		string(MigrationSubStateWaitingForTargetDBOverwriteConfirmation),
	}
}

func (s *MigrationSubState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMigrationSubState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMigrationSubState(input string) (*MigrationSubState, error) {
	vals := map[string]MigrationSubState{
		"cancelingrequesteddbmigrations":                     MigrationSubStateCancelingRequestedDBMigrations,
		"completed":                                          MigrationSubStateCompleted,
		"completingmigration":                                MigrationSubStateCompletingMigration,
		"migratingdata":                                      MigrationSubStateMigratingData,
		"performingprerequisitesteps":                        MigrationSubStatePerformingPreRequisiteSteps,
		"validationinprogress":                               MigrationSubStateValidationInProgress,
		"waitingforcutovertrigger":                           MigrationSubStateWaitingForCutoverTrigger,
		"waitingfordbstomigratespecification":                MigrationSubStateWaitingForDBsToMigrateSpecification,
		"waitingfordatamigrationscheduling":                  MigrationSubStateWaitingForDataMigrationScheduling,
		"waitingfordatamigrationwindow":                      MigrationSubStateWaitingForDataMigrationWindow,
		"waitingforlogicalreplicationsetuprequestonsourcedb": MigrationSubStateWaitingForLogicalReplicationSetupRequestOnSourceDB,
		"waitingfortargetdboverwriteconfirmation":            MigrationSubStateWaitingForTargetDBOverwriteConfirmation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationSubState(input)
	return &out, nil
}

type OverwriteDbsInTargetEnum string

const (
	OverwriteDbsInTargetEnumFalse OverwriteDbsInTargetEnum = "False"
	OverwriteDbsInTargetEnumTrue  OverwriteDbsInTargetEnum = "True"
)

func PossibleValuesForOverwriteDbsInTargetEnum() []string {
	return []string{
		string(OverwriteDbsInTargetEnumFalse),
		string(OverwriteDbsInTargetEnumTrue),
	}
}

func (s *OverwriteDbsInTargetEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOverwriteDbsInTargetEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOverwriteDbsInTargetEnum(input string) (*OverwriteDbsInTargetEnum, error) {
	vals := map[string]OverwriteDbsInTargetEnum{
		"false": OverwriteDbsInTargetEnumFalse,
		"true":  OverwriteDbsInTargetEnumTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OverwriteDbsInTargetEnum(input)
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
	SourceTypeAWS                    SourceType = "AWS"
	SourceTypeAWSAURORA              SourceType = "AWS_AURORA"
	SourceTypeAWSECTwo               SourceType = "AWS_EC2"
	SourceTypeAWSRDS                 SourceType = "AWS_RDS"
	SourceTypeAzureVM                SourceType = "AzureVM"
	SourceTypeEDB                    SourceType = "EDB"
	SourceTypeGCP                    SourceType = "GCP"
	SourceTypeGCPAlloyDB             SourceType = "GCP_AlloyDB"
	SourceTypeGCPCloudSQL            SourceType = "GCP_CloudSQL"
	SourceTypeGCPCompute             SourceType = "GCP_Compute"
	SourceTypeOnPremises             SourceType = "OnPremises"
	SourceTypePostgreSQLSingleServer SourceType = "PostgreSQLSingleServer"
)

func PossibleValuesForSourceType() []string {
	return []string{
		string(SourceTypeAWS),
		string(SourceTypeAWSAURORA),
		string(SourceTypeAWSECTwo),
		string(SourceTypeAWSRDS),
		string(SourceTypeAzureVM),
		string(SourceTypeEDB),
		string(SourceTypeGCP),
		string(SourceTypeGCPAlloyDB),
		string(SourceTypeGCPCloudSQL),
		string(SourceTypeGCPCompute),
		string(SourceTypeOnPremises),
		string(SourceTypePostgreSQLSingleServer),
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
		"aws":                    SourceTypeAWS,
		"aws_aurora":             SourceTypeAWSAURORA,
		"aws_ec2":                SourceTypeAWSECTwo,
		"aws_rds":                SourceTypeAWSRDS,
		"azurevm":                SourceTypeAzureVM,
		"edb":                    SourceTypeEDB,
		"gcp":                    SourceTypeGCP,
		"gcp_alloydb":            SourceTypeGCPAlloyDB,
		"gcp_cloudsql":           SourceTypeGCPCloudSQL,
		"gcp_compute":            SourceTypeGCPCompute,
		"onpremises":             SourceTypeOnPremises,
		"postgresqlsingleserver": SourceTypePostgreSQLSingleServer,
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

type StartDataMigrationEnum string

const (
	StartDataMigrationEnumFalse StartDataMigrationEnum = "False"
	StartDataMigrationEnumTrue  StartDataMigrationEnum = "True"
)

func PossibleValuesForStartDataMigrationEnum() []string {
	return []string{
		string(StartDataMigrationEnumFalse),
		string(StartDataMigrationEnumTrue),
	}
}

func (s *StartDataMigrationEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStartDataMigrationEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStartDataMigrationEnum(input string) (*StartDataMigrationEnum, error) {
	vals := map[string]StartDataMigrationEnum{
		"false": StartDataMigrationEnumFalse,
		"true":  StartDataMigrationEnumTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StartDataMigrationEnum(input)
	return &out, nil
}

type TriggerCutoverEnum string

const (
	TriggerCutoverEnumFalse TriggerCutoverEnum = "False"
	TriggerCutoverEnumTrue  TriggerCutoverEnum = "True"
)

func PossibleValuesForTriggerCutoverEnum() []string {
	return []string{
		string(TriggerCutoverEnumFalse),
		string(TriggerCutoverEnumTrue),
	}
}

func (s *TriggerCutoverEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggerCutoverEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggerCutoverEnum(input string) (*TriggerCutoverEnum, error) {
	vals := map[string]TriggerCutoverEnum{
		"false": TriggerCutoverEnumFalse,
		"true":  TriggerCutoverEnumTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerCutoverEnum(input)
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
