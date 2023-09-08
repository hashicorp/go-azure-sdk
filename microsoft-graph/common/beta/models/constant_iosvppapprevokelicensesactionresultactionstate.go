package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppRevokeLicensesActionResultActionState string

const (
	IosVppAppRevokeLicensesActionResultActionStateactive       IosVppAppRevokeLicensesActionResultActionState = "Active"
	IosVppAppRevokeLicensesActionResultActionStatecanceled     IosVppAppRevokeLicensesActionResultActionState = "Canceled"
	IosVppAppRevokeLicensesActionResultActionStatedone         IosVppAppRevokeLicensesActionResultActionState = "Done"
	IosVppAppRevokeLicensesActionResultActionStatefailed       IosVppAppRevokeLicensesActionResultActionState = "Failed"
	IosVppAppRevokeLicensesActionResultActionStatenone         IosVppAppRevokeLicensesActionResultActionState = "None"
	IosVppAppRevokeLicensesActionResultActionStatenotSupported IosVppAppRevokeLicensesActionResultActionState = "NotSupported"
	IosVppAppRevokeLicensesActionResultActionStatepending      IosVppAppRevokeLicensesActionResultActionState = "Pending"
)

func PossibleValuesForIosVppAppRevokeLicensesActionResultActionState() []string {
	return []string{
		string(IosVppAppRevokeLicensesActionResultActionStateactive),
		string(IosVppAppRevokeLicensesActionResultActionStatecanceled),
		string(IosVppAppRevokeLicensesActionResultActionStatedone),
		string(IosVppAppRevokeLicensesActionResultActionStatefailed),
		string(IosVppAppRevokeLicensesActionResultActionStatenone),
		string(IosVppAppRevokeLicensesActionResultActionStatenotSupported),
		string(IosVppAppRevokeLicensesActionResultActionStatepending),
	}
}

func parseIosVppAppRevokeLicensesActionResultActionState(input string) (*IosVppAppRevokeLicensesActionResultActionState, error) {
	vals := map[string]IosVppAppRevokeLicensesActionResultActionState{
		"active":       IosVppAppRevokeLicensesActionResultActionStateactive,
		"canceled":     IosVppAppRevokeLicensesActionResultActionStatecanceled,
		"done":         IosVppAppRevokeLicensesActionResultActionStatedone,
		"failed":       IosVppAppRevokeLicensesActionResultActionStatefailed,
		"none":         IosVppAppRevokeLicensesActionResultActionStatenone,
		"notsupported": IosVppAppRevokeLicensesActionResultActionStatenotSupported,
		"pending":      IosVppAppRevokeLicensesActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVppAppRevokeLicensesActionResultActionState(input)
	return &out, nil
}
