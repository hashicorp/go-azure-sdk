package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleInstanceMemberType string

const (
	PrivilegedAccessGroupAssignmentScheduleInstanceMemberType_Direct PrivilegedAccessGroupAssignmentScheduleInstanceMemberType = "direct"
	PrivilegedAccessGroupAssignmentScheduleInstanceMemberType_Group  PrivilegedAccessGroupAssignmentScheduleInstanceMemberType = "group"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleInstanceMemberType() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleInstanceMemberType_Direct),
		string(PrivilegedAccessGroupAssignmentScheduleInstanceMemberType_Group),
	}
}

func (s *PrivilegedAccessGroupAssignmentScheduleInstanceMemberType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupAssignmentScheduleInstanceMemberType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupAssignmentScheduleInstanceMemberType(input string) (*PrivilegedAccessGroupAssignmentScheduleInstanceMemberType, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleInstanceMemberType{
		"direct": PrivilegedAccessGroupAssignmentScheduleInstanceMemberType_Direct,
		"group":  PrivilegedAccessGroupAssignmentScheduleInstanceMemberType_Group,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleInstanceMemberType(input)
	return &out, nil
}
