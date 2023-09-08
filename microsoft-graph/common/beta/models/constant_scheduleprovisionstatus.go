package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleProvisionStatus string

const (
	ScheduleProvisionStatusCompleted  ScheduleProvisionStatus = "Completed"
	ScheduleProvisionStatusFailed     ScheduleProvisionStatus = "Failed"
	ScheduleProvisionStatusNotStarted ScheduleProvisionStatus = "NotStarted"
	ScheduleProvisionStatusRunning    ScheduleProvisionStatus = "Running"
)

func PossibleValuesForScheduleProvisionStatus() []string {
	return []string{
		string(ScheduleProvisionStatusCompleted),
		string(ScheduleProvisionStatusFailed),
		string(ScheduleProvisionStatusNotStarted),
		string(ScheduleProvisionStatusRunning),
	}
}

func parseScheduleProvisionStatus(input string) (*ScheduleProvisionStatus, error) {
	vals := map[string]ScheduleProvisionStatus{
		"completed":  ScheduleProvisionStatusCompleted,
		"failed":     ScheduleProvisionStatusFailed,
		"notstarted": ScheduleProvisionStatusNotStarted,
		"running":    ScheduleProvisionStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleProvisionStatus(input)
	return &out, nil
}
