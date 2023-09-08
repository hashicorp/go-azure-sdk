package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RevokeAppleVppLicensesActionResultActionState string

const (
	RevokeAppleVppLicensesActionResultActionStateactive       RevokeAppleVppLicensesActionResultActionState = "Active"
	RevokeAppleVppLicensesActionResultActionStatecanceled     RevokeAppleVppLicensesActionResultActionState = "Canceled"
	RevokeAppleVppLicensesActionResultActionStatedone         RevokeAppleVppLicensesActionResultActionState = "Done"
	RevokeAppleVppLicensesActionResultActionStatefailed       RevokeAppleVppLicensesActionResultActionState = "Failed"
	RevokeAppleVppLicensesActionResultActionStatenone         RevokeAppleVppLicensesActionResultActionState = "None"
	RevokeAppleVppLicensesActionResultActionStatenotSupported RevokeAppleVppLicensesActionResultActionState = "NotSupported"
	RevokeAppleVppLicensesActionResultActionStatepending      RevokeAppleVppLicensesActionResultActionState = "Pending"
)

func PossibleValuesForRevokeAppleVppLicensesActionResultActionState() []string {
	return []string{
		string(RevokeAppleVppLicensesActionResultActionStateactive),
		string(RevokeAppleVppLicensesActionResultActionStatecanceled),
		string(RevokeAppleVppLicensesActionResultActionStatedone),
		string(RevokeAppleVppLicensesActionResultActionStatefailed),
		string(RevokeAppleVppLicensesActionResultActionStatenone),
		string(RevokeAppleVppLicensesActionResultActionStatenotSupported),
		string(RevokeAppleVppLicensesActionResultActionStatepending),
	}
}

func parseRevokeAppleVppLicensesActionResultActionState(input string) (*RevokeAppleVppLicensesActionResultActionState, error) {
	vals := map[string]RevokeAppleVppLicensesActionResultActionState{
		"active":       RevokeAppleVppLicensesActionResultActionStateactive,
		"canceled":     RevokeAppleVppLicensesActionResultActionStatecanceled,
		"done":         RevokeAppleVppLicensesActionResultActionStatedone,
		"failed":       RevokeAppleVppLicensesActionResultActionStatefailed,
		"none":         RevokeAppleVppLicensesActionResultActionStatenone,
		"notsupported": RevokeAppleVppLicensesActionResultActionStatenotSupported,
		"pending":      RevokeAppleVppLicensesActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RevokeAppleVppLicensesActionResultActionState(input)
	return &out, nil
}
