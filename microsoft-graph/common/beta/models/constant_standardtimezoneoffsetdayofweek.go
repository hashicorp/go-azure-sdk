package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardTimeZoneOffsetDayOfWeek string

const (
	StandardTimeZoneOffsetDayOfWeekfriday    StandardTimeZoneOffsetDayOfWeek = "Friday"
	StandardTimeZoneOffsetDayOfWeekmonday    StandardTimeZoneOffsetDayOfWeek = "Monday"
	StandardTimeZoneOffsetDayOfWeeksaturday  StandardTimeZoneOffsetDayOfWeek = "Saturday"
	StandardTimeZoneOffsetDayOfWeeksunday    StandardTimeZoneOffsetDayOfWeek = "Sunday"
	StandardTimeZoneOffsetDayOfWeekthursday  StandardTimeZoneOffsetDayOfWeek = "Thursday"
	StandardTimeZoneOffsetDayOfWeektuesday   StandardTimeZoneOffsetDayOfWeek = "Tuesday"
	StandardTimeZoneOffsetDayOfWeekwednesday StandardTimeZoneOffsetDayOfWeek = "Wednesday"
)

func PossibleValuesForStandardTimeZoneOffsetDayOfWeek() []string {
	return []string{
		string(StandardTimeZoneOffsetDayOfWeekfriday),
		string(StandardTimeZoneOffsetDayOfWeekmonday),
		string(StandardTimeZoneOffsetDayOfWeeksaturday),
		string(StandardTimeZoneOffsetDayOfWeeksunday),
		string(StandardTimeZoneOffsetDayOfWeekthursday),
		string(StandardTimeZoneOffsetDayOfWeektuesday),
		string(StandardTimeZoneOffsetDayOfWeekwednesday),
	}
}

func parseStandardTimeZoneOffsetDayOfWeek(input string) (*StandardTimeZoneOffsetDayOfWeek, error) {
	vals := map[string]StandardTimeZoneOffsetDayOfWeek{
		"friday":    StandardTimeZoneOffsetDayOfWeekfriday,
		"monday":    StandardTimeZoneOffsetDayOfWeekmonday,
		"saturday":  StandardTimeZoneOffsetDayOfWeeksaturday,
		"sunday":    StandardTimeZoneOffsetDayOfWeeksunday,
		"thursday":  StandardTimeZoneOffsetDayOfWeekthursday,
		"tuesday":   StandardTimeZoneOffsetDayOfWeektuesday,
		"wednesday": StandardTimeZoneOffsetDayOfWeekwednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StandardTimeZoneOffsetDayOfWeek(input)
	return &out, nil
}
