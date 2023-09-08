package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkingHoursDaysOfWeek string

const (
	WorkingHoursDaysOfWeekfriday    WorkingHoursDaysOfWeek = "Friday"
	WorkingHoursDaysOfWeekmonday    WorkingHoursDaysOfWeek = "Monday"
	WorkingHoursDaysOfWeeksaturday  WorkingHoursDaysOfWeek = "Saturday"
	WorkingHoursDaysOfWeeksunday    WorkingHoursDaysOfWeek = "Sunday"
	WorkingHoursDaysOfWeekthursday  WorkingHoursDaysOfWeek = "Thursday"
	WorkingHoursDaysOfWeektuesday   WorkingHoursDaysOfWeek = "Tuesday"
	WorkingHoursDaysOfWeekwednesday WorkingHoursDaysOfWeek = "Wednesday"
)

func PossibleValuesForWorkingHoursDaysOfWeek() []string {
	return []string{
		string(WorkingHoursDaysOfWeekfriday),
		string(WorkingHoursDaysOfWeekmonday),
		string(WorkingHoursDaysOfWeeksaturday),
		string(WorkingHoursDaysOfWeeksunday),
		string(WorkingHoursDaysOfWeekthursday),
		string(WorkingHoursDaysOfWeektuesday),
		string(WorkingHoursDaysOfWeekwednesday),
	}
}

func parseWorkingHoursDaysOfWeek(input string) (*WorkingHoursDaysOfWeek, error) {
	vals := map[string]WorkingHoursDaysOfWeek{
		"friday":    WorkingHoursDaysOfWeekfriday,
		"monday":    WorkingHoursDaysOfWeekmonday,
		"saturday":  WorkingHoursDaysOfWeeksaturday,
		"sunday":    WorkingHoursDaysOfWeeksunday,
		"thursday":  WorkingHoursDaysOfWeekthursday,
		"tuesday":   WorkingHoursDaysOfWeektuesday,
		"wednesday": WorkingHoursDaysOfWeekwednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkingHoursDaysOfWeek(input)
	return &out, nil
}
