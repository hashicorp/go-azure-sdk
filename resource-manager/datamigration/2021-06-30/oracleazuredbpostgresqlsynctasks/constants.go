package oracleazuredbpostgresqlsynctasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
