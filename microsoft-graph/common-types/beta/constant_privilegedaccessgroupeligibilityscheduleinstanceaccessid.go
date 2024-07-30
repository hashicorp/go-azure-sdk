package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleInstanceAccessId string

const (
	PrivilegedAccessGroupEligibilityScheduleInstanceAccessId_Member PrivilegedAccessGroupEligibilityScheduleInstanceAccessId = "member"
	PrivilegedAccessGroupEligibilityScheduleInstanceAccessId_Owner  PrivilegedAccessGroupEligibilityScheduleInstanceAccessId = "owner"
)

func PossibleValuesForPrivilegedAccessGroupEligibilityScheduleInstanceAccessId() []string {
	return []string{
		string(PrivilegedAccessGroupEligibilityScheduleInstanceAccessId_Member),
		string(PrivilegedAccessGroupEligibilityScheduleInstanceAccessId_Owner),
	}
}

func (s *PrivilegedAccessGroupEligibilityScheduleInstanceAccessId) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupEligibilityScheduleInstanceAccessId(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupEligibilityScheduleInstanceAccessId(input string) (*PrivilegedAccessGroupEligibilityScheduleInstanceAccessId, error) {
	vals := map[string]PrivilegedAccessGroupEligibilityScheduleInstanceAccessId{
		"member": PrivilegedAccessGroupEligibilityScheduleInstanceAccessId_Member,
		"owner":  PrivilegedAccessGroupEligibilityScheduleInstanceAccessId_Owner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupEligibilityScheduleInstanceAccessId(input)
	return &out, nil
}
