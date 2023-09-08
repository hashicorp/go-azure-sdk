package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomUpdateTimeWindowStartDay string

const (
	CustomUpdateTimeWindowStartDayfriday    CustomUpdateTimeWindowStartDay = "Friday"
	CustomUpdateTimeWindowStartDaymonday    CustomUpdateTimeWindowStartDay = "Monday"
	CustomUpdateTimeWindowStartDaysaturday  CustomUpdateTimeWindowStartDay = "Saturday"
	CustomUpdateTimeWindowStartDaysunday    CustomUpdateTimeWindowStartDay = "Sunday"
	CustomUpdateTimeWindowStartDaythursday  CustomUpdateTimeWindowStartDay = "Thursday"
	CustomUpdateTimeWindowStartDaytuesday   CustomUpdateTimeWindowStartDay = "Tuesday"
	CustomUpdateTimeWindowStartDaywednesday CustomUpdateTimeWindowStartDay = "Wednesday"
)

func PossibleValuesForCustomUpdateTimeWindowStartDay() []string {
	return []string{
		string(CustomUpdateTimeWindowStartDayfriday),
		string(CustomUpdateTimeWindowStartDaymonday),
		string(CustomUpdateTimeWindowStartDaysaturday),
		string(CustomUpdateTimeWindowStartDaysunday),
		string(CustomUpdateTimeWindowStartDaythursday),
		string(CustomUpdateTimeWindowStartDaytuesday),
		string(CustomUpdateTimeWindowStartDaywednesday),
	}
}

func parseCustomUpdateTimeWindowStartDay(input string) (*CustomUpdateTimeWindowStartDay, error) {
	vals := map[string]CustomUpdateTimeWindowStartDay{
		"friday":    CustomUpdateTimeWindowStartDayfriday,
		"monday":    CustomUpdateTimeWindowStartDaymonday,
		"saturday":  CustomUpdateTimeWindowStartDaysaturday,
		"sunday":    CustomUpdateTimeWindowStartDaysunday,
		"thursday":  CustomUpdateTimeWindowStartDaythursday,
		"tuesday":   CustomUpdateTimeWindowStartDaytuesday,
		"wednesday": CustomUpdateTimeWindowStartDaywednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomUpdateTimeWindowStartDay(input)
	return &out, nil
}
