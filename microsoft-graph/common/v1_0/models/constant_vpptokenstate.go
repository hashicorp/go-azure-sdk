package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenState string

const (
	VppTokenStateassignedToExternalMDM VppTokenState = "AssignedToExternalMDM"
	VppTokenStateexpired               VppTokenState = "Expired"
	VppTokenStateinvalid               VppTokenState = "Invalid"
	VppTokenStateunknown               VppTokenState = "Unknown"
	VppTokenStatevalid                 VppTokenState = "Valid"
)

func PossibleValuesForVppTokenState() []string {
	return []string{
		string(VppTokenStateassignedToExternalMDM),
		string(VppTokenStateexpired),
		string(VppTokenStateinvalid),
		string(VppTokenStateunknown),
		string(VppTokenStatevalid),
	}
}

func parseVppTokenState(input string) (*VppTokenState, error) {
	vals := map[string]VppTokenState{
		"assignedtoexternalmdm": VppTokenStateassignedToExternalMDM,
		"expired":               VppTokenStateexpired,
		"invalid":               VppTokenStateinvalid,
		"unknown":               VppTokenStateunknown,
		"valid":                 VppTokenStatevalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenState(input)
	return &out, nil
}
