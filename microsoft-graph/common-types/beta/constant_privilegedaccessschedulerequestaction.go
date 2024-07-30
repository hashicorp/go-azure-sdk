package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessScheduleRequestAction string

const (
	PrivilegedAccessScheduleRequestAction_AdminAssign    PrivilegedAccessScheduleRequestAction = "adminAssign"
	PrivilegedAccessScheduleRequestAction_AdminExtend    PrivilegedAccessScheduleRequestAction = "adminExtend"
	PrivilegedAccessScheduleRequestAction_AdminRemove    PrivilegedAccessScheduleRequestAction = "adminRemove"
	PrivilegedAccessScheduleRequestAction_AdminRenew     PrivilegedAccessScheduleRequestAction = "adminRenew"
	PrivilegedAccessScheduleRequestAction_AdminUpdate    PrivilegedAccessScheduleRequestAction = "adminUpdate"
	PrivilegedAccessScheduleRequestAction_SelfActivate   PrivilegedAccessScheduleRequestAction = "selfActivate"
	PrivilegedAccessScheduleRequestAction_SelfDeactivate PrivilegedAccessScheduleRequestAction = "selfDeactivate"
	PrivilegedAccessScheduleRequestAction_SelfExtend     PrivilegedAccessScheduleRequestAction = "selfExtend"
	PrivilegedAccessScheduleRequestAction_SelfRenew      PrivilegedAccessScheduleRequestAction = "selfRenew"
)

func PossibleValuesForPrivilegedAccessScheduleRequestAction() []string {
	return []string{
		string(PrivilegedAccessScheduleRequestAction_AdminAssign),
		string(PrivilegedAccessScheduleRequestAction_AdminExtend),
		string(PrivilegedAccessScheduleRequestAction_AdminRemove),
		string(PrivilegedAccessScheduleRequestAction_AdminRenew),
		string(PrivilegedAccessScheduleRequestAction_AdminUpdate),
		string(PrivilegedAccessScheduleRequestAction_SelfActivate),
		string(PrivilegedAccessScheduleRequestAction_SelfDeactivate),
		string(PrivilegedAccessScheduleRequestAction_SelfExtend),
		string(PrivilegedAccessScheduleRequestAction_SelfRenew),
	}
}

func (s *PrivilegedAccessScheduleRequestAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessScheduleRequestAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessScheduleRequestAction(input string) (*PrivilegedAccessScheduleRequestAction, error) {
	vals := map[string]PrivilegedAccessScheduleRequestAction{
		"adminassign":    PrivilegedAccessScheduleRequestAction_AdminAssign,
		"adminextend":    PrivilegedAccessScheduleRequestAction_AdminExtend,
		"adminremove":    PrivilegedAccessScheduleRequestAction_AdminRemove,
		"adminrenew":     PrivilegedAccessScheduleRequestAction_AdminRenew,
		"adminupdate":    PrivilegedAccessScheduleRequestAction_AdminUpdate,
		"selfactivate":   PrivilegedAccessScheduleRequestAction_SelfActivate,
		"selfdeactivate": PrivilegedAccessScheduleRequestAction_SelfDeactivate,
		"selfextend":     PrivilegedAccessScheduleRequestAction_SelfExtend,
		"selfrenew":      PrivilegedAccessScheduleRequestAction_SelfRenew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessScheduleRequestAction(input)
	return &out, nil
}
