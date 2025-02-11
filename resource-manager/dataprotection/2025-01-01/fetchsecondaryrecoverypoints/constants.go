package fetchsecondaryrecoverypoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointCompletionState string

const (
	RecoveryPointCompletionStateCompleted RecoveryPointCompletionState = "Completed"
	RecoveryPointCompletionStatePartial   RecoveryPointCompletionState = "Partial"
)

func PossibleValuesForRecoveryPointCompletionState() []string {
	return []string{
		string(RecoveryPointCompletionStateCompleted),
		string(RecoveryPointCompletionStatePartial),
	}
}

func (s *RecoveryPointCompletionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecoveryPointCompletionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecoveryPointCompletionState(input string) (*RecoveryPointCompletionState, error) {
	vals := map[string]RecoveryPointCompletionState{
		"completed": RecoveryPointCompletionStateCompleted,
		"partial":   RecoveryPointCompletionStatePartial,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecoveryPointCompletionState(input)
	return &out, nil
}

type RehydrationStatus string

const (
	RehydrationStatusCOMPLETED        RehydrationStatus = "COMPLETED"
	RehydrationStatusCREATEINPROGRESS RehydrationStatus = "CREATE_IN_PROGRESS"
	RehydrationStatusDELETED          RehydrationStatus = "DELETED"
	RehydrationStatusDELETEINPROGRESS RehydrationStatus = "DELETE_IN_PROGRESS"
	RehydrationStatusFAILED           RehydrationStatus = "FAILED"
)

func PossibleValuesForRehydrationStatus() []string {
	return []string{
		string(RehydrationStatusCOMPLETED),
		string(RehydrationStatusCREATEINPROGRESS),
		string(RehydrationStatusDELETED),
		string(RehydrationStatusDELETEINPROGRESS),
		string(RehydrationStatusFAILED),
	}
}

func (s *RehydrationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRehydrationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRehydrationStatus(input string) (*RehydrationStatus, error) {
	vals := map[string]RehydrationStatus{
		"completed":          RehydrationStatusCOMPLETED,
		"create_in_progress": RehydrationStatusCREATEINPROGRESS,
		"deleted":            RehydrationStatusDELETED,
		"delete_in_progress": RehydrationStatusDELETEINPROGRESS,
		"failed":             RehydrationStatusFAILED,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RehydrationStatus(input)
	return &out, nil
}
