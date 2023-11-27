package databaseoperations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementOperationState string

const (
	ManagementOperationStateCancelInProgress ManagementOperationState = "CancelInProgress"
	ManagementOperationStateCancelled        ManagementOperationState = "Cancelled"
	ManagementOperationStateFailed           ManagementOperationState = "Failed"
	ManagementOperationStateInProgress       ManagementOperationState = "InProgress"
	ManagementOperationStatePending          ManagementOperationState = "Pending"
	ManagementOperationStateSucceeded        ManagementOperationState = "Succeeded"
)

func PossibleValuesForManagementOperationState() []string {
	return []string{
		string(ManagementOperationStateCancelInProgress),
		string(ManagementOperationStateCancelled),
		string(ManagementOperationStateFailed),
		string(ManagementOperationStateInProgress),
		string(ManagementOperationStatePending),
		string(ManagementOperationStateSucceeded),
	}
}

func (s *ManagementOperationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagementOperationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagementOperationState(input string) (*ManagementOperationState, error) {
	vals := map[string]ManagementOperationState{
		"cancelinprogress": ManagementOperationStateCancelInProgress,
		"cancelled":        ManagementOperationStateCancelled,
		"failed":           ManagementOperationStateFailed,
		"inprogress":       ManagementOperationStateInProgress,
		"pending":          ManagementOperationStatePending,
		"succeeded":        ManagementOperationStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagementOperationState(input)
	return &out, nil
}
