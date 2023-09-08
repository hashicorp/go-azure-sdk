package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleAccessId string

const (
	PrivilegedAccessGroupAssignmentScheduleAccessIdmember PrivilegedAccessGroupAssignmentScheduleAccessId = "Member"
	PrivilegedAccessGroupAssignmentScheduleAccessIdowner  PrivilegedAccessGroupAssignmentScheduleAccessId = "Owner"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleAccessId() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleAccessIdmember),
		string(PrivilegedAccessGroupAssignmentScheduleAccessIdowner),
	}
}

func parsePrivilegedAccessGroupAssignmentScheduleAccessId(input string) (*PrivilegedAccessGroupAssignmentScheduleAccessId, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleAccessId{
		"member": PrivilegedAccessGroupAssignmentScheduleAccessIdmember,
		"owner":  PrivilegedAccessGroupAssignmentScheduleAccessIdowner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleAccessId(input)
	return &out, nil
}
