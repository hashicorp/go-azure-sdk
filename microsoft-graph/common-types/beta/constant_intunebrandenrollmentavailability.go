package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntuneBrandEnrollmentAvailability string

const (
	IntuneBrandEnrollmentAvailability_AvailableWithPrompts    IntuneBrandEnrollmentAvailability = "availableWithPrompts"
	IntuneBrandEnrollmentAvailability_AvailableWithoutPrompts IntuneBrandEnrollmentAvailability = "availableWithoutPrompts"
	IntuneBrandEnrollmentAvailability_Unavailable             IntuneBrandEnrollmentAvailability = "unavailable"
)

func PossibleValuesForIntuneBrandEnrollmentAvailability() []string {
	return []string{
		string(IntuneBrandEnrollmentAvailability_AvailableWithPrompts),
		string(IntuneBrandEnrollmentAvailability_AvailableWithoutPrompts),
		string(IntuneBrandEnrollmentAvailability_Unavailable),
	}
}

func (s *IntuneBrandEnrollmentAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIntuneBrandEnrollmentAvailability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIntuneBrandEnrollmentAvailability(input string) (*IntuneBrandEnrollmentAvailability, error) {
	vals := map[string]IntuneBrandEnrollmentAvailability{
		"availablewithprompts":    IntuneBrandEnrollmentAvailability_AvailableWithPrompts,
		"availablewithoutprompts": IntuneBrandEnrollmentAvailability_AvailableWithoutPrompts,
		"unavailable":             IntuneBrandEnrollmentAvailability_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntuneBrandEnrollmentAvailability(input)
	return &out, nil
}
