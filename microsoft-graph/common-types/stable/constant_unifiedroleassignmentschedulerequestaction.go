package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleAssignmentScheduleRequestAction string

const (
	UnifiedRoleAssignmentScheduleRequestAction_AdminAssign    UnifiedRoleAssignmentScheduleRequestAction = "adminAssign"
	UnifiedRoleAssignmentScheduleRequestAction_AdminExtend    UnifiedRoleAssignmentScheduleRequestAction = "adminExtend"
	UnifiedRoleAssignmentScheduleRequestAction_AdminRemove    UnifiedRoleAssignmentScheduleRequestAction = "adminRemove"
	UnifiedRoleAssignmentScheduleRequestAction_AdminRenew     UnifiedRoleAssignmentScheduleRequestAction = "adminRenew"
	UnifiedRoleAssignmentScheduleRequestAction_AdminUpdate    UnifiedRoleAssignmentScheduleRequestAction = "adminUpdate"
	UnifiedRoleAssignmentScheduleRequestAction_SelfActivate   UnifiedRoleAssignmentScheduleRequestAction = "selfActivate"
	UnifiedRoleAssignmentScheduleRequestAction_SelfDeactivate UnifiedRoleAssignmentScheduleRequestAction = "selfDeactivate"
	UnifiedRoleAssignmentScheduleRequestAction_SelfExtend     UnifiedRoleAssignmentScheduleRequestAction = "selfExtend"
	UnifiedRoleAssignmentScheduleRequestAction_SelfRenew      UnifiedRoleAssignmentScheduleRequestAction = "selfRenew"
)

func PossibleValuesForUnifiedRoleAssignmentScheduleRequestAction() []string {
	return []string{
		string(UnifiedRoleAssignmentScheduleRequestAction_AdminAssign),
		string(UnifiedRoleAssignmentScheduleRequestAction_AdminExtend),
		string(UnifiedRoleAssignmentScheduleRequestAction_AdminRemove),
		string(UnifiedRoleAssignmentScheduleRequestAction_AdminRenew),
		string(UnifiedRoleAssignmentScheduleRequestAction_AdminUpdate),
		string(UnifiedRoleAssignmentScheduleRequestAction_SelfActivate),
		string(UnifiedRoleAssignmentScheduleRequestAction_SelfDeactivate),
		string(UnifiedRoleAssignmentScheduleRequestAction_SelfExtend),
		string(UnifiedRoleAssignmentScheduleRequestAction_SelfRenew),
	}
}

func (s *UnifiedRoleAssignmentScheduleRequestAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnifiedRoleAssignmentScheduleRequestAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnifiedRoleAssignmentScheduleRequestAction(input string) (*UnifiedRoleAssignmentScheduleRequestAction, error) {
	vals := map[string]UnifiedRoleAssignmentScheduleRequestAction{
		"adminassign":    UnifiedRoleAssignmentScheduleRequestAction_AdminAssign,
		"adminextend":    UnifiedRoleAssignmentScheduleRequestAction_AdminExtend,
		"adminremove":    UnifiedRoleAssignmentScheduleRequestAction_AdminRemove,
		"adminrenew":     UnifiedRoleAssignmentScheduleRequestAction_AdminRenew,
		"adminupdate":    UnifiedRoleAssignmentScheduleRequestAction_AdminUpdate,
		"selfactivate":   UnifiedRoleAssignmentScheduleRequestAction_SelfActivate,
		"selfdeactivate": UnifiedRoleAssignmentScheduleRequestAction_SelfDeactivate,
		"selfextend":     UnifiedRoleAssignmentScheduleRequestAction_SelfExtend,
		"selfrenew":      UnifiedRoleAssignmentScheduleRequestAction_SelfRenew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnifiedRoleAssignmentScheduleRequestAction(input)
	return &out, nil
}
