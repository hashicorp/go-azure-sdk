package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingMessageActionAction string

const (
	CalendarSharingMessageActionActionaccept                CalendarSharingMessageActionAction = "Accept"
	CalendarSharingMessageActionActionacceptAndViewCalendar CalendarSharingMessageActionAction = "AcceptAndViewCalendar"
	CalendarSharingMessageActionActionaddThisCalendar       CalendarSharingMessageActionAction = "AddThisCalendar"
	CalendarSharingMessageActionActionviewCalendar          CalendarSharingMessageActionAction = "ViewCalendar"
)

func PossibleValuesForCalendarSharingMessageActionAction() []string {
	return []string{
		string(CalendarSharingMessageActionActionaccept),
		string(CalendarSharingMessageActionActionacceptAndViewCalendar),
		string(CalendarSharingMessageActionActionaddThisCalendar),
		string(CalendarSharingMessageActionActionviewCalendar),
	}
}

func parseCalendarSharingMessageActionAction(input string) (*CalendarSharingMessageActionAction, error) {
	vals := map[string]CalendarSharingMessageActionAction{
		"accept":                CalendarSharingMessageActionActionaccept,
		"acceptandviewcalendar": CalendarSharingMessageActionActionacceptAndViewCalendar,
		"addthiscalendar":       CalendarSharingMessageActionActionaddThisCalendar,
		"viewcalendar":          CalendarSharingMessageActionActionviewCalendar,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarSharingMessageActionAction(input)
	return &out, nil
}
