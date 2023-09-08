package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrationBaseAllowedRegistrant string

const (
	MeetingRegistrationBaseAllowedRegistranteveryone     MeetingRegistrationBaseAllowedRegistrant = "Everyone"
	MeetingRegistrationBaseAllowedRegistrantorganization MeetingRegistrationBaseAllowedRegistrant = "Organization"
)

func PossibleValuesForMeetingRegistrationBaseAllowedRegistrant() []string {
	return []string{
		string(MeetingRegistrationBaseAllowedRegistranteveryone),
		string(MeetingRegistrationBaseAllowedRegistrantorganization),
	}
}

func parseMeetingRegistrationBaseAllowedRegistrant(input string) (*MeetingRegistrationBaseAllowedRegistrant, error) {
	vals := map[string]MeetingRegistrationBaseAllowedRegistrant{
		"everyone":     MeetingRegistrationBaseAllowedRegistranteveryone,
		"organization": MeetingRegistrationBaseAllowedRegistrantorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingRegistrationBaseAllowedRegistrant(input)
	return &out, nil
}
