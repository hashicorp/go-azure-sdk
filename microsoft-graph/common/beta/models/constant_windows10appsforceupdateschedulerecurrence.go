package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10AppsForceUpdateScheduleRecurrence string

const (
	Windows10AppsForceUpdateScheduleRecurrencedaily   Windows10AppsForceUpdateScheduleRecurrence = "Daily"
	Windows10AppsForceUpdateScheduleRecurrencemonthly Windows10AppsForceUpdateScheduleRecurrence = "Monthly"
	Windows10AppsForceUpdateScheduleRecurrencenone    Windows10AppsForceUpdateScheduleRecurrence = "None"
	Windows10AppsForceUpdateScheduleRecurrenceweekly  Windows10AppsForceUpdateScheduleRecurrence = "Weekly"
)

func PossibleValuesForWindows10AppsForceUpdateScheduleRecurrence() []string {
	return []string{
		string(Windows10AppsForceUpdateScheduleRecurrencedaily),
		string(Windows10AppsForceUpdateScheduleRecurrencemonthly),
		string(Windows10AppsForceUpdateScheduleRecurrencenone),
		string(Windows10AppsForceUpdateScheduleRecurrenceweekly),
	}
}

func parseWindows10AppsForceUpdateScheduleRecurrence(input string) (*Windows10AppsForceUpdateScheduleRecurrence, error) {
	vals := map[string]Windows10AppsForceUpdateScheduleRecurrence{
		"daily":   Windows10AppsForceUpdateScheduleRecurrencedaily,
		"monthly": Windows10AppsForceUpdateScheduleRecurrencemonthly,
		"none":    Windows10AppsForceUpdateScheduleRecurrencenone,
		"weekly":  Windows10AppsForceUpdateScheduleRecurrenceweekly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10AppsForceUpdateScheduleRecurrence(input)
	return &out, nil
}
