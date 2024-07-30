package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleAssignmentType string

const (
	PrivilegedAccessGroupAssignmentScheduleAssignmentType_Activated PrivilegedAccessGroupAssignmentScheduleAssignmentType = "activated"
	PrivilegedAccessGroupAssignmentScheduleAssignmentType_Assigned  PrivilegedAccessGroupAssignmentScheduleAssignmentType = "assigned"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleAssignmentType() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleAssignmentType_Activated),
		string(PrivilegedAccessGroupAssignmentScheduleAssignmentType_Assigned),
	}
}

func (s *PrivilegedAccessGroupAssignmentScheduleAssignmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupAssignmentScheduleAssignmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupAssignmentScheduleAssignmentType(input string) (*PrivilegedAccessGroupAssignmentScheduleAssignmentType, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleAssignmentType{
		"activated": PrivilegedAccessGroupAssignmentScheduleAssignmentType_Activated,
		"assigned":  PrivilegedAccessGroupAssignmentScheduleAssignmentType_Assigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleAssignmentType(input)
	return &out, nil
}
