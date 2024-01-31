package migratepostgresqlazuredbforpostgresqlsynctasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
