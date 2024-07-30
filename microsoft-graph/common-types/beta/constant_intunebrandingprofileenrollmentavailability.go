package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntuneBrandingProfileEnrollmentAvailability string

const (
	IntuneBrandingProfileEnrollmentAvailability_AvailableWithPrompts    IntuneBrandingProfileEnrollmentAvailability = "availableWithPrompts"
	IntuneBrandingProfileEnrollmentAvailability_AvailableWithoutPrompts IntuneBrandingProfileEnrollmentAvailability = "availableWithoutPrompts"
	IntuneBrandingProfileEnrollmentAvailability_Unavailable             IntuneBrandingProfileEnrollmentAvailability = "unavailable"
)

func PossibleValuesForIntuneBrandingProfileEnrollmentAvailability() []string {
	return []string{
		string(IntuneBrandingProfileEnrollmentAvailability_AvailableWithPrompts),
		string(IntuneBrandingProfileEnrollmentAvailability_AvailableWithoutPrompts),
		string(IntuneBrandingProfileEnrollmentAvailability_Unavailable),
	}
}

func (s *IntuneBrandingProfileEnrollmentAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIntuneBrandingProfileEnrollmentAvailability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIntuneBrandingProfileEnrollmentAvailability(input string) (*IntuneBrandingProfileEnrollmentAvailability, error) {
	vals := map[string]IntuneBrandingProfileEnrollmentAvailability{
		"availablewithprompts":    IntuneBrandingProfileEnrollmentAvailability_AvailableWithPrompts,
		"availablewithoutprompts": IntuneBrandingProfileEnrollmentAvailability_AvailableWithoutPrompts,
		"unavailable":             IntuneBrandingProfileEnrollmentAvailability_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntuneBrandingProfileEnrollmentAvailability(input)
	return &out, nil
}
