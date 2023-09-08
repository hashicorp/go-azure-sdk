package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleItemStatus string

const (
	ScheduleItemStatusbusy             ScheduleItemStatus = "Busy"
	ScheduleItemStatusfree             ScheduleItemStatus = "Free"
	ScheduleItemStatusoof              ScheduleItemStatus = "Oof"
	ScheduleItemStatustentative        ScheduleItemStatus = "Tentative"
	ScheduleItemStatusunknown          ScheduleItemStatus = "Unknown"
	ScheduleItemStatusworkingElsewhere ScheduleItemStatus = "WorkingElsewhere"
)

func PossibleValuesForScheduleItemStatus() []string {
	return []string{
		string(ScheduleItemStatusbusy),
		string(ScheduleItemStatusfree),
		string(ScheduleItemStatusoof),
		string(ScheduleItemStatustentative),
		string(ScheduleItemStatusunknown),
		string(ScheduleItemStatusworkingElsewhere),
	}
}

func parseScheduleItemStatus(input string) (*ScheduleItemStatus, error) {
	vals := map[string]ScheduleItemStatus{
		"busy":             ScheduleItemStatusbusy,
		"free":             ScheduleItemStatusfree,
		"oof":              ScheduleItemStatusoof,
		"tentative":        ScheduleItemStatustentative,
		"unknown":          ScheduleItemStatusunknown,
		"workingelsewhere": ScheduleItemStatusworkingElsewhere,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleItemStatus(input)
	return &out, nil
}
