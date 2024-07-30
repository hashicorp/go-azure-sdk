package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType string

const (
	PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType_Activated PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType = "activated"
	PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType_Assigned  PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType = "assigned"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType_Activated),
		string(PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType_Assigned),
	}
}

func (s *PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType(input string) (*PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType{
		"activated": PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType_Activated,
		"assigned":  PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType_Assigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType(input)
	return &out, nil
}
