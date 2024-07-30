package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleRequestAction string

const (
	PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminAssign    PrivilegedAccessGroupAssignmentScheduleRequestAction = "adminAssign"
	PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminExtend    PrivilegedAccessGroupAssignmentScheduleRequestAction = "adminExtend"
	PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminRemove    PrivilegedAccessGroupAssignmentScheduleRequestAction = "adminRemove"
	PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminRenew     PrivilegedAccessGroupAssignmentScheduleRequestAction = "adminRenew"
	PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminUpdate    PrivilegedAccessGroupAssignmentScheduleRequestAction = "adminUpdate"
	PrivilegedAccessGroupAssignmentScheduleRequestAction_SelfActivate   PrivilegedAccessGroupAssignmentScheduleRequestAction = "selfActivate"
	PrivilegedAccessGroupAssignmentScheduleRequestAction_SelfDeactivate PrivilegedAccessGroupAssignmentScheduleRequestAction = "selfDeactivate"
	PrivilegedAccessGroupAssignmentScheduleRequestAction_SelfExtend     PrivilegedAccessGroupAssignmentScheduleRequestAction = "selfExtend"
	PrivilegedAccessGroupAssignmentScheduleRequestAction_SelfRenew      PrivilegedAccessGroupAssignmentScheduleRequestAction = "selfRenew"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleRequestAction() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminAssign),
		string(PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminExtend),
		string(PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminRemove),
		string(PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminRenew),
		string(PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminUpdate),
		string(PrivilegedAccessGroupAssignmentScheduleRequestAction_SelfActivate),
		string(PrivilegedAccessGroupAssignmentScheduleRequestAction_SelfDeactivate),
		string(PrivilegedAccessGroupAssignmentScheduleRequestAction_SelfExtend),
		string(PrivilegedAccessGroupAssignmentScheduleRequestAction_SelfRenew),
	}
}

func (s *PrivilegedAccessGroupAssignmentScheduleRequestAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupAssignmentScheduleRequestAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupAssignmentScheduleRequestAction(input string) (*PrivilegedAccessGroupAssignmentScheduleRequestAction, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleRequestAction{
		"adminassign":    PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminAssign,
		"adminextend":    PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminExtend,
		"adminremove":    PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminRemove,
		"adminrenew":     PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminRenew,
		"adminupdate":    PrivilegedAccessGroupAssignmentScheduleRequestAction_AdminUpdate,
		"selfactivate":   PrivilegedAccessGroupAssignmentScheduleRequestAction_SelfActivate,
		"selfdeactivate": PrivilegedAccessGroupAssignmentScheduleRequestAction_SelfDeactivate,
		"selfextend":     PrivilegedAccessGroupAssignmentScheduleRequestAction_SelfExtend,
		"selfrenew":      PrivilegedAccessGroupAssignmentScheduleRequestAction_SelfRenew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleRequestAction(input)
	return &out, nil
}
