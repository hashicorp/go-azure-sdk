package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenActionResultActionState string

const (
	VppTokenActionResultActionStateactive       VppTokenActionResultActionState = "Active"
	VppTokenActionResultActionStatecanceled     VppTokenActionResultActionState = "Canceled"
	VppTokenActionResultActionStatedone         VppTokenActionResultActionState = "Done"
	VppTokenActionResultActionStatefailed       VppTokenActionResultActionState = "Failed"
	VppTokenActionResultActionStatenone         VppTokenActionResultActionState = "None"
	VppTokenActionResultActionStatenotSupported VppTokenActionResultActionState = "NotSupported"
	VppTokenActionResultActionStatepending      VppTokenActionResultActionState = "Pending"
)

func PossibleValuesForVppTokenActionResultActionState() []string {
	return []string{
		string(VppTokenActionResultActionStateactive),
		string(VppTokenActionResultActionStatecanceled),
		string(VppTokenActionResultActionStatedone),
		string(VppTokenActionResultActionStatefailed),
		string(VppTokenActionResultActionStatenone),
		string(VppTokenActionResultActionStatenotSupported),
		string(VppTokenActionResultActionStatepending),
	}
}

func parseVppTokenActionResultActionState(input string) (*VppTokenActionResultActionState, error) {
	vals := map[string]VppTokenActionResultActionState{
		"active":       VppTokenActionResultActionStateactive,
		"canceled":     VppTokenActionResultActionStatecanceled,
		"done":         VppTokenActionResultActionStatedone,
		"failed":       VppTokenActionResultActionStatefailed,
		"none":         VppTokenActionResultActionStatenone,
		"notsupported": VppTokenActionResultActionStatenotSupported,
		"pending":      VppTokenActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenActionResultActionState(input)
	return &out, nil
}
