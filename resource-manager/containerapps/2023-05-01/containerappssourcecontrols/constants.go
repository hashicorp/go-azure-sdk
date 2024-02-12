package containerappssourcecontrols

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlOperationState string

const (
	SourceControlOperationStateCanceled   SourceControlOperationState = "Canceled"
	SourceControlOperationStateFailed     SourceControlOperationState = "Failed"
	SourceControlOperationStateInProgress SourceControlOperationState = "InProgress"
	SourceControlOperationStateSucceeded  SourceControlOperationState = "Succeeded"
)

func PossibleValuesForSourceControlOperationState() []string {
	return []string{
		string(SourceControlOperationStateCanceled),
		string(SourceControlOperationStateFailed),
		string(SourceControlOperationStateInProgress),
		string(SourceControlOperationStateSucceeded),
	}
}

func (s *SourceControlOperationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSourceControlOperationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSourceControlOperationState(input string) (*SourceControlOperationState, error) {
	vals := map[string]SourceControlOperationState{
		"canceled":   SourceControlOperationStateCanceled,
		"failed":     SourceControlOperationStateFailed,
		"inprogress": SourceControlOperationStateInProgress,
		"succeeded":  SourceControlOperationStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SourceControlOperationState(input)
	return &out, nil
}
