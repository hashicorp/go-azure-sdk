package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskForceUpdateScheduleRecurrence string

const (
	WindowsKioskForceUpdateScheduleRecurrencedaily   WindowsKioskForceUpdateScheduleRecurrence = "Daily"
	WindowsKioskForceUpdateScheduleRecurrencemonthly WindowsKioskForceUpdateScheduleRecurrence = "Monthly"
	WindowsKioskForceUpdateScheduleRecurrencenone    WindowsKioskForceUpdateScheduleRecurrence = "None"
	WindowsKioskForceUpdateScheduleRecurrenceweekly  WindowsKioskForceUpdateScheduleRecurrence = "Weekly"
)

func PossibleValuesForWindowsKioskForceUpdateScheduleRecurrence() []string {
	return []string{
		string(WindowsKioskForceUpdateScheduleRecurrencedaily),
		string(WindowsKioskForceUpdateScheduleRecurrencemonthly),
		string(WindowsKioskForceUpdateScheduleRecurrencenone),
		string(WindowsKioskForceUpdateScheduleRecurrenceweekly),
	}
}

func parseWindowsKioskForceUpdateScheduleRecurrence(input string) (*WindowsKioskForceUpdateScheduleRecurrence, error) {
	vals := map[string]WindowsKioskForceUpdateScheduleRecurrence{
		"daily":   WindowsKioskForceUpdateScheduleRecurrencedaily,
		"monthly": WindowsKioskForceUpdateScheduleRecurrencemonthly,
		"none":    WindowsKioskForceUpdateScheduleRecurrencenone,
		"weekly":  WindowsKioskForceUpdateScheduleRecurrenceweekly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskForceUpdateScheduleRecurrence(input)
	return &out, nil
}
