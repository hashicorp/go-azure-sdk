package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingMessageActionAction string

const (
	CalendarSharingMessageActionAction_Accept                CalendarSharingMessageActionAction = "accept"
	CalendarSharingMessageActionAction_AcceptAndViewCalendar CalendarSharingMessageActionAction = "acceptAndViewCalendar"
	CalendarSharingMessageActionAction_AddThisCalendar       CalendarSharingMessageActionAction = "addThisCalendar"
	CalendarSharingMessageActionAction_ViewCalendar          CalendarSharingMessageActionAction = "viewCalendar"
)

func PossibleValuesForCalendarSharingMessageActionAction() []string {
	return []string{
		string(CalendarSharingMessageActionAction_Accept),
		string(CalendarSharingMessageActionAction_AcceptAndViewCalendar),
		string(CalendarSharingMessageActionAction_AddThisCalendar),
		string(CalendarSharingMessageActionAction_ViewCalendar),
	}
}

func (s *CalendarSharingMessageActionAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarSharingMessageActionAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarSharingMessageActionAction(input string) (*CalendarSharingMessageActionAction, error) {
	vals := map[string]CalendarSharingMessageActionAction{
		"accept":                CalendarSharingMessageActionAction_Accept,
		"acceptandviewcalendar": CalendarSharingMessageActionAction_AcceptAndViewCalendar,
		"addthiscalendar":       CalendarSharingMessageActionAction_AddThisCalendar,
		"viewcalendar":          CalendarSharingMessageActionAction_ViewCalendar,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarSharingMessageActionAction(input)
	return &out, nil
}
