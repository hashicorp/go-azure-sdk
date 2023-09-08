package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateScheduledInstallScheduledInstallDay string

const (
	WindowsUpdateScheduledInstallScheduledInstallDayeveryday    WindowsUpdateScheduledInstallScheduledInstallDay = "Everyday"
	WindowsUpdateScheduledInstallScheduledInstallDayfriday      WindowsUpdateScheduledInstallScheduledInstallDay = "Friday"
	WindowsUpdateScheduledInstallScheduledInstallDaymonday      WindowsUpdateScheduledInstallScheduledInstallDay = "Monday"
	WindowsUpdateScheduledInstallScheduledInstallDaysaturday    WindowsUpdateScheduledInstallScheduledInstallDay = "Saturday"
	WindowsUpdateScheduledInstallScheduledInstallDaysunday      WindowsUpdateScheduledInstallScheduledInstallDay = "Sunday"
	WindowsUpdateScheduledInstallScheduledInstallDaythursday    WindowsUpdateScheduledInstallScheduledInstallDay = "Thursday"
	WindowsUpdateScheduledInstallScheduledInstallDaytuesday     WindowsUpdateScheduledInstallScheduledInstallDay = "Tuesday"
	WindowsUpdateScheduledInstallScheduledInstallDayuserDefined WindowsUpdateScheduledInstallScheduledInstallDay = "UserDefined"
	WindowsUpdateScheduledInstallScheduledInstallDaywednesday   WindowsUpdateScheduledInstallScheduledInstallDay = "Wednesday"
)

func PossibleValuesForWindowsUpdateScheduledInstallScheduledInstallDay() []string {
	return []string{
		string(WindowsUpdateScheduledInstallScheduledInstallDayeveryday),
		string(WindowsUpdateScheduledInstallScheduledInstallDayfriday),
		string(WindowsUpdateScheduledInstallScheduledInstallDaymonday),
		string(WindowsUpdateScheduledInstallScheduledInstallDaysaturday),
		string(WindowsUpdateScheduledInstallScheduledInstallDaysunday),
		string(WindowsUpdateScheduledInstallScheduledInstallDaythursday),
		string(WindowsUpdateScheduledInstallScheduledInstallDaytuesday),
		string(WindowsUpdateScheduledInstallScheduledInstallDayuserDefined),
		string(WindowsUpdateScheduledInstallScheduledInstallDaywednesday),
	}
}

func parseWindowsUpdateScheduledInstallScheduledInstallDay(input string) (*WindowsUpdateScheduledInstallScheduledInstallDay, error) {
	vals := map[string]WindowsUpdateScheduledInstallScheduledInstallDay{
		"everyday":    WindowsUpdateScheduledInstallScheduledInstallDayeveryday,
		"friday":      WindowsUpdateScheduledInstallScheduledInstallDayfriday,
		"monday":      WindowsUpdateScheduledInstallScheduledInstallDaymonday,
		"saturday":    WindowsUpdateScheduledInstallScheduledInstallDaysaturday,
		"sunday":      WindowsUpdateScheduledInstallScheduledInstallDaysunday,
		"thursday":    WindowsUpdateScheduledInstallScheduledInstallDaythursday,
		"tuesday":     WindowsUpdateScheduledInstallScheduledInstallDaytuesday,
		"userdefined": WindowsUpdateScheduledInstallScheduledInstallDayuserDefined,
		"wednesday":   WindowsUpdateScheduledInstallScheduledInstallDaywednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateScheduledInstallScheduledInstallDay(input)
	return &out, nil
}
