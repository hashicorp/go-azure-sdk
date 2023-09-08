package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleAssignmentScheduleRequestAction string

const (
	UnifiedRoleAssignmentScheduleRequestActionadminAssign    UnifiedRoleAssignmentScheduleRequestAction = "AdminAssign"
	UnifiedRoleAssignmentScheduleRequestActionadminExtend    UnifiedRoleAssignmentScheduleRequestAction = "AdminExtend"
	UnifiedRoleAssignmentScheduleRequestActionadminRemove    UnifiedRoleAssignmentScheduleRequestAction = "AdminRemove"
	UnifiedRoleAssignmentScheduleRequestActionadminRenew     UnifiedRoleAssignmentScheduleRequestAction = "AdminRenew"
	UnifiedRoleAssignmentScheduleRequestActionadminUpdate    UnifiedRoleAssignmentScheduleRequestAction = "AdminUpdate"
	UnifiedRoleAssignmentScheduleRequestActionselfActivate   UnifiedRoleAssignmentScheduleRequestAction = "SelfActivate"
	UnifiedRoleAssignmentScheduleRequestActionselfDeactivate UnifiedRoleAssignmentScheduleRequestAction = "SelfDeactivate"
	UnifiedRoleAssignmentScheduleRequestActionselfExtend     UnifiedRoleAssignmentScheduleRequestAction = "SelfExtend"
	UnifiedRoleAssignmentScheduleRequestActionselfRenew      UnifiedRoleAssignmentScheduleRequestAction = "SelfRenew"
)

func PossibleValuesForUnifiedRoleAssignmentScheduleRequestAction() []string {
	return []string{
		string(UnifiedRoleAssignmentScheduleRequestActionadminAssign),
		string(UnifiedRoleAssignmentScheduleRequestActionadminExtend),
		string(UnifiedRoleAssignmentScheduleRequestActionadminRemove),
		string(UnifiedRoleAssignmentScheduleRequestActionadminRenew),
		string(UnifiedRoleAssignmentScheduleRequestActionadminUpdate),
		string(UnifiedRoleAssignmentScheduleRequestActionselfActivate),
		string(UnifiedRoleAssignmentScheduleRequestActionselfDeactivate),
		string(UnifiedRoleAssignmentScheduleRequestActionselfExtend),
		string(UnifiedRoleAssignmentScheduleRequestActionselfRenew),
	}
}

func parseUnifiedRoleAssignmentScheduleRequestAction(input string) (*UnifiedRoleAssignmentScheduleRequestAction, error) {
	vals := map[string]UnifiedRoleAssignmentScheduleRequestAction{
		"adminassign":    UnifiedRoleAssignmentScheduleRequestActionadminAssign,
		"adminextend":    UnifiedRoleAssignmentScheduleRequestActionadminExtend,
		"adminremove":    UnifiedRoleAssignmentScheduleRequestActionadminRemove,
		"adminrenew":     UnifiedRoleAssignmentScheduleRequestActionadminRenew,
		"adminupdate":    UnifiedRoleAssignmentScheduleRequestActionadminUpdate,
		"selfactivate":   UnifiedRoleAssignmentScheduleRequestActionselfActivate,
		"selfdeactivate": UnifiedRoleAssignmentScheduleRequestActionselfDeactivate,
		"selfextend":     UnifiedRoleAssignmentScheduleRequestActionselfExtend,
		"selfrenew":      UnifiedRoleAssignmentScheduleRequestActionselfRenew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnifiedRoleAssignmentScheduleRequestAction(input)
	return &out, nil
}
