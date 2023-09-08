package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteLockActionResultActionState string

const (
	RemoteLockActionResultActionStateactive       RemoteLockActionResultActionState = "Active"
	RemoteLockActionResultActionStatecanceled     RemoteLockActionResultActionState = "Canceled"
	RemoteLockActionResultActionStatedone         RemoteLockActionResultActionState = "Done"
	RemoteLockActionResultActionStatefailed       RemoteLockActionResultActionState = "Failed"
	RemoteLockActionResultActionStatenone         RemoteLockActionResultActionState = "None"
	RemoteLockActionResultActionStatenotSupported RemoteLockActionResultActionState = "NotSupported"
	RemoteLockActionResultActionStatepending      RemoteLockActionResultActionState = "Pending"
)

func PossibleValuesForRemoteLockActionResultActionState() []string {
	return []string{
		string(RemoteLockActionResultActionStateactive),
		string(RemoteLockActionResultActionStatecanceled),
		string(RemoteLockActionResultActionStatedone),
		string(RemoteLockActionResultActionStatefailed),
		string(RemoteLockActionResultActionStatenone),
		string(RemoteLockActionResultActionStatenotSupported),
		string(RemoteLockActionResultActionStatepending),
	}
}

func parseRemoteLockActionResultActionState(input string) (*RemoteLockActionResultActionState, error) {
	vals := map[string]RemoteLockActionResultActionState{
		"active":       RemoteLockActionResultActionStateactive,
		"canceled":     RemoteLockActionResultActionStatecanceled,
		"done":         RemoteLockActionResultActionStatedone,
		"failed":       RemoteLockActionResultActionStatefailed,
		"none":         RemoteLockActionResultActionStatenone,
		"notsupported": RemoteLockActionResultActionStatenotSupported,
		"pending":      RemoteLockActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteLockActionResultActionState(input)
	return &out, nil
}
