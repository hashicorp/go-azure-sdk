package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrencePatternFirstDayOfWeek string

const (
	RecurrencePatternFirstDayOfWeek_Friday    RecurrencePatternFirstDayOfWeek = "friday"
	RecurrencePatternFirstDayOfWeek_Monday    RecurrencePatternFirstDayOfWeek = "monday"
	RecurrencePatternFirstDayOfWeek_Saturday  RecurrencePatternFirstDayOfWeek = "saturday"
	RecurrencePatternFirstDayOfWeek_Sunday    RecurrencePatternFirstDayOfWeek = "sunday"
	RecurrencePatternFirstDayOfWeek_Thursday  RecurrencePatternFirstDayOfWeek = "thursday"
	RecurrencePatternFirstDayOfWeek_Tuesday   RecurrencePatternFirstDayOfWeek = "tuesday"
	RecurrencePatternFirstDayOfWeek_Wednesday RecurrencePatternFirstDayOfWeek = "wednesday"
)

func PossibleValuesForRecurrencePatternFirstDayOfWeek() []string {
	return []string{
		string(RecurrencePatternFirstDayOfWeek_Friday),
		string(RecurrencePatternFirstDayOfWeek_Monday),
		string(RecurrencePatternFirstDayOfWeek_Saturday),
		string(RecurrencePatternFirstDayOfWeek_Sunday),
		string(RecurrencePatternFirstDayOfWeek_Thursday),
		string(RecurrencePatternFirstDayOfWeek_Tuesday),
		string(RecurrencePatternFirstDayOfWeek_Wednesday),
	}
}

func (s *RecurrencePatternFirstDayOfWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecurrencePatternFirstDayOfWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecurrencePatternFirstDayOfWeek(input string) (*RecurrencePatternFirstDayOfWeek, error) {
	vals := map[string]RecurrencePatternFirstDayOfWeek{
		"friday":    RecurrencePatternFirstDayOfWeek_Friday,
		"monday":    RecurrencePatternFirstDayOfWeek_Monday,
		"saturday":  RecurrencePatternFirstDayOfWeek_Saturday,
		"sunday":    RecurrencePatternFirstDayOfWeek_Sunday,
		"thursday":  RecurrencePatternFirstDayOfWeek_Thursday,
		"tuesday":   RecurrencePatternFirstDayOfWeek_Tuesday,
		"wednesday": RecurrencePatternFirstDayOfWeek_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrencePatternFirstDayOfWeek(input)
	return &out, nil
}
