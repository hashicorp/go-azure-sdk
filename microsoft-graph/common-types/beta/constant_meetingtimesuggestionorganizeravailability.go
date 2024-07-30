package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingTimeSuggestionOrganizerAvailability string

const (
	MeetingTimeSuggestionOrganizerAvailability_Busy             MeetingTimeSuggestionOrganizerAvailability = "busy"
	MeetingTimeSuggestionOrganizerAvailability_Free             MeetingTimeSuggestionOrganizerAvailability = "free"
	MeetingTimeSuggestionOrganizerAvailability_Oof              MeetingTimeSuggestionOrganizerAvailability = "oof"
	MeetingTimeSuggestionOrganizerAvailability_Tentative        MeetingTimeSuggestionOrganizerAvailability = "tentative"
	MeetingTimeSuggestionOrganizerAvailability_Unknown          MeetingTimeSuggestionOrganizerAvailability = "unknown"
	MeetingTimeSuggestionOrganizerAvailability_WorkingElsewhere MeetingTimeSuggestionOrganizerAvailability = "workingElsewhere"
)

func PossibleValuesForMeetingTimeSuggestionOrganizerAvailability() []string {
	return []string{
		string(MeetingTimeSuggestionOrganizerAvailability_Busy),
		string(MeetingTimeSuggestionOrganizerAvailability_Free),
		string(MeetingTimeSuggestionOrganizerAvailability_Oof),
		string(MeetingTimeSuggestionOrganizerAvailability_Tentative),
		string(MeetingTimeSuggestionOrganizerAvailability_Unknown),
		string(MeetingTimeSuggestionOrganizerAvailability_WorkingElsewhere),
	}
}

func (s *MeetingTimeSuggestionOrganizerAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingTimeSuggestionOrganizerAvailability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingTimeSuggestionOrganizerAvailability(input string) (*MeetingTimeSuggestionOrganizerAvailability, error) {
	vals := map[string]MeetingTimeSuggestionOrganizerAvailability{
		"busy":             MeetingTimeSuggestionOrganizerAvailability_Busy,
		"free":             MeetingTimeSuggestionOrganizerAvailability_Free,
		"oof":              MeetingTimeSuggestionOrganizerAvailability_Oof,
		"tentative":        MeetingTimeSuggestionOrganizerAvailability_Tentative,
		"unknown":          MeetingTimeSuggestionOrganizerAvailability_Unknown,
		"workingelsewhere": MeetingTimeSuggestionOrganizerAvailability_WorkingElsewhere,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingTimeSuggestionOrganizerAvailability(input)
	return &out, nil
}
