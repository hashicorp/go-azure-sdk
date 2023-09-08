package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingMessageImportance string

const (
	CalendarSharingMessageImportancehigh   CalendarSharingMessageImportance = "High"
	CalendarSharingMessageImportancelow    CalendarSharingMessageImportance = "Low"
	CalendarSharingMessageImportancenormal CalendarSharingMessageImportance = "Normal"
)

func PossibleValuesForCalendarSharingMessageImportance() []string {
	return []string{
		string(CalendarSharingMessageImportancehigh),
		string(CalendarSharingMessageImportancelow),
		string(CalendarSharingMessageImportancenormal),
	}
}

func parseCalendarSharingMessageImportance(input string) (*CalendarSharingMessageImportance, error) {
	vals := map[string]CalendarSharingMessageImportance{
		"high":   CalendarSharingMessageImportancehigh,
		"low":    CalendarSharingMessageImportancelow,
		"normal": CalendarSharingMessageImportancenormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarSharingMessageImportance(input)
	return &out, nil
}
