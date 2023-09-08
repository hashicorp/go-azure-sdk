package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessScheduleRequestAction string

const (
	PrivilegedAccessScheduleRequestActionadminAssign    PrivilegedAccessScheduleRequestAction = "AdminAssign"
	PrivilegedAccessScheduleRequestActionadminExtend    PrivilegedAccessScheduleRequestAction = "AdminExtend"
	PrivilegedAccessScheduleRequestActionadminRemove    PrivilegedAccessScheduleRequestAction = "AdminRemove"
	PrivilegedAccessScheduleRequestActionadminRenew     PrivilegedAccessScheduleRequestAction = "AdminRenew"
	PrivilegedAccessScheduleRequestActionadminUpdate    PrivilegedAccessScheduleRequestAction = "AdminUpdate"
	PrivilegedAccessScheduleRequestActionselfActivate   PrivilegedAccessScheduleRequestAction = "SelfActivate"
	PrivilegedAccessScheduleRequestActionselfDeactivate PrivilegedAccessScheduleRequestAction = "SelfDeactivate"
	PrivilegedAccessScheduleRequestActionselfExtend     PrivilegedAccessScheduleRequestAction = "SelfExtend"
	PrivilegedAccessScheduleRequestActionselfRenew      PrivilegedAccessScheduleRequestAction = "SelfRenew"
)

func PossibleValuesForPrivilegedAccessScheduleRequestAction() []string {
	return []string{
		string(PrivilegedAccessScheduleRequestActionadminAssign),
		string(PrivilegedAccessScheduleRequestActionadminExtend),
		string(PrivilegedAccessScheduleRequestActionadminRemove),
		string(PrivilegedAccessScheduleRequestActionadminRenew),
		string(PrivilegedAccessScheduleRequestActionadminUpdate),
		string(PrivilegedAccessScheduleRequestActionselfActivate),
		string(PrivilegedAccessScheduleRequestActionselfDeactivate),
		string(PrivilegedAccessScheduleRequestActionselfExtend),
		string(PrivilegedAccessScheduleRequestActionselfRenew),
	}
}

func parsePrivilegedAccessScheduleRequestAction(input string) (*PrivilegedAccessScheduleRequestAction, error) {
	vals := map[string]PrivilegedAccessScheduleRequestAction{
		"adminassign":    PrivilegedAccessScheduleRequestActionadminAssign,
		"adminextend":    PrivilegedAccessScheduleRequestActionadminExtend,
		"adminremove":    PrivilegedAccessScheduleRequestActionadminRemove,
		"adminrenew":     PrivilegedAccessScheduleRequestActionadminRenew,
		"adminupdate":    PrivilegedAccessScheduleRequestActionadminUpdate,
		"selfactivate":   PrivilegedAccessScheduleRequestActionselfActivate,
		"selfdeactivate": PrivilegedAccessScheduleRequestActionselfDeactivate,
		"selfextend":     PrivilegedAccessScheduleRequestActionselfExtend,
		"selfrenew":      PrivilegedAccessScheduleRequestActionselfRenew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessScheduleRequestAction(input)
	return &out, nil
}
