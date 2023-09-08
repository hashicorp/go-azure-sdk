package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetPasscodeActionResultActionState string

const (
	ResetPasscodeActionResultActionStateactive       ResetPasscodeActionResultActionState = "Active"
	ResetPasscodeActionResultActionStatecanceled     ResetPasscodeActionResultActionState = "Canceled"
	ResetPasscodeActionResultActionStatedone         ResetPasscodeActionResultActionState = "Done"
	ResetPasscodeActionResultActionStatefailed       ResetPasscodeActionResultActionState = "Failed"
	ResetPasscodeActionResultActionStatenone         ResetPasscodeActionResultActionState = "None"
	ResetPasscodeActionResultActionStatenotSupported ResetPasscodeActionResultActionState = "NotSupported"
	ResetPasscodeActionResultActionStatepending      ResetPasscodeActionResultActionState = "Pending"
)

func PossibleValuesForResetPasscodeActionResultActionState() []string {
	return []string{
		string(ResetPasscodeActionResultActionStateactive),
		string(ResetPasscodeActionResultActionStatecanceled),
		string(ResetPasscodeActionResultActionStatedone),
		string(ResetPasscodeActionResultActionStatefailed),
		string(ResetPasscodeActionResultActionStatenone),
		string(ResetPasscodeActionResultActionStatenotSupported),
		string(ResetPasscodeActionResultActionStatepending),
	}
}

func parseResetPasscodeActionResultActionState(input string) (*ResetPasscodeActionResultActionState, error) {
	vals := map[string]ResetPasscodeActionResultActionState{
		"active":       ResetPasscodeActionResultActionStateactive,
		"canceled":     ResetPasscodeActionResultActionStatecanceled,
		"done":         ResetPasscodeActionResultActionStatedone,
		"failed":       ResetPasscodeActionResultActionStatefailed,
		"none":         ResetPasscodeActionResultActionStatenone,
		"notsupported": ResetPasscodeActionResultActionStatenotSupported,
		"pending":      ResetPasscodeActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResetPasscodeActionResultActionState(input)
	return &out, nil
}
