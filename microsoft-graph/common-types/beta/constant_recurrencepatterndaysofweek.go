package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrencePatternDaysOfWeek string

const (
	RecurrencePatternDaysOfWeek_Friday    RecurrencePatternDaysOfWeek = "friday"
	RecurrencePatternDaysOfWeek_Monday    RecurrencePatternDaysOfWeek = "monday"
	RecurrencePatternDaysOfWeek_Saturday  RecurrencePatternDaysOfWeek = "saturday"
	RecurrencePatternDaysOfWeek_Sunday    RecurrencePatternDaysOfWeek = "sunday"
	RecurrencePatternDaysOfWeek_Thursday  RecurrencePatternDaysOfWeek = "thursday"
	RecurrencePatternDaysOfWeek_Tuesday   RecurrencePatternDaysOfWeek = "tuesday"
	RecurrencePatternDaysOfWeek_Wednesday RecurrencePatternDaysOfWeek = "wednesday"
)

func PossibleValuesForRecurrencePatternDaysOfWeek() []string {
	return []string{
		string(RecurrencePatternDaysOfWeek_Friday),
		string(RecurrencePatternDaysOfWeek_Monday),
		string(RecurrencePatternDaysOfWeek_Saturday),
		string(RecurrencePatternDaysOfWeek_Sunday),
		string(RecurrencePatternDaysOfWeek_Thursday),
		string(RecurrencePatternDaysOfWeek_Tuesday),
		string(RecurrencePatternDaysOfWeek_Wednesday),
	}
}

func (s *RecurrencePatternDaysOfWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecurrencePatternDaysOfWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecurrencePatternDaysOfWeek(input string) (*RecurrencePatternDaysOfWeek, error) {
	vals := map[string]RecurrencePatternDaysOfWeek{
		"friday":    RecurrencePatternDaysOfWeek_Friday,
		"monday":    RecurrencePatternDaysOfWeek_Monday,
		"saturday":  RecurrencePatternDaysOfWeek_Saturday,
		"sunday":    RecurrencePatternDaysOfWeek_Sunday,
		"thursday":  RecurrencePatternDaysOfWeek_Thursday,
		"tuesday":   RecurrencePatternDaysOfWeek_Tuesday,
		"wednesday": RecurrencePatternDaysOfWeek_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrencePatternDaysOfWeek(input)
	return &out, nil
}
