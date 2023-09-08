package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntuneBrandEnrollmentAvailability string

const (
	IntuneBrandEnrollmentAvailabilityavailableWithPrompts    IntuneBrandEnrollmentAvailability = "AvailableWithPrompts"
	IntuneBrandEnrollmentAvailabilityavailableWithoutPrompts IntuneBrandEnrollmentAvailability = "AvailableWithoutPrompts"
	IntuneBrandEnrollmentAvailabilityunavailable             IntuneBrandEnrollmentAvailability = "Unavailable"
)

func PossibleValuesForIntuneBrandEnrollmentAvailability() []string {
	return []string{
		string(IntuneBrandEnrollmentAvailabilityavailableWithPrompts),
		string(IntuneBrandEnrollmentAvailabilityavailableWithoutPrompts),
		string(IntuneBrandEnrollmentAvailabilityunavailable),
	}
}

func parseIntuneBrandEnrollmentAvailability(input string) (*IntuneBrandEnrollmentAvailability, error) {
	vals := map[string]IntuneBrandEnrollmentAvailability{
		"availablewithprompts":    IntuneBrandEnrollmentAvailabilityavailableWithPrompts,
		"availablewithoutprompts": IntuneBrandEnrollmentAvailabilityavailableWithoutPrompts,
		"unavailable":             IntuneBrandEnrollmentAvailabilityunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntuneBrandEnrollmentAvailability(input)
	return &out, nil
}
