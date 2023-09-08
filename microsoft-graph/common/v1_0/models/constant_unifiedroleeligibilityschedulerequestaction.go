package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleEligibilityScheduleRequestAction string

const (
	UnifiedRoleEligibilityScheduleRequestActionadminAssign    UnifiedRoleEligibilityScheduleRequestAction = "AdminAssign"
	UnifiedRoleEligibilityScheduleRequestActionadminExtend    UnifiedRoleEligibilityScheduleRequestAction = "AdminExtend"
	UnifiedRoleEligibilityScheduleRequestActionadminRemove    UnifiedRoleEligibilityScheduleRequestAction = "AdminRemove"
	UnifiedRoleEligibilityScheduleRequestActionadminRenew     UnifiedRoleEligibilityScheduleRequestAction = "AdminRenew"
	UnifiedRoleEligibilityScheduleRequestActionadminUpdate    UnifiedRoleEligibilityScheduleRequestAction = "AdminUpdate"
	UnifiedRoleEligibilityScheduleRequestActionselfActivate   UnifiedRoleEligibilityScheduleRequestAction = "SelfActivate"
	UnifiedRoleEligibilityScheduleRequestActionselfDeactivate UnifiedRoleEligibilityScheduleRequestAction = "SelfDeactivate"
	UnifiedRoleEligibilityScheduleRequestActionselfExtend     UnifiedRoleEligibilityScheduleRequestAction = "SelfExtend"
	UnifiedRoleEligibilityScheduleRequestActionselfRenew      UnifiedRoleEligibilityScheduleRequestAction = "SelfRenew"
)

func PossibleValuesForUnifiedRoleEligibilityScheduleRequestAction() []string {
	return []string{
		string(UnifiedRoleEligibilityScheduleRequestActionadminAssign),
		string(UnifiedRoleEligibilityScheduleRequestActionadminExtend),
		string(UnifiedRoleEligibilityScheduleRequestActionadminRemove),
		string(UnifiedRoleEligibilityScheduleRequestActionadminRenew),
		string(UnifiedRoleEligibilityScheduleRequestActionadminUpdate),
		string(UnifiedRoleEligibilityScheduleRequestActionselfActivate),
		string(UnifiedRoleEligibilityScheduleRequestActionselfDeactivate),
		string(UnifiedRoleEligibilityScheduleRequestActionselfExtend),
		string(UnifiedRoleEligibilityScheduleRequestActionselfRenew),
	}
}

func parseUnifiedRoleEligibilityScheduleRequestAction(input string) (*UnifiedRoleEligibilityScheduleRequestAction, error) {
	vals := map[string]UnifiedRoleEligibilityScheduleRequestAction{
		"adminassign":    UnifiedRoleEligibilityScheduleRequestActionadminAssign,
		"adminextend":    UnifiedRoleEligibilityScheduleRequestActionadminExtend,
		"adminremove":    UnifiedRoleEligibilityScheduleRequestActionadminRemove,
		"adminrenew":     UnifiedRoleEligibilityScheduleRequestActionadminRenew,
		"adminupdate":    UnifiedRoleEligibilityScheduleRequestActionadminUpdate,
		"selfactivate":   UnifiedRoleEligibilityScheduleRequestActionselfActivate,
		"selfdeactivate": UnifiedRoleEligibilityScheduleRequestActionselfDeactivate,
		"selfextend":     UnifiedRoleEligibilityScheduleRequestActionselfExtend,
		"selfrenew":      UnifiedRoleEligibilityScheduleRequestActionselfRenew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnifiedRoleEligibilityScheduleRequestAction(input)
	return &out, nil
}
