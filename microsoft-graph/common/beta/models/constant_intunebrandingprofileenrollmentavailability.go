package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntuneBrandingProfileEnrollmentAvailability string

const (
	IntuneBrandingProfileEnrollmentAvailabilityavailableWithPrompts    IntuneBrandingProfileEnrollmentAvailability = "AvailableWithPrompts"
	IntuneBrandingProfileEnrollmentAvailabilityavailableWithoutPrompts IntuneBrandingProfileEnrollmentAvailability = "AvailableWithoutPrompts"
	IntuneBrandingProfileEnrollmentAvailabilityunavailable             IntuneBrandingProfileEnrollmentAvailability = "Unavailable"
)

func PossibleValuesForIntuneBrandingProfileEnrollmentAvailability() []string {
	return []string{
		string(IntuneBrandingProfileEnrollmentAvailabilityavailableWithPrompts),
		string(IntuneBrandingProfileEnrollmentAvailabilityavailableWithoutPrompts),
		string(IntuneBrandingProfileEnrollmentAvailabilityunavailable),
	}
}

func parseIntuneBrandingProfileEnrollmentAvailability(input string) (*IntuneBrandingProfileEnrollmentAvailability, error) {
	vals := map[string]IntuneBrandingProfileEnrollmentAvailability{
		"availablewithprompts":    IntuneBrandingProfileEnrollmentAvailabilityavailableWithPrompts,
		"availablewithoutprompts": IntuneBrandingProfileEnrollmentAvailabilityavailableWithoutPrompts,
		"unavailable":             IntuneBrandingProfileEnrollmentAvailabilityunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntuneBrandingProfileEnrollmentAvailability(input)
	return &out, nil
}
