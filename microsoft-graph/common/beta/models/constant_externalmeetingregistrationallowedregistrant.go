package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalMeetingRegistrationAllowedRegistrant string

const (
	ExternalMeetingRegistrationAllowedRegistranteveryone     ExternalMeetingRegistrationAllowedRegistrant = "Everyone"
	ExternalMeetingRegistrationAllowedRegistrantorganization ExternalMeetingRegistrationAllowedRegistrant = "Organization"
)

func PossibleValuesForExternalMeetingRegistrationAllowedRegistrant() []string {
	return []string{
		string(ExternalMeetingRegistrationAllowedRegistranteveryone),
		string(ExternalMeetingRegistrationAllowedRegistrantorganization),
	}
}

func parseExternalMeetingRegistrationAllowedRegistrant(input string) (*ExternalMeetingRegistrationAllowedRegistrant, error) {
	vals := map[string]ExternalMeetingRegistrationAllowedRegistrant{
		"everyone":     ExternalMeetingRegistrationAllowedRegistranteveryone,
		"organization": ExternalMeetingRegistrationAllowedRegistrantorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalMeetingRegistrationAllowedRegistrant(input)
	return &out, nil
}
