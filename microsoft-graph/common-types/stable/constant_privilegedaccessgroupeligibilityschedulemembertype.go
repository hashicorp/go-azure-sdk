package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleMemberType string

const (
	PrivilegedAccessGroupEligibilityScheduleMemberType_Direct PrivilegedAccessGroupEligibilityScheduleMemberType = "direct"
	PrivilegedAccessGroupEligibilityScheduleMemberType_Group  PrivilegedAccessGroupEligibilityScheduleMemberType = "group"
)

func PossibleValuesForPrivilegedAccessGroupEligibilityScheduleMemberType() []string {
	return []string{
		string(PrivilegedAccessGroupEligibilityScheduleMemberType_Direct),
		string(PrivilegedAccessGroupEligibilityScheduleMemberType_Group),
	}
}

func (s *PrivilegedAccessGroupEligibilityScheduleMemberType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupEligibilityScheduleMemberType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupEligibilityScheduleMemberType(input string) (*PrivilegedAccessGroupEligibilityScheduleMemberType, error) {
	vals := map[string]PrivilegedAccessGroupEligibilityScheduleMemberType{
		"direct": PrivilegedAccessGroupEligibilityScheduleMemberType_Direct,
		"group":  PrivilegedAccessGroupEligibilityScheduleMemberType_Group,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupEligibilityScheduleMemberType(input)
	return &out, nil
}
