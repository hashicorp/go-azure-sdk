package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleEligibilityScheduleRequestAction string

const (
	UnifiedRoleEligibilityScheduleRequestAction_AdminAssign    UnifiedRoleEligibilityScheduleRequestAction = "adminAssign"
	UnifiedRoleEligibilityScheduleRequestAction_AdminExtend    UnifiedRoleEligibilityScheduleRequestAction = "adminExtend"
	UnifiedRoleEligibilityScheduleRequestAction_AdminRemove    UnifiedRoleEligibilityScheduleRequestAction = "adminRemove"
	UnifiedRoleEligibilityScheduleRequestAction_AdminRenew     UnifiedRoleEligibilityScheduleRequestAction = "adminRenew"
	UnifiedRoleEligibilityScheduleRequestAction_AdminUpdate    UnifiedRoleEligibilityScheduleRequestAction = "adminUpdate"
	UnifiedRoleEligibilityScheduleRequestAction_SelfActivate   UnifiedRoleEligibilityScheduleRequestAction = "selfActivate"
	UnifiedRoleEligibilityScheduleRequestAction_SelfDeactivate UnifiedRoleEligibilityScheduleRequestAction = "selfDeactivate"
	UnifiedRoleEligibilityScheduleRequestAction_SelfExtend     UnifiedRoleEligibilityScheduleRequestAction = "selfExtend"
	UnifiedRoleEligibilityScheduleRequestAction_SelfRenew      UnifiedRoleEligibilityScheduleRequestAction = "selfRenew"
)

func PossibleValuesForUnifiedRoleEligibilityScheduleRequestAction() []string {
	return []string{
		string(UnifiedRoleEligibilityScheduleRequestAction_AdminAssign),
		string(UnifiedRoleEligibilityScheduleRequestAction_AdminExtend),
		string(UnifiedRoleEligibilityScheduleRequestAction_AdminRemove),
		string(UnifiedRoleEligibilityScheduleRequestAction_AdminRenew),
		string(UnifiedRoleEligibilityScheduleRequestAction_AdminUpdate),
		string(UnifiedRoleEligibilityScheduleRequestAction_SelfActivate),
		string(UnifiedRoleEligibilityScheduleRequestAction_SelfDeactivate),
		string(UnifiedRoleEligibilityScheduleRequestAction_SelfExtend),
		string(UnifiedRoleEligibilityScheduleRequestAction_SelfRenew),
	}
}

func (s *UnifiedRoleEligibilityScheduleRequestAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnifiedRoleEligibilityScheduleRequestAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnifiedRoleEligibilityScheduleRequestAction(input string) (*UnifiedRoleEligibilityScheduleRequestAction, error) {
	vals := map[string]UnifiedRoleEligibilityScheduleRequestAction{
		"adminassign":    UnifiedRoleEligibilityScheduleRequestAction_AdminAssign,
		"adminextend":    UnifiedRoleEligibilityScheduleRequestAction_AdminExtend,
		"adminremove":    UnifiedRoleEligibilityScheduleRequestAction_AdminRemove,
		"adminrenew":     UnifiedRoleEligibilityScheduleRequestAction_AdminRenew,
		"adminupdate":    UnifiedRoleEligibilityScheduleRequestAction_AdminUpdate,
		"selfactivate":   UnifiedRoleEligibilityScheduleRequestAction_SelfActivate,
		"selfdeactivate": UnifiedRoleEligibilityScheduleRequestAction_SelfDeactivate,
		"selfextend":     UnifiedRoleEligibilityScheduleRequestAction_SelfExtend,
		"selfrenew":      UnifiedRoleEligibilityScheduleRequestAction_SelfRenew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnifiedRoleEligibilityScheduleRequestAction(input)
	return &out, nil
}
