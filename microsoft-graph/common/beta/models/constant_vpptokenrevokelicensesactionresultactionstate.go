package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenRevokeLicensesActionResultActionState string

const (
	VppTokenRevokeLicensesActionResultActionStateactive       VppTokenRevokeLicensesActionResultActionState = "Active"
	VppTokenRevokeLicensesActionResultActionStatecanceled     VppTokenRevokeLicensesActionResultActionState = "Canceled"
	VppTokenRevokeLicensesActionResultActionStatedone         VppTokenRevokeLicensesActionResultActionState = "Done"
	VppTokenRevokeLicensesActionResultActionStatefailed       VppTokenRevokeLicensesActionResultActionState = "Failed"
	VppTokenRevokeLicensesActionResultActionStatenone         VppTokenRevokeLicensesActionResultActionState = "None"
	VppTokenRevokeLicensesActionResultActionStatenotSupported VppTokenRevokeLicensesActionResultActionState = "NotSupported"
	VppTokenRevokeLicensesActionResultActionStatepending      VppTokenRevokeLicensesActionResultActionState = "Pending"
)

func PossibleValuesForVppTokenRevokeLicensesActionResultActionState() []string {
	return []string{
		string(VppTokenRevokeLicensesActionResultActionStateactive),
		string(VppTokenRevokeLicensesActionResultActionStatecanceled),
		string(VppTokenRevokeLicensesActionResultActionStatedone),
		string(VppTokenRevokeLicensesActionResultActionStatefailed),
		string(VppTokenRevokeLicensesActionResultActionStatenone),
		string(VppTokenRevokeLicensesActionResultActionStatenotSupported),
		string(VppTokenRevokeLicensesActionResultActionStatepending),
	}
}

func parseVppTokenRevokeLicensesActionResultActionState(input string) (*VppTokenRevokeLicensesActionResultActionState, error) {
	vals := map[string]VppTokenRevokeLicensesActionResultActionState{
		"active":       VppTokenRevokeLicensesActionResultActionStateactive,
		"canceled":     VppTokenRevokeLicensesActionResultActionStatecanceled,
		"done":         VppTokenRevokeLicensesActionResultActionStatedone,
		"failed":       VppTokenRevokeLicensesActionResultActionStatefailed,
		"none":         VppTokenRevokeLicensesActionResultActionStatenone,
		"notsupported": VppTokenRevokeLicensesActionResultActionStatenotSupported,
		"pending":      VppTokenRevokeLicensesActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenRevokeLicensesActionResultActionState(input)
	return &out, nil
}
