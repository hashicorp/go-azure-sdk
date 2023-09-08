package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleRequestAction string

const (
	PrivilegedAccessGroupAssignmentScheduleRequestActionadminAssign    PrivilegedAccessGroupAssignmentScheduleRequestAction = "AdminAssign"
	PrivilegedAccessGroupAssignmentScheduleRequestActionadminExtend    PrivilegedAccessGroupAssignmentScheduleRequestAction = "AdminExtend"
	PrivilegedAccessGroupAssignmentScheduleRequestActionadminRemove    PrivilegedAccessGroupAssignmentScheduleRequestAction = "AdminRemove"
	PrivilegedAccessGroupAssignmentScheduleRequestActionadminRenew     PrivilegedAccessGroupAssignmentScheduleRequestAction = "AdminRenew"
	PrivilegedAccessGroupAssignmentScheduleRequestActionadminUpdate    PrivilegedAccessGroupAssignmentScheduleRequestAction = "AdminUpdate"
	PrivilegedAccessGroupAssignmentScheduleRequestActionselfActivate   PrivilegedAccessGroupAssignmentScheduleRequestAction = "SelfActivate"
	PrivilegedAccessGroupAssignmentScheduleRequestActionselfDeactivate PrivilegedAccessGroupAssignmentScheduleRequestAction = "SelfDeactivate"
	PrivilegedAccessGroupAssignmentScheduleRequestActionselfExtend     PrivilegedAccessGroupAssignmentScheduleRequestAction = "SelfExtend"
	PrivilegedAccessGroupAssignmentScheduleRequestActionselfRenew      PrivilegedAccessGroupAssignmentScheduleRequestAction = "SelfRenew"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleRequestAction() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleRequestActionadminAssign),
		string(PrivilegedAccessGroupAssignmentScheduleRequestActionadminExtend),
		string(PrivilegedAccessGroupAssignmentScheduleRequestActionadminRemove),
		string(PrivilegedAccessGroupAssignmentScheduleRequestActionadminRenew),
		string(PrivilegedAccessGroupAssignmentScheduleRequestActionadminUpdate),
		string(PrivilegedAccessGroupAssignmentScheduleRequestActionselfActivate),
		string(PrivilegedAccessGroupAssignmentScheduleRequestActionselfDeactivate),
		string(PrivilegedAccessGroupAssignmentScheduleRequestActionselfExtend),
		string(PrivilegedAccessGroupAssignmentScheduleRequestActionselfRenew),
	}
}

func parsePrivilegedAccessGroupAssignmentScheduleRequestAction(input string) (*PrivilegedAccessGroupAssignmentScheduleRequestAction, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleRequestAction{
		"adminassign":    PrivilegedAccessGroupAssignmentScheduleRequestActionadminAssign,
		"adminextend":    PrivilegedAccessGroupAssignmentScheduleRequestActionadminExtend,
		"adminremove":    PrivilegedAccessGroupAssignmentScheduleRequestActionadminRemove,
		"adminrenew":     PrivilegedAccessGroupAssignmentScheduleRequestActionadminRenew,
		"adminupdate":    PrivilegedAccessGroupAssignmentScheduleRequestActionadminUpdate,
		"selfactivate":   PrivilegedAccessGroupAssignmentScheduleRequestActionselfActivate,
		"selfdeactivate": PrivilegedAccessGroupAssignmentScheduleRequestActionselfDeactivate,
		"selfextend":     PrivilegedAccessGroupAssignmentScheduleRequestActionselfExtend,
		"selfrenew":      PrivilegedAccessGroupAssignmentScheduleRequestActionselfRenew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleRequestAction(input)
	return &out, nil
}
