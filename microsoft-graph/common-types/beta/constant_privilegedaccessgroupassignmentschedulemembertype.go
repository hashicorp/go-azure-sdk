package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleMemberType string

const (
	PrivilegedAccessGroupAssignmentScheduleMemberType_Direct PrivilegedAccessGroupAssignmentScheduleMemberType = "direct"
	PrivilegedAccessGroupAssignmentScheduleMemberType_Group  PrivilegedAccessGroupAssignmentScheduleMemberType = "group"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleMemberType() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleMemberType_Direct),
		string(PrivilegedAccessGroupAssignmentScheduleMemberType_Group),
	}
}

func (s *PrivilegedAccessGroupAssignmentScheduleMemberType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupAssignmentScheduleMemberType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupAssignmentScheduleMemberType(input string) (*PrivilegedAccessGroupAssignmentScheduleMemberType, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleMemberType{
		"direct": PrivilegedAccessGroupAssignmentScheduleMemberType_Direct,
		"group":  PrivilegedAccessGroupAssignmentScheduleMemberType_Group,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleMemberType(input)
	return &out, nil
}
