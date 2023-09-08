package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleRequestAccessId string

const (
	PrivilegedAccessGroupAssignmentScheduleRequestAccessIdmember PrivilegedAccessGroupAssignmentScheduleRequestAccessId = "Member"
	PrivilegedAccessGroupAssignmentScheduleRequestAccessIdowner  PrivilegedAccessGroupAssignmentScheduleRequestAccessId = "Owner"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleRequestAccessId() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleRequestAccessIdmember),
		string(PrivilegedAccessGroupAssignmentScheduleRequestAccessIdowner),
	}
}

func parsePrivilegedAccessGroupAssignmentScheduleRequestAccessId(input string) (*PrivilegedAccessGroupAssignmentScheduleRequestAccessId, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleRequestAccessId{
		"member": PrivilegedAccessGroupAssignmentScheduleRequestAccessIdmember,
		"owner":  PrivilegedAccessGroupAssignmentScheduleRequestAccessIdowner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleRequestAccessId(input)
	return &out, nil
}
