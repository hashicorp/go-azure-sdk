package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleAccessId string

const (
	PrivilegedAccessGroupEligibilityScheduleAccessId_Member PrivilegedAccessGroupEligibilityScheduleAccessId = "member"
	PrivilegedAccessGroupEligibilityScheduleAccessId_Owner  PrivilegedAccessGroupEligibilityScheduleAccessId = "owner"
)

func PossibleValuesForPrivilegedAccessGroupEligibilityScheduleAccessId() []string {
	return []string{
		string(PrivilegedAccessGroupEligibilityScheduleAccessId_Member),
		string(PrivilegedAccessGroupEligibilityScheduleAccessId_Owner),
	}
}

func (s *PrivilegedAccessGroupEligibilityScheduleAccessId) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupEligibilityScheduleAccessId(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupEligibilityScheduleAccessId(input string) (*PrivilegedAccessGroupEligibilityScheduleAccessId, error) {
	vals := map[string]PrivilegedAccessGroupEligibilityScheduleAccessId{
		"member": PrivilegedAccessGroupEligibilityScheduleAccessId_Member,
		"owner":  PrivilegedAccessGroupEligibilityScheduleAccessId_Owner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupEligibilityScheduleAccessId(input)
	return &out, nil
}
