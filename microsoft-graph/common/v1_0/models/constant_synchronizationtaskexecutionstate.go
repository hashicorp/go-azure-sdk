package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationTaskExecutionState string

const (
	SynchronizationTaskExecutionStateEntryLevelErrors SynchronizationTaskExecutionState = "EntryLevelErrors"
	SynchronizationTaskExecutionStateFailed           SynchronizationTaskExecutionState = "Failed"
	SynchronizationTaskExecutionStateSucceeded        SynchronizationTaskExecutionState = "Succeeded"
)

func PossibleValuesForSynchronizationTaskExecutionState() []string {
	return []string{
		string(SynchronizationTaskExecutionStateEntryLevelErrors),
		string(SynchronizationTaskExecutionStateFailed),
		string(SynchronizationTaskExecutionStateSucceeded),
	}
}

func parseSynchronizationTaskExecutionState(input string) (*SynchronizationTaskExecutionState, error) {
	vals := map[string]SynchronizationTaskExecutionState{
		"entrylevelerrors": SynchronizationTaskExecutionStateEntryLevelErrors,
		"failed":           SynchronizationTaskExecutionStateFailed,
		"succeeded":        SynchronizationTaskExecutionStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationTaskExecutionState(input)
	return &out, nil
}
