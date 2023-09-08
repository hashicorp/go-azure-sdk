package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrencePatternDaysOfWeek string

const (
	RecurrencePatternDaysOfWeekfriday    RecurrencePatternDaysOfWeek = "Friday"
	RecurrencePatternDaysOfWeekmonday    RecurrencePatternDaysOfWeek = "Monday"
	RecurrencePatternDaysOfWeeksaturday  RecurrencePatternDaysOfWeek = "Saturday"
	RecurrencePatternDaysOfWeeksunday    RecurrencePatternDaysOfWeek = "Sunday"
	RecurrencePatternDaysOfWeekthursday  RecurrencePatternDaysOfWeek = "Thursday"
	RecurrencePatternDaysOfWeektuesday   RecurrencePatternDaysOfWeek = "Tuesday"
	RecurrencePatternDaysOfWeekwednesday RecurrencePatternDaysOfWeek = "Wednesday"
)

func PossibleValuesForRecurrencePatternDaysOfWeek() []string {
	return []string{
		string(RecurrencePatternDaysOfWeekfriday),
		string(RecurrencePatternDaysOfWeekmonday),
		string(RecurrencePatternDaysOfWeeksaturday),
		string(RecurrencePatternDaysOfWeeksunday),
		string(RecurrencePatternDaysOfWeekthursday),
		string(RecurrencePatternDaysOfWeektuesday),
		string(RecurrencePatternDaysOfWeekwednesday),
	}
}

func parseRecurrencePatternDaysOfWeek(input string) (*RecurrencePatternDaysOfWeek, error) {
	vals := map[string]RecurrencePatternDaysOfWeek{
		"friday":    RecurrencePatternDaysOfWeekfriday,
		"monday":    RecurrencePatternDaysOfWeekmonday,
		"saturday":  RecurrencePatternDaysOfWeeksaturday,
		"sunday":    RecurrencePatternDaysOfWeeksunday,
		"thursday":  RecurrencePatternDaysOfWeekthursday,
		"tuesday":   RecurrencePatternDaysOfWeektuesday,
		"wednesday": RecurrencePatternDaysOfWeekwednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrencePatternDaysOfWeek(input)
	return &out, nil
}
