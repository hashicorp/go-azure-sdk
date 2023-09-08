package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleChangeRequestAssignedTo string

const (
	ScheduleChangeRequestAssignedTomanager   ScheduleChangeRequestAssignedTo = "Manager"
	ScheduleChangeRequestAssignedTorecipient ScheduleChangeRequestAssignedTo = "Recipient"
	ScheduleChangeRequestAssignedTosender    ScheduleChangeRequestAssignedTo = "Sender"
	ScheduleChangeRequestAssignedTosystem    ScheduleChangeRequestAssignedTo = "System"
)

func PossibleValuesForScheduleChangeRequestAssignedTo() []string {
	return []string{
		string(ScheduleChangeRequestAssignedTomanager),
		string(ScheduleChangeRequestAssignedTorecipient),
		string(ScheduleChangeRequestAssignedTosender),
		string(ScheduleChangeRequestAssignedTosystem),
	}
}

func parseScheduleChangeRequestAssignedTo(input string) (*ScheduleChangeRequestAssignedTo, error) {
	vals := map[string]ScheduleChangeRequestAssignedTo{
		"manager":   ScheduleChangeRequestAssignedTomanager,
		"recipient": ScheduleChangeRequestAssignedTorecipient,
		"sender":    ScheduleChangeRequestAssignedTosender,
		"system":    ScheduleChangeRequestAssignedTosystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleChangeRequestAssignedTo(input)
	return &out, nil
}
