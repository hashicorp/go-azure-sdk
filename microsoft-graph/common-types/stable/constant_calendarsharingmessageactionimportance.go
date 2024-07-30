package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingMessageActionImportance string

const (
	CalendarSharingMessageActionImportance_Primary   CalendarSharingMessageActionImportance = "primary"
	CalendarSharingMessageActionImportance_Secondary CalendarSharingMessageActionImportance = "secondary"
)

func PossibleValuesForCalendarSharingMessageActionImportance() []string {
	return []string{
		string(CalendarSharingMessageActionImportance_Primary),
		string(CalendarSharingMessageActionImportance_Secondary),
	}
}

func (s *CalendarSharingMessageActionImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarSharingMessageActionImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarSharingMessageActionImportance(input string) (*CalendarSharingMessageActionImportance, error) {
	vals := map[string]CalendarSharingMessageActionImportance{
		"primary":   CalendarSharingMessageActionImportance_Primary,
		"secondary": CalendarSharingMessageActionImportance_Secondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarSharingMessageActionImportance(input)
	return &out, nil
}
