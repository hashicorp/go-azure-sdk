package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendeeAvailabilityAvailability string

const (
	AttendeeAvailabilityAvailabilitybusy             AttendeeAvailabilityAvailability = "Busy"
	AttendeeAvailabilityAvailabilityfree             AttendeeAvailabilityAvailability = "Free"
	AttendeeAvailabilityAvailabilityoof              AttendeeAvailabilityAvailability = "Oof"
	AttendeeAvailabilityAvailabilitytentative        AttendeeAvailabilityAvailability = "Tentative"
	AttendeeAvailabilityAvailabilityunknown          AttendeeAvailabilityAvailability = "Unknown"
	AttendeeAvailabilityAvailabilityworkingElsewhere AttendeeAvailabilityAvailability = "WorkingElsewhere"
)

func PossibleValuesForAttendeeAvailabilityAvailability() []string {
	return []string{
		string(AttendeeAvailabilityAvailabilitybusy),
		string(AttendeeAvailabilityAvailabilityfree),
		string(AttendeeAvailabilityAvailabilityoof),
		string(AttendeeAvailabilityAvailabilitytentative),
		string(AttendeeAvailabilityAvailabilityunknown),
		string(AttendeeAvailabilityAvailabilityworkingElsewhere),
	}
}

func parseAttendeeAvailabilityAvailability(input string) (*AttendeeAvailabilityAvailability, error) {
	vals := map[string]AttendeeAvailabilityAvailability{
		"busy":             AttendeeAvailabilityAvailabilitybusy,
		"free":             AttendeeAvailabilityAvailabilityfree,
		"oof":              AttendeeAvailabilityAvailabilityoof,
		"tentative":        AttendeeAvailabilityAvailabilitytentative,
		"unknown":          AttendeeAvailabilityAvailabilityunknown,
		"workingelsewhere": AttendeeAvailabilityAvailabilityworkingElsewhere,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttendeeAvailabilityAvailability(input)
	return &out, nil
}
