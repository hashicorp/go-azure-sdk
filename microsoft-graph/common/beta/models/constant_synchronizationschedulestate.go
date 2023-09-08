package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationScheduleState string

const (
	SynchronizationScheduleStateActive   SynchronizationScheduleState = "Active"
	SynchronizationScheduleStateDisabled SynchronizationScheduleState = "Disabled"
	SynchronizationScheduleStatePaused   SynchronizationScheduleState = "Paused"
)

func PossibleValuesForSynchronizationScheduleState() []string {
	return []string{
		string(SynchronizationScheduleStateActive),
		string(SynchronizationScheduleStateDisabled),
		string(SynchronizationScheduleStatePaused),
	}
}

func parseSynchronizationScheduleState(input string) (*SynchronizationScheduleState, error) {
	vals := map[string]SynchronizationScheduleState{
		"active":   SynchronizationScheduleStateActive,
		"disabled": SynchronizationScheduleStateDisabled,
		"paused":   SynchronizationScheduleStatePaused,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationScheduleState(input)
	return &out, nil
}
