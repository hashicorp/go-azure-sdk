package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrantStatus string

const (
	MeetingRegistrantStatuscanceled   MeetingRegistrantStatus = "Canceled"
	MeetingRegistrantStatusprocessing MeetingRegistrantStatus = "Processing"
	MeetingRegistrantStatusregistered MeetingRegistrantStatus = "Registered"
)

func PossibleValuesForMeetingRegistrantStatus() []string {
	return []string{
		string(MeetingRegistrantStatuscanceled),
		string(MeetingRegistrantStatusprocessing),
		string(MeetingRegistrantStatusregistered),
	}
}

func parseMeetingRegistrantStatus(input string) (*MeetingRegistrantStatus, error) {
	vals := map[string]MeetingRegistrantStatus{
		"canceled":   MeetingRegistrantStatuscanceled,
		"processing": MeetingRegistrantStatusprocessing,
		"registered": MeetingRegistrantStatusregistered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingRegistrantStatus(input)
	return &out, nil
}
