package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingTimeSuggestionOrganizerAvailability string

const (
	MeetingTimeSuggestionOrganizerAvailabilitybusy             MeetingTimeSuggestionOrganizerAvailability = "Busy"
	MeetingTimeSuggestionOrganizerAvailabilityfree             MeetingTimeSuggestionOrganizerAvailability = "Free"
	MeetingTimeSuggestionOrganizerAvailabilityoof              MeetingTimeSuggestionOrganizerAvailability = "Oof"
	MeetingTimeSuggestionOrganizerAvailabilitytentative        MeetingTimeSuggestionOrganizerAvailability = "Tentative"
	MeetingTimeSuggestionOrganizerAvailabilityunknown          MeetingTimeSuggestionOrganizerAvailability = "Unknown"
	MeetingTimeSuggestionOrganizerAvailabilityworkingElsewhere MeetingTimeSuggestionOrganizerAvailability = "WorkingElsewhere"
)

func PossibleValuesForMeetingTimeSuggestionOrganizerAvailability() []string {
	return []string{
		string(MeetingTimeSuggestionOrganizerAvailabilitybusy),
		string(MeetingTimeSuggestionOrganizerAvailabilityfree),
		string(MeetingTimeSuggestionOrganizerAvailabilityoof),
		string(MeetingTimeSuggestionOrganizerAvailabilitytentative),
		string(MeetingTimeSuggestionOrganizerAvailabilityunknown),
		string(MeetingTimeSuggestionOrganizerAvailabilityworkingElsewhere),
	}
}

func parseMeetingTimeSuggestionOrganizerAvailability(input string) (*MeetingTimeSuggestionOrganizerAvailability, error) {
	vals := map[string]MeetingTimeSuggestionOrganizerAvailability{
		"busy":             MeetingTimeSuggestionOrganizerAvailabilitybusy,
		"free":             MeetingTimeSuggestionOrganizerAvailabilityfree,
		"oof":              MeetingTimeSuggestionOrganizerAvailabilityoof,
		"tentative":        MeetingTimeSuggestionOrganizerAvailabilitytentative,
		"unknown":          MeetingTimeSuggestionOrganizerAvailabilityunknown,
		"workingelsewhere": MeetingTimeSuggestionOrganizerAvailabilityworkingElsewhere,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingTimeSuggestionOrganizerAvailability(input)
	return &out, nil
}
