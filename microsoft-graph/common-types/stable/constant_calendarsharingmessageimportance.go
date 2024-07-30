package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingMessageImportance string

const (
	CalendarSharingMessageImportance_High   CalendarSharingMessageImportance = "high"
	CalendarSharingMessageImportance_Low    CalendarSharingMessageImportance = "low"
	CalendarSharingMessageImportance_Normal CalendarSharingMessageImportance = "normal"
)

func PossibleValuesForCalendarSharingMessageImportance() []string {
	return []string{
		string(CalendarSharingMessageImportance_High),
		string(CalendarSharingMessageImportance_Low),
		string(CalendarSharingMessageImportance_Normal),
	}
}

func (s *CalendarSharingMessageImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarSharingMessageImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarSharingMessageImportance(input string) (*CalendarSharingMessageImportance, error) {
	vals := map[string]CalendarSharingMessageImportance{
		"high":   CalendarSharingMessageImportance_High,
		"low":    CalendarSharingMessageImportance_Low,
		"normal": CalendarSharingMessageImportance_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarSharingMessageImportance(input)
	return &out, nil
}
