package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleInstanceMemberType string

const (
	PrivilegedAccessGroupEligibilityScheduleInstanceMemberType_Direct PrivilegedAccessGroupEligibilityScheduleInstanceMemberType = "direct"
	PrivilegedAccessGroupEligibilityScheduleInstanceMemberType_Group  PrivilegedAccessGroupEligibilityScheduleInstanceMemberType = "group"
)

func PossibleValuesForPrivilegedAccessGroupEligibilityScheduleInstanceMemberType() []string {
	return []string{
		string(PrivilegedAccessGroupEligibilityScheduleInstanceMemberType_Direct),
		string(PrivilegedAccessGroupEligibilityScheduleInstanceMemberType_Group),
	}
}

func (s *PrivilegedAccessGroupEligibilityScheduleInstanceMemberType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupEligibilityScheduleInstanceMemberType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupEligibilityScheduleInstanceMemberType(input string) (*PrivilegedAccessGroupEligibilityScheduleInstanceMemberType, error) {
	vals := map[string]PrivilegedAccessGroupEligibilityScheduleInstanceMemberType{
		"direct": PrivilegedAccessGroupEligibilityScheduleInstanceMemberType_Direct,
		"group":  PrivilegedAccessGroupEligibilityScheduleInstanceMemberType_Group,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupEligibilityScheduleInstanceMemberType(input)
	return &out, nil
}
