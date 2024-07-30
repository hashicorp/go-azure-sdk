package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteLockActionResultActionState string

const (
	RemoteLockActionResultActionState_Active       RemoteLockActionResultActionState = "active"
	RemoteLockActionResultActionState_Canceled     RemoteLockActionResultActionState = "canceled"
	RemoteLockActionResultActionState_Done         RemoteLockActionResultActionState = "done"
	RemoteLockActionResultActionState_Failed       RemoteLockActionResultActionState = "failed"
	RemoteLockActionResultActionState_None         RemoteLockActionResultActionState = "none"
	RemoteLockActionResultActionState_NotSupported RemoteLockActionResultActionState = "notSupported"
	RemoteLockActionResultActionState_Pending      RemoteLockActionResultActionState = "pending"
)

func PossibleValuesForRemoteLockActionResultActionState() []string {
	return []string{
		string(RemoteLockActionResultActionState_Active),
		string(RemoteLockActionResultActionState_Canceled),
		string(RemoteLockActionResultActionState_Done),
		string(RemoteLockActionResultActionState_Failed),
		string(RemoteLockActionResultActionState_None),
		string(RemoteLockActionResultActionState_NotSupported),
		string(RemoteLockActionResultActionState_Pending),
	}
}

func (s *RemoteLockActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemoteLockActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemoteLockActionResultActionState(input string) (*RemoteLockActionResultActionState, error) {
	vals := map[string]RemoteLockActionResultActionState{
		"active":       RemoteLockActionResultActionState_Active,
		"canceled":     RemoteLockActionResultActionState_Canceled,
		"done":         RemoteLockActionResultActionState_Done,
		"failed":       RemoteLockActionResultActionState_Failed,
		"none":         RemoteLockActionResultActionState_None,
		"notsupported": RemoteLockActionResultActionState_NotSupported,
		"pending":      RemoteLockActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteLockActionResultActionState(input)
	return &out, nil
}
