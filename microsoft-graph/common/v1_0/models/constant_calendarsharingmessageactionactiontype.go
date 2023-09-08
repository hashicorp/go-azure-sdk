package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingMessageActionActionType string

const (
	CalendarSharingMessageActionActionTypeaccept CalendarSharingMessageActionActionType = "Accept"
)

func PossibleValuesForCalendarSharingMessageActionActionType() []string {
	return []string{
		string(CalendarSharingMessageActionActionTypeaccept),
	}
}

func parseCalendarSharingMessageActionActionType(input string) (*CalendarSharingMessageActionActionType, error) {
	vals := map[string]CalendarSharingMessageActionActionType{
		"accept": CalendarSharingMessageActionActionTypeaccept,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarSharingMessageActionActionType(input)
	return &out, nil
}
