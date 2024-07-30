package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppAssignmentIntent string

const (
	MobileAppAssignmentIntent_Available                  MobileAppAssignmentIntent = "available"
	MobileAppAssignmentIntent_AvailableWithoutEnrollment MobileAppAssignmentIntent = "availableWithoutEnrollment"
	MobileAppAssignmentIntent_Required                   MobileAppAssignmentIntent = "required"
	MobileAppAssignmentIntent_Uninstall                  MobileAppAssignmentIntent = "uninstall"
)

func PossibleValuesForMobileAppAssignmentIntent() []string {
	return []string{
		string(MobileAppAssignmentIntent_Available),
		string(MobileAppAssignmentIntent_AvailableWithoutEnrollment),
		string(MobileAppAssignmentIntent_Required),
		string(MobileAppAssignmentIntent_Uninstall),
	}
}

func (s *MobileAppAssignmentIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppAssignmentIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppAssignmentIntent(input string) (*MobileAppAssignmentIntent, error) {
	vals := map[string]MobileAppAssignmentIntent{
		"available":                  MobileAppAssignmentIntent_Available,
		"availablewithoutenrollment": MobileAppAssignmentIntent_AvailableWithoutEnrollment,
		"required":                   MobileAppAssignmentIntent_Required,
		"uninstall":                  MobileAppAssignmentIntent_Uninstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppAssignmentIntent(input)
	return &out, nil
}
