package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrationAllowedRegistrant string

const (
	MeetingRegistrationAllowedRegistranteveryone     MeetingRegistrationAllowedRegistrant = "Everyone"
	MeetingRegistrationAllowedRegistrantorganization MeetingRegistrationAllowedRegistrant = "Organization"
)

func PossibleValuesForMeetingRegistrationAllowedRegistrant() []string {
	return []string{
		string(MeetingRegistrationAllowedRegistranteveryone),
		string(MeetingRegistrationAllowedRegistrantorganization),
	}
}

func parseMeetingRegistrationAllowedRegistrant(input string) (*MeetingRegistrationAllowedRegistrant, error) {
	vals := map[string]MeetingRegistrationAllowedRegistrant{
		"everyone":     MeetingRegistrationAllowedRegistranteveryone,
		"organization": MeetingRegistrationAllowedRegistrantorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingRegistrationAllowedRegistrant(input)
	return &out, nil
}
