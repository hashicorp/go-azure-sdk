package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleInstanceAccessId string

const (
	PrivilegedAccessGroupAssignmentScheduleInstanceAccessIdmember PrivilegedAccessGroupAssignmentScheduleInstanceAccessId = "Member"
	PrivilegedAccessGroupAssignmentScheduleInstanceAccessIdowner  PrivilegedAccessGroupAssignmentScheduleInstanceAccessId = "Owner"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleInstanceAccessId() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleInstanceAccessIdmember),
		string(PrivilegedAccessGroupAssignmentScheduleInstanceAccessIdowner),
	}
}

func parsePrivilegedAccessGroupAssignmentScheduleInstanceAccessId(input string) (*PrivilegedAccessGroupAssignmentScheduleInstanceAccessId, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleInstanceAccessId{
		"member": PrivilegedAccessGroupAssignmentScheduleInstanceAccessIdmember,
		"owner":  PrivilegedAccessGroupAssignmentScheduleInstanceAccessIdowner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleInstanceAccessId(input)
	return &out, nil
}
