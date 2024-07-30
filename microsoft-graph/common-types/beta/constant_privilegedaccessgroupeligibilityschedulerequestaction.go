package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleRequestAction string

const (
	PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminAssign    PrivilegedAccessGroupEligibilityScheduleRequestAction = "adminAssign"
	PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminExtend    PrivilegedAccessGroupEligibilityScheduleRequestAction = "adminExtend"
	PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminRemove    PrivilegedAccessGroupEligibilityScheduleRequestAction = "adminRemove"
	PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminRenew     PrivilegedAccessGroupEligibilityScheduleRequestAction = "adminRenew"
	PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminUpdate    PrivilegedAccessGroupEligibilityScheduleRequestAction = "adminUpdate"
	PrivilegedAccessGroupEligibilityScheduleRequestAction_SelfActivate   PrivilegedAccessGroupEligibilityScheduleRequestAction = "selfActivate"
	PrivilegedAccessGroupEligibilityScheduleRequestAction_SelfDeactivate PrivilegedAccessGroupEligibilityScheduleRequestAction = "selfDeactivate"
	PrivilegedAccessGroupEligibilityScheduleRequestAction_SelfExtend     PrivilegedAccessGroupEligibilityScheduleRequestAction = "selfExtend"
	PrivilegedAccessGroupEligibilityScheduleRequestAction_SelfRenew      PrivilegedAccessGroupEligibilityScheduleRequestAction = "selfRenew"
)

func PossibleValuesForPrivilegedAccessGroupEligibilityScheduleRequestAction() []string {
	return []string{
		string(PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminAssign),
		string(PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminExtend),
		string(PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminRemove),
		string(PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminRenew),
		string(PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminUpdate),
		string(PrivilegedAccessGroupEligibilityScheduleRequestAction_SelfActivate),
		string(PrivilegedAccessGroupEligibilityScheduleRequestAction_SelfDeactivate),
		string(PrivilegedAccessGroupEligibilityScheduleRequestAction_SelfExtend),
		string(PrivilegedAccessGroupEligibilityScheduleRequestAction_SelfRenew),
	}
}

func (s *PrivilegedAccessGroupEligibilityScheduleRequestAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupEligibilityScheduleRequestAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupEligibilityScheduleRequestAction(input string) (*PrivilegedAccessGroupEligibilityScheduleRequestAction, error) {
	vals := map[string]PrivilegedAccessGroupEligibilityScheduleRequestAction{
		"adminassign":    PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminAssign,
		"adminextend":    PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminExtend,
		"adminremove":    PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminRemove,
		"adminrenew":     PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminRenew,
		"adminupdate":    PrivilegedAccessGroupEligibilityScheduleRequestAction_AdminUpdate,
		"selfactivate":   PrivilegedAccessGroupEligibilityScheduleRequestAction_SelfActivate,
		"selfdeactivate": PrivilegedAccessGroupEligibilityScheduleRequestAction_SelfDeactivate,
		"selfextend":     PrivilegedAccessGroupEligibilityScheduleRequestAction_SelfExtend,
		"selfrenew":      PrivilegedAccessGroupEligibilityScheduleRequestAction_SelfRenew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupEligibilityScheduleRequestAction(input)
	return &out, nil
}
