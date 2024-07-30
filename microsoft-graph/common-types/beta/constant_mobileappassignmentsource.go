package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppAssignmentSource string

const (
	MobileAppAssignmentSource_Direct     MobileAppAssignmentSource = "direct"
	MobileAppAssignmentSource_PolicySets MobileAppAssignmentSource = "policySets"
)

func PossibleValuesForMobileAppAssignmentSource() []string {
	return []string{
		string(MobileAppAssignmentSource_Direct),
		string(MobileAppAssignmentSource_PolicySets),
	}
}

func (s *MobileAppAssignmentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppAssignmentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppAssignmentSource(input string) (*MobileAppAssignmentSource, error) {
	vals := map[string]MobileAppAssignmentSource{
		"direct":     MobileAppAssignmentSource_Direct,
		"policysets": MobileAppAssignmentSource_PolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppAssignmentSource(input)
	return &out, nil
}
