package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomUpdateTimeWindowEndDay string

const (
	CustomUpdateTimeWindowEndDayfriday    CustomUpdateTimeWindowEndDay = "Friday"
	CustomUpdateTimeWindowEndDaymonday    CustomUpdateTimeWindowEndDay = "Monday"
	CustomUpdateTimeWindowEndDaysaturday  CustomUpdateTimeWindowEndDay = "Saturday"
	CustomUpdateTimeWindowEndDaysunday    CustomUpdateTimeWindowEndDay = "Sunday"
	CustomUpdateTimeWindowEndDaythursday  CustomUpdateTimeWindowEndDay = "Thursday"
	CustomUpdateTimeWindowEndDaytuesday   CustomUpdateTimeWindowEndDay = "Tuesday"
	CustomUpdateTimeWindowEndDaywednesday CustomUpdateTimeWindowEndDay = "Wednesday"
)

func PossibleValuesForCustomUpdateTimeWindowEndDay() []string {
	return []string{
		string(CustomUpdateTimeWindowEndDayfriday),
		string(CustomUpdateTimeWindowEndDaymonday),
		string(CustomUpdateTimeWindowEndDaysaturday),
		string(CustomUpdateTimeWindowEndDaysunday),
		string(CustomUpdateTimeWindowEndDaythursday),
		string(CustomUpdateTimeWindowEndDaytuesday),
		string(CustomUpdateTimeWindowEndDaywednesday),
	}
}

func parseCustomUpdateTimeWindowEndDay(input string) (*CustomUpdateTimeWindowEndDay, error) {
	vals := map[string]CustomUpdateTimeWindowEndDay{
		"friday":    CustomUpdateTimeWindowEndDayfriday,
		"monday":    CustomUpdateTimeWindowEndDaymonday,
		"saturday":  CustomUpdateTimeWindowEndDaysaturday,
		"sunday":    CustomUpdateTimeWindowEndDaysunday,
		"thursday":  CustomUpdateTimeWindowEndDaythursday,
		"tuesday":   CustomUpdateTimeWindowEndDaytuesday,
		"wednesday": CustomUpdateTimeWindowEndDaywednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomUpdateTimeWindowEndDay(input)
	return &out, nil
}
