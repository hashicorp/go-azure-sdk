package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleMemberType string

const (
	PrivilegedAccessGroupAssignmentScheduleMemberTypedirect PrivilegedAccessGroupAssignmentScheduleMemberType = "Direct"
	PrivilegedAccessGroupAssignmentScheduleMemberTypegroup  PrivilegedAccessGroupAssignmentScheduleMemberType = "Group"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleMemberType() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleMemberTypedirect),
		string(PrivilegedAccessGroupAssignmentScheduleMemberTypegroup),
	}
}

func parsePrivilegedAccessGroupAssignmentScheduleMemberType(input string) (*PrivilegedAccessGroupAssignmentScheduleMemberType, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleMemberType{
		"direct": PrivilegedAccessGroupAssignmentScheduleMemberTypedirect,
		"group":  PrivilegedAccessGroupAssignmentScheduleMemberTypegroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleMemberType(input)
	return &out, nil
}
