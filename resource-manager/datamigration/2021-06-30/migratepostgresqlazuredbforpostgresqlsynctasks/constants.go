package migratepostgresqlazuredbforpostgresqlsynctasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicateMigrationState string

const (
	ReplicateMigrationStateACTIONREQUIRED ReplicateMigrationState = "ACTION_REQUIRED"
	ReplicateMigrationStateCOMPLETE       ReplicateMigrationState = "COMPLETE"
	ReplicateMigrationStateFAILED         ReplicateMigrationState = "FAILED"
	ReplicateMigrationStatePENDING        ReplicateMigrationState = "PENDING"
	ReplicateMigrationStateUNDEFINED      ReplicateMigrationState = "UNDEFINED"
	ReplicateMigrationStateVALIDATING     ReplicateMigrationState = "VALIDATING"
)

func PossibleValuesForReplicateMigrationState() []string {
	return []string{
		string(ReplicateMigrationStateACTIONREQUIRED),
		string(ReplicateMigrationStateCOMPLETE),
		string(ReplicateMigrationStateFAILED),
		string(ReplicateMigrationStatePENDING),
		string(ReplicateMigrationStateUNDEFINED),
		string(ReplicateMigrationStateVALIDATING),
	}
}

func (s *ReplicateMigrationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReplicateMigrationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReplicateMigrationState(input string) (*ReplicateMigrationState, error) {
	vals := map[string]ReplicateMigrationState{
		"action_required": ReplicateMigrationStateACTIONREQUIRED,
		"complete":        ReplicateMigrationStateCOMPLETE,
		"failed":          ReplicateMigrationStateFAILED,
		"pending":         ReplicateMigrationStatePENDING,
		"undefined":       ReplicateMigrationStateUNDEFINED,
		"validating":      ReplicateMigrationStateVALIDATING,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicateMigrationState(input)
	return &out, nil
}

type ScenarioSource string

const (
	ScenarioSourceAccess        ScenarioSource = "Access"
	ScenarioSourceDBTwo         ScenarioSource = "DB2"
	ScenarioSourceMongoDB       ScenarioSource = "MongoDB"
	ScenarioSourceMySQL         ScenarioSource = "MySQL"
	ScenarioSourceMySQLRDS      ScenarioSource = "MySQLRDS"
	ScenarioSourceOracle        ScenarioSource = "Oracle"
	ScenarioSourcePostgreSQL    ScenarioSource = "PostgreSQL"
	ScenarioSourcePostgreSQLRDS ScenarioSource = "PostgreSQLRDS"
	ScenarioSourceSQL           ScenarioSource = "SQL"
	ScenarioSourceSQLRDS        ScenarioSource = "SQLRDS"
	ScenarioSourceSybase        ScenarioSource = "Sybase"
)

func PossibleValuesForScenarioSource() []string {
	return []string{
		string(ScenarioSourceAccess),
		string(ScenarioSourceDBTwo),
		string(ScenarioSourceMongoDB),
		string(ScenarioSourceMySQL),
		string(ScenarioSourceMySQLRDS),
		string(ScenarioSourceOracle),
		string(ScenarioSourcePostgreSQL),
		string(ScenarioSourcePostgreSQLRDS),
		string(ScenarioSourceSQL),
		string(ScenarioSourceSQLRDS),
		string(ScenarioSourceSybase),
	}
}

func (s *ScenarioSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScenarioSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScenarioSource(input string) (*ScenarioSource, error) {
	vals := map[string]ScenarioSource{
		"access":        ScenarioSourceAccess,
		"db2":           ScenarioSourceDBTwo,
		"mongodb":       ScenarioSourceMongoDB,
		"mysql":         ScenarioSourceMySQL,
		"mysqlrds":      ScenarioSourceMySQLRDS,
		"oracle":        ScenarioSourceOracle,
		"postgresql":    ScenarioSourcePostgreSQL,
		"postgresqlrds": ScenarioSourcePostgreSQLRDS,
		"sql":           ScenarioSourceSQL,
		"sqlrds":        ScenarioSourceSQLRDS,
		"sybase":        ScenarioSourceSybase,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScenarioSource(input)
	return &out, nil
}

type ScenarioTarget string

const (
	ScenarioTargetAzureDBForMySql       ScenarioTarget = "AzureDBForMySql"
	ScenarioTargetAzureDBForPostgresSQL ScenarioTarget = "AzureDBForPostgresSQL"
	ScenarioTargetMongoDB               ScenarioTarget = "MongoDB"
	ScenarioTargetSQLDB                 ScenarioTarget = "SQLDB"
	ScenarioTargetSQLDW                 ScenarioTarget = "SQLDW"
	ScenarioTargetSQLMI                 ScenarioTarget = "SQLMI"
	ScenarioTargetSQLServer             ScenarioTarget = "SQLServer"
)

func PossibleValuesForScenarioTarget() []string {
	return []string{
		string(ScenarioTargetAzureDBForMySql),
		string(ScenarioTargetAzureDBForPostgresSQL),
		string(ScenarioTargetMongoDB),
		string(ScenarioTargetSQLDB),
		string(ScenarioTargetSQLDW),
		string(ScenarioTargetSQLMI),
		string(ScenarioTargetSQLServer),
	}
}

func (s *ScenarioTarget) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScenarioTarget(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScenarioTarget(input string) (*ScenarioTarget, error) {
	vals := map[string]ScenarioTarget{
		"azuredbformysql":       ScenarioTargetAzureDBForMySql,
		"azuredbforpostgressql": ScenarioTargetAzureDBForPostgresSQL,
		"mongodb":               ScenarioTargetMongoDB,
		"sqldb":                 ScenarioTargetSQLDB,
		"sqldw":                 ScenarioTargetSQLDW,
		"sqlmi":                 ScenarioTargetSQLMI,
		"sqlserver":             ScenarioTargetSQLServer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScenarioTarget(input)
	return &out, nil
}

type SyncDatabaseMigrationReportingState string

const (
	SyncDatabaseMigrationReportingStateBACKUPCOMPLETED    SyncDatabaseMigrationReportingState = "BACKUP_COMPLETED"
	SyncDatabaseMigrationReportingStateBACKUPINPROGRESS   SyncDatabaseMigrationReportingState = "BACKUP_IN_PROGRESS"
	SyncDatabaseMigrationReportingStateCANCELLED          SyncDatabaseMigrationReportingState = "CANCELLED"
	SyncDatabaseMigrationReportingStateCANCELLING         SyncDatabaseMigrationReportingState = "CANCELLING"
	SyncDatabaseMigrationReportingStateCOMPLETE           SyncDatabaseMigrationReportingState = "COMPLETE"
	SyncDatabaseMigrationReportingStateCOMPLETING         SyncDatabaseMigrationReportingState = "COMPLETING"
	SyncDatabaseMigrationReportingStateCONFIGURING        SyncDatabaseMigrationReportingState = "CONFIGURING"
	SyncDatabaseMigrationReportingStateFAILED             SyncDatabaseMigrationReportingState = "FAILED"
	SyncDatabaseMigrationReportingStateINITIALIAZING      SyncDatabaseMigrationReportingState = "INITIALIAZING"
	SyncDatabaseMigrationReportingStateREADYTOCOMPLETE    SyncDatabaseMigrationReportingState = "READY_TO_COMPLETE"
	SyncDatabaseMigrationReportingStateRESTORECOMPLETED   SyncDatabaseMigrationReportingState = "RESTORE_COMPLETED"
	SyncDatabaseMigrationReportingStateRESTOREINPROGRESS  SyncDatabaseMigrationReportingState = "RESTORE_IN_PROGRESS"
	SyncDatabaseMigrationReportingStateRUNNING            SyncDatabaseMigrationReportingState = "RUNNING"
	SyncDatabaseMigrationReportingStateSTARTING           SyncDatabaseMigrationReportingState = "STARTING"
	SyncDatabaseMigrationReportingStateUNDEFINED          SyncDatabaseMigrationReportingState = "UNDEFINED"
	SyncDatabaseMigrationReportingStateVALIDATING         SyncDatabaseMigrationReportingState = "VALIDATING"
	SyncDatabaseMigrationReportingStateVALIDATIONCOMPLETE SyncDatabaseMigrationReportingState = "VALIDATION_COMPLETE"
	SyncDatabaseMigrationReportingStateVALIDATIONFAILED   SyncDatabaseMigrationReportingState = "VALIDATION_FAILED"
)

func PossibleValuesForSyncDatabaseMigrationReportingState() []string {
	return []string{
		string(SyncDatabaseMigrationReportingStateBACKUPCOMPLETED),
		string(SyncDatabaseMigrationReportingStateBACKUPINPROGRESS),
		string(SyncDatabaseMigrationReportingStateCANCELLED),
		string(SyncDatabaseMigrationReportingStateCANCELLING),
		string(SyncDatabaseMigrationReportingStateCOMPLETE),
		string(SyncDatabaseMigrationReportingStateCOMPLETING),
		string(SyncDatabaseMigrationReportingStateCONFIGURING),
		string(SyncDatabaseMigrationReportingStateFAILED),
		string(SyncDatabaseMigrationReportingStateINITIALIAZING),
		string(SyncDatabaseMigrationReportingStateREADYTOCOMPLETE),
		string(SyncDatabaseMigrationReportingStateRESTORECOMPLETED),
		string(SyncDatabaseMigrationReportingStateRESTOREINPROGRESS),
		string(SyncDatabaseMigrationReportingStateRUNNING),
		string(SyncDatabaseMigrationReportingStateSTARTING),
		string(SyncDatabaseMigrationReportingStateUNDEFINED),
		string(SyncDatabaseMigrationReportingStateVALIDATING),
		string(SyncDatabaseMigrationReportingStateVALIDATIONCOMPLETE),
		string(SyncDatabaseMigrationReportingStateVALIDATIONFAILED),
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
		"backup_completed":    SyncDatabaseMigrationReportingStateBACKUPCOMPLETED,
		"backup_in_progress":  SyncDatabaseMigrationReportingStateBACKUPINPROGRESS,
		"cancelled":           SyncDatabaseMigrationReportingStateCANCELLED,
		"cancelling":          SyncDatabaseMigrationReportingStateCANCELLING,
		"complete":            SyncDatabaseMigrationReportingStateCOMPLETE,
		"completing":          SyncDatabaseMigrationReportingStateCOMPLETING,
		"configuring":         SyncDatabaseMigrationReportingStateCONFIGURING,
		"failed":              SyncDatabaseMigrationReportingStateFAILED,
		"initialiazing":       SyncDatabaseMigrationReportingStateINITIALIAZING,
		"ready_to_complete":   SyncDatabaseMigrationReportingStateREADYTOCOMPLETE,
		"restore_completed":   SyncDatabaseMigrationReportingStateRESTORECOMPLETED,
		"restore_in_progress": SyncDatabaseMigrationReportingStateRESTOREINPROGRESS,
		"running":             SyncDatabaseMigrationReportingStateRUNNING,
		"starting":            SyncDatabaseMigrationReportingStateSTARTING,
		"undefined":           SyncDatabaseMigrationReportingStateUNDEFINED,
		"validating":          SyncDatabaseMigrationReportingStateVALIDATING,
		"validation_complete": SyncDatabaseMigrationReportingStateVALIDATIONCOMPLETE,
		"validation_failed":   SyncDatabaseMigrationReportingStateVALIDATIONFAILED,
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
