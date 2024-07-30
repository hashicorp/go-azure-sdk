package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleInstanceAccessId string

const (
	PrivilegedAccessGroupAssignmentScheduleInstanceAccessId_Member PrivilegedAccessGroupAssignmentScheduleInstanceAccessId = "member"
	PrivilegedAccessGroupAssignmentScheduleInstanceAccessId_Owner  PrivilegedAccessGroupAssignmentScheduleInstanceAccessId = "owner"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleInstanceAccessId() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleInstanceAccessId_Member),
		string(PrivilegedAccessGroupAssignmentScheduleInstanceAccessId_Owner),
	}
}

func (s *PrivilegedAccessGroupAssignmentScheduleInstanceAccessId) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupAssignmentScheduleInstanceAccessId(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupAssignmentScheduleInstanceAccessId(input string) (*PrivilegedAccessGroupAssignmentScheduleInstanceAccessId, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleInstanceAccessId{
		"member": PrivilegedAccessGroupAssignmentScheduleInstanceAccessId_Member,
		"owner":  PrivilegedAccessGroupAssignmentScheduleInstanceAccessId_Owner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleInstanceAccessId(input)
	return &out, nil
}
