package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaylightTimeZoneOffsetDayOfWeek string

const (
	DaylightTimeZoneOffsetDayOfWeekfriday    DaylightTimeZoneOffsetDayOfWeek = "Friday"
	DaylightTimeZoneOffsetDayOfWeekmonday    DaylightTimeZoneOffsetDayOfWeek = "Monday"
	DaylightTimeZoneOffsetDayOfWeeksaturday  DaylightTimeZoneOffsetDayOfWeek = "Saturday"
	DaylightTimeZoneOffsetDayOfWeeksunday    DaylightTimeZoneOffsetDayOfWeek = "Sunday"
	DaylightTimeZoneOffsetDayOfWeekthursday  DaylightTimeZoneOffsetDayOfWeek = "Thursday"
	DaylightTimeZoneOffsetDayOfWeektuesday   DaylightTimeZoneOffsetDayOfWeek = "Tuesday"
	DaylightTimeZoneOffsetDayOfWeekwednesday DaylightTimeZoneOffsetDayOfWeek = "Wednesday"
)

func PossibleValuesForDaylightTimeZoneOffsetDayOfWeek() []string {
	return []string{
		string(DaylightTimeZoneOffsetDayOfWeekfriday),
		string(DaylightTimeZoneOffsetDayOfWeekmonday),
		string(DaylightTimeZoneOffsetDayOfWeeksaturday),
		string(DaylightTimeZoneOffsetDayOfWeeksunday),
		string(DaylightTimeZoneOffsetDayOfWeekthursday),
		string(DaylightTimeZoneOffsetDayOfWeektuesday),
		string(DaylightTimeZoneOffsetDayOfWeekwednesday),
	}
}

func parseDaylightTimeZoneOffsetDayOfWeek(input string) (*DaylightTimeZoneOffsetDayOfWeek, error) {
	vals := map[string]DaylightTimeZoneOffsetDayOfWeek{
		"friday":    DaylightTimeZoneOffsetDayOfWeekfriday,
		"monday":    DaylightTimeZoneOffsetDayOfWeekmonday,
		"saturday":  DaylightTimeZoneOffsetDayOfWeeksaturday,
		"sunday":    DaylightTimeZoneOffsetDayOfWeeksunday,
		"thursday":  DaylightTimeZoneOffsetDayOfWeekthursday,
		"tuesday":   DaylightTimeZoneOffsetDayOfWeektuesday,
		"wednesday": DaylightTimeZoneOffsetDayOfWeekwednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DaylightTimeZoneOffsetDayOfWeek(input)
	return &out, nil
}
