package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOsVppAppRevokeLicensesActionResultActionState string

const (
	MacOsVppAppRevokeLicensesActionResultActionStateactive       MacOsVppAppRevokeLicensesActionResultActionState = "Active"
	MacOsVppAppRevokeLicensesActionResultActionStatecanceled     MacOsVppAppRevokeLicensesActionResultActionState = "Canceled"
	MacOsVppAppRevokeLicensesActionResultActionStatedone         MacOsVppAppRevokeLicensesActionResultActionState = "Done"
	MacOsVppAppRevokeLicensesActionResultActionStatefailed       MacOsVppAppRevokeLicensesActionResultActionState = "Failed"
	MacOsVppAppRevokeLicensesActionResultActionStatenone         MacOsVppAppRevokeLicensesActionResultActionState = "None"
	MacOsVppAppRevokeLicensesActionResultActionStatenotSupported MacOsVppAppRevokeLicensesActionResultActionState = "NotSupported"
	MacOsVppAppRevokeLicensesActionResultActionStatepending      MacOsVppAppRevokeLicensesActionResultActionState = "Pending"
)

func PossibleValuesForMacOsVppAppRevokeLicensesActionResultActionState() []string {
	return []string{
		string(MacOsVppAppRevokeLicensesActionResultActionStateactive),
		string(MacOsVppAppRevokeLicensesActionResultActionStatecanceled),
		string(MacOsVppAppRevokeLicensesActionResultActionStatedone),
		string(MacOsVppAppRevokeLicensesActionResultActionStatefailed),
		string(MacOsVppAppRevokeLicensesActionResultActionStatenone),
		string(MacOsVppAppRevokeLicensesActionResultActionStatenotSupported),
		string(MacOsVppAppRevokeLicensesActionResultActionStatepending),
	}
}

func parseMacOsVppAppRevokeLicensesActionResultActionState(input string) (*MacOsVppAppRevokeLicensesActionResultActionState, error) {
	vals := map[string]MacOsVppAppRevokeLicensesActionResultActionState{
		"active":       MacOsVppAppRevokeLicensesActionResultActionStateactive,
		"canceled":     MacOsVppAppRevokeLicensesActionResultActionStatecanceled,
		"done":         MacOsVppAppRevokeLicensesActionResultActionStatedone,
		"failed":       MacOsVppAppRevokeLicensesActionResultActionStatefailed,
		"none":         MacOsVppAppRevokeLicensesActionResultActionStatenone,
		"notsupported": MacOsVppAppRevokeLicensesActionResultActionStatenotSupported,
		"pending":      MacOsVppAppRevokeLicensesActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOsVppAppRevokeLicensesActionResultActionState(input)
	return &out, nil
}
