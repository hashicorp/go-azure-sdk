package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendeeAvailabilityAvailability string

const (
	AttendeeAvailabilityAvailability_Busy             AttendeeAvailabilityAvailability = "busy"
	AttendeeAvailabilityAvailability_Free             AttendeeAvailabilityAvailability = "free"
	AttendeeAvailabilityAvailability_Oof              AttendeeAvailabilityAvailability = "oof"
	AttendeeAvailabilityAvailability_Tentative        AttendeeAvailabilityAvailability = "tentative"
	AttendeeAvailabilityAvailability_Unknown          AttendeeAvailabilityAvailability = "unknown"
	AttendeeAvailabilityAvailability_WorkingElsewhere AttendeeAvailabilityAvailability = "workingElsewhere"
)

func PossibleValuesForAttendeeAvailabilityAvailability() []string {
	return []string{
		string(AttendeeAvailabilityAvailability_Busy),
		string(AttendeeAvailabilityAvailability_Free),
		string(AttendeeAvailabilityAvailability_Oof),
		string(AttendeeAvailabilityAvailability_Tentative),
		string(AttendeeAvailabilityAvailability_Unknown),
		string(AttendeeAvailabilityAvailability_WorkingElsewhere),
	}
}

func (s *AttendeeAvailabilityAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttendeeAvailabilityAvailability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttendeeAvailabilityAvailability(input string) (*AttendeeAvailabilityAvailability, error) {
	vals := map[string]AttendeeAvailabilityAvailability{
		"busy":             AttendeeAvailabilityAvailability_Busy,
		"free":             AttendeeAvailabilityAvailability_Free,
		"oof":              AttendeeAvailabilityAvailability_Oof,
		"tentative":        AttendeeAvailabilityAvailability_Tentative,
		"unknown":          AttendeeAvailabilityAvailability_Unknown,
		"workingelsewhere": AttendeeAvailabilityAvailability_WorkingElsewhere,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttendeeAvailabilityAvailability(input)
	return &out, nil
}
