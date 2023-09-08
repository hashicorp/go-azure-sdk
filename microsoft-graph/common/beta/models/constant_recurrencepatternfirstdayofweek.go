package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrencePatternFirstDayOfWeek string

const (
	RecurrencePatternFirstDayOfWeekfriday    RecurrencePatternFirstDayOfWeek = "Friday"
	RecurrencePatternFirstDayOfWeekmonday    RecurrencePatternFirstDayOfWeek = "Monday"
	RecurrencePatternFirstDayOfWeeksaturday  RecurrencePatternFirstDayOfWeek = "Saturday"
	RecurrencePatternFirstDayOfWeeksunday    RecurrencePatternFirstDayOfWeek = "Sunday"
	RecurrencePatternFirstDayOfWeekthursday  RecurrencePatternFirstDayOfWeek = "Thursday"
	RecurrencePatternFirstDayOfWeektuesday   RecurrencePatternFirstDayOfWeek = "Tuesday"
	RecurrencePatternFirstDayOfWeekwednesday RecurrencePatternFirstDayOfWeek = "Wednesday"
)

func PossibleValuesForRecurrencePatternFirstDayOfWeek() []string {
	return []string{
		string(RecurrencePatternFirstDayOfWeekfriday),
		string(RecurrencePatternFirstDayOfWeekmonday),
		string(RecurrencePatternFirstDayOfWeeksaturday),
		string(RecurrencePatternFirstDayOfWeeksunday),
		string(RecurrencePatternFirstDayOfWeekthursday),
		string(RecurrencePatternFirstDayOfWeektuesday),
		string(RecurrencePatternFirstDayOfWeekwednesday),
	}
}

func parseRecurrencePatternFirstDayOfWeek(input string) (*RecurrencePatternFirstDayOfWeek, error) {
	vals := map[string]RecurrencePatternFirstDayOfWeek{
		"friday":    RecurrencePatternFirstDayOfWeekfriday,
		"monday":    RecurrencePatternFirstDayOfWeekmonday,
		"saturday":  RecurrencePatternFirstDayOfWeeksaturday,
		"sunday":    RecurrencePatternFirstDayOfWeeksunday,
		"thursday":  RecurrencePatternFirstDayOfWeekthursday,
		"tuesday":   RecurrencePatternFirstDayOfWeektuesday,
		"wednesday": RecurrencePatternFirstDayOfWeekwednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrencePatternFirstDayOfWeek(input)
	return &out, nil
}
