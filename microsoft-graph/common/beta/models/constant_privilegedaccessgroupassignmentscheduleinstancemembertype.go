package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleInstanceMemberType string

const (
	PrivilegedAccessGroupAssignmentScheduleInstanceMemberTypedirect PrivilegedAccessGroupAssignmentScheduleInstanceMemberType = "Direct"
	PrivilegedAccessGroupAssignmentScheduleInstanceMemberTypegroup  PrivilegedAccessGroupAssignmentScheduleInstanceMemberType = "Group"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleInstanceMemberType() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleInstanceMemberTypedirect),
		string(PrivilegedAccessGroupAssignmentScheduleInstanceMemberTypegroup),
	}
}

func parsePrivilegedAccessGroupAssignmentScheduleInstanceMemberType(input string) (*PrivilegedAccessGroupAssignmentScheduleInstanceMemberType, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleInstanceMemberType{
		"direct": PrivilegedAccessGroupAssignmentScheduleInstanceMemberTypedirect,
		"group":  PrivilegedAccessGroupAssignmentScheduleInstanceMemberTypegroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleInstanceMemberType(input)
	return &out, nil
}
