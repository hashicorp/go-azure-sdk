package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleRequestAccessId string

const (
	PrivilegedAccessGroupEligibilityScheduleRequestAccessId_Member PrivilegedAccessGroupEligibilityScheduleRequestAccessId = "member"
	PrivilegedAccessGroupEligibilityScheduleRequestAccessId_Owner  PrivilegedAccessGroupEligibilityScheduleRequestAccessId = "owner"
)

func PossibleValuesForPrivilegedAccessGroupEligibilityScheduleRequestAccessId() []string {
	return []string{
		string(PrivilegedAccessGroupEligibilityScheduleRequestAccessId_Member),
		string(PrivilegedAccessGroupEligibilityScheduleRequestAccessId_Owner),
	}
}

func (s *PrivilegedAccessGroupEligibilityScheduleRequestAccessId) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupEligibilityScheduleRequestAccessId(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupEligibilityScheduleRequestAccessId(input string) (*PrivilegedAccessGroupEligibilityScheduleRequestAccessId, error) {
	vals := map[string]PrivilegedAccessGroupEligibilityScheduleRequestAccessId{
		"member": PrivilegedAccessGroupEligibilityScheduleRequestAccessId_Member,
		"owner":  PrivilegedAccessGroupEligibilityScheduleRequestAccessId_Owner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupEligibilityScheduleRequestAccessId(input)
	return &out, nil
}
