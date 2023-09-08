package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingMessageActionImportance string

const (
	CalendarSharingMessageActionImportanceprimary   CalendarSharingMessageActionImportance = "Primary"
	CalendarSharingMessageActionImportancesecondary CalendarSharingMessageActionImportance = "Secondary"
)

func PossibleValuesForCalendarSharingMessageActionImportance() []string {
	return []string{
		string(CalendarSharingMessageActionImportanceprimary),
		string(CalendarSharingMessageActionImportancesecondary),
	}
}

func parseCalendarSharingMessageActionImportance(input string) (*CalendarSharingMessageActionImportance, error) {
	vals := map[string]CalendarSharingMessageActionImportance{
		"primary":   CalendarSharingMessageActionImportanceprimary,
		"secondary": CalendarSharingMessageActionImportancesecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarSharingMessageActionImportance(input)
	return &out, nil
}
