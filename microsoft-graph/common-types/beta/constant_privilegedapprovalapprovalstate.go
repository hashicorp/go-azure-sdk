package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedApprovalApprovalState string

const (
	PrivilegedApprovalApprovalState_Aborted  PrivilegedApprovalApprovalState = "aborted"
	PrivilegedApprovalApprovalState_Approved PrivilegedApprovalApprovalState = "approved"
	PrivilegedApprovalApprovalState_Canceled PrivilegedApprovalApprovalState = "canceled"
	PrivilegedApprovalApprovalState_Denied   PrivilegedApprovalApprovalState = "denied"
	PrivilegedApprovalApprovalState_Pending  PrivilegedApprovalApprovalState = "pending"
)

func PossibleValuesForPrivilegedApprovalApprovalState() []string {
	return []string{
		string(PrivilegedApprovalApprovalState_Aborted),
		string(PrivilegedApprovalApprovalState_Approved),
		string(PrivilegedApprovalApprovalState_Canceled),
		string(PrivilegedApprovalApprovalState_Denied),
		string(PrivilegedApprovalApprovalState_Pending),
	}
}

func (s *PrivilegedApprovalApprovalState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedApprovalApprovalState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedApprovalApprovalState(input string) (*PrivilegedApprovalApprovalState, error) {
	vals := map[string]PrivilegedApprovalApprovalState{
		"aborted":  PrivilegedApprovalApprovalState_Aborted,
		"approved": PrivilegedApprovalApprovalState_Approved,
		"canceled": PrivilegedApprovalApprovalState_Canceled,
		"denied":   PrivilegedApprovalApprovalState_Denied,
		"pending":  PrivilegedApprovalApprovalState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedApprovalApprovalState(input)
	return &out, nil
}
