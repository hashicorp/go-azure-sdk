package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskForceUpdateScheduleDayofWeek string

const (
	WindowsKioskForceUpdateScheduleDayofWeekfriday    WindowsKioskForceUpdateScheduleDayofWeek = "Friday"
	WindowsKioskForceUpdateScheduleDayofWeekmonday    WindowsKioskForceUpdateScheduleDayofWeek = "Monday"
	WindowsKioskForceUpdateScheduleDayofWeeksaturday  WindowsKioskForceUpdateScheduleDayofWeek = "Saturday"
	WindowsKioskForceUpdateScheduleDayofWeeksunday    WindowsKioskForceUpdateScheduleDayofWeek = "Sunday"
	WindowsKioskForceUpdateScheduleDayofWeekthursday  WindowsKioskForceUpdateScheduleDayofWeek = "Thursday"
	WindowsKioskForceUpdateScheduleDayofWeektuesday   WindowsKioskForceUpdateScheduleDayofWeek = "Tuesday"
	WindowsKioskForceUpdateScheduleDayofWeekwednesday WindowsKioskForceUpdateScheduleDayofWeek = "Wednesday"
)

func PossibleValuesForWindowsKioskForceUpdateScheduleDayofWeek() []string {
	return []string{
		string(WindowsKioskForceUpdateScheduleDayofWeekfriday),
		string(WindowsKioskForceUpdateScheduleDayofWeekmonday),
		string(WindowsKioskForceUpdateScheduleDayofWeeksaturday),
		string(WindowsKioskForceUpdateScheduleDayofWeeksunday),
		string(WindowsKioskForceUpdateScheduleDayofWeekthursday),
		string(WindowsKioskForceUpdateScheduleDayofWeektuesday),
		string(WindowsKioskForceUpdateScheduleDayofWeekwednesday),
	}
}

func parseWindowsKioskForceUpdateScheduleDayofWeek(input string) (*WindowsKioskForceUpdateScheduleDayofWeek, error) {
	vals := map[string]WindowsKioskForceUpdateScheduleDayofWeek{
		"friday":    WindowsKioskForceUpdateScheduleDayofWeekfriday,
		"monday":    WindowsKioskForceUpdateScheduleDayofWeekmonday,
		"saturday":  WindowsKioskForceUpdateScheduleDayofWeeksaturday,
		"sunday":    WindowsKioskForceUpdateScheduleDayofWeeksunday,
		"thursday":  WindowsKioskForceUpdateScheduleDayofWeekthursday,
		"tuesday":   WindowsKioskForceUpdateScheduleDayofWeektuesday,
		"wednesday": WindowsKioskForceUpdateScheduleDayofWeekwednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskForceUpdateScheduleDayofWeek(input)
	return &out, nil
}
