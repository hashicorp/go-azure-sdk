package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationTaskExecutionState string

const (
	SynchronizationTaskExecutionState_EntryLevelErrors SynchronizationTaskExecutionState = "EntryLevelErrors"
	SynchronizationTaskExecutionState_Failed           SynchronizationTaskExecutionState = "Failed"
	SynchronizationTaskExecutionState_Succeeded        SynchronizationTaskExecutionState = "Succeeded"
)

func PossibleValuesForSynchronizationTaskExecutionState() []string {
	return []string{
		string(SynchronizationTaskExecutionState_EntryLevelErrors),
		string(SynchronizationTaskExecutionState_Failed),
		string(SynchronizationTaskExecutionState_Succeeded),
	}
}

func (s *SynchronizationTaskExecutionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationTaskExecutionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationTaskExecutionState(input string) (*SynchronizationTaskExecutionState, error) {
	vals := map[string]SynchronizationTaskExecutionState{
		"entrylevelerrors": SynchronizationTaskExecutionState_EntryLevelErrors,
		"failed":           SynchronizationTaskExecutionState_Failed,
		"succeeded":        SynchronizationTaskExecutionState_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationTaskExecutionState(input)
	return &out, nil
}
