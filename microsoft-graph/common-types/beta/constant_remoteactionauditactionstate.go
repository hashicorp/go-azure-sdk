package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteActionAuditActionState string

const (
	RemoteActionAuditActionState_Active       RemoteActionAuditActionState = "active"
	RemoteActionAuditActionState_Canceled     RemoteActionAuditActionState = "canceled"
	RemoteActionAuditActionState_Done         RemoteActionAuditActionState = "done"
	RemoteActionAuditActionState_Failed       RemoteActionAuditActionState = "failed"
	RemoteActionAuditActionState_None         RemoteActionAuditActionState = "none"
	RemoteActionAuditActionState_NotSupported RemoteActionAuditActionState = "notSupported"
	RemoteActionAuditActionState_Pending      RemoteActionAuditActionState = "pending"
)

func PossibleValuesForRemoteActionAuditActionState() []string {
	return []string{
		string(RemoteActionAuditActionState_Active),
		string(RemoteActionAuditActionState_Canceled),
		string(RemoteActionAuditActionState_Done),
		string(RemoteActionAuditActionState_Failed),
		string(RemoteActionAuditActionState_None),
		string(RemoteActionAuditActionState_NotSupported),
		string(RemoteActionAuditActionState_Pending),
	}
}

func (s *RemoteActionAuditActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemoteActionAuditActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemoteActionAuditActionState(input string) (*RemoteActionAuditActionState, error) {
	vals := map[string]RemoteActionAuditActionState{
		"active":       RemoteActionAuditActionState_Active,
		"canceled":     RemoteActionAuditActionState_Canceled,
		"done":         RemoteActionAuditActionState_Done,
		"failed":       RemoteActionAuditActionState_Failed,
		"none":         RemoteActionAuditActionState_None,
		"notsupported": RemoteActionAuditActionState_NotSupported,
		"pending":      RemoteActionAuditActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteActionAuditActionState(input)
	return &out, nil
}
