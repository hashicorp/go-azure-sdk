package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleRequestAccessId string

const (
	PrivilegedAccessGroupAssignmentScheduleRequestAccessId_Member PrivilegedAccessGroupAssignmentScheduleRequestAccessId = "member"
	PrivilegedAccessGroupAssignmentScheduleRequestAccessId_Owner  PrivilegedAccessGroupAssignmentScheduleRequestAccessId = "owner"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleRequestAccessId() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleRequestAccessId_Member),
		string(PrivilegedAccessGroupAssignmentScheduleRequestAccessId_Owner),
	}
}

func (s *PrivilegedAccessGroupAssignmentScheduleRequestAccessId) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupAssignmentScheduleRequestAccessId(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupAssignmentScheduleRequestAccessId(input string) (*PrivilegedAccessGroupAssignmentScheduleRequestAccessId, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleRequestAccessId{
		"member": PrivilegedAccessGroupAssignmentScheduleRequestAccessId_Member,
		"owner":  PrivilegedAccessGroupAssignmentScheduleRequestAccessId_Owner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleRequestAccessId(input)
	return &out, nil
}
