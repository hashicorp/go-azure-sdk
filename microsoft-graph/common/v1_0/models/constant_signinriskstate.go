package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInRiskState string

const (
	SignInRiskStateatRisk               SignInRiskState = "AtRisk"
	SignInRiskStateconfirmedCompromised SignInRiskState = "ConfirmedCompromised"
	SignInRiskStateconfirmedSafe        SignInRiskState = "ConfirmedSafe"
	SignInRiskStatedismissed            SignInRiskState = "Dismissed"
	SignInRiskStatenone                 SignInRiskState = "None"
	SignInRiskStateremediated           SignInRiskState = "Remediated"
)

func PossibleValuesForSignInRiskState() []string {
	return []string{
		string(SignInRiskStateatRisk),
		string(SignInRiskStateconfirmedCompromised),
		string(SignInRiskStateconfirmedSafe),
		string(SignInRiskStatedismissed),
		string(SignInRiskStatenone),
		string(SignInRiskStateremediated),
	}
}

func parseSignInRiskState(input string) (*SignInRiskState, error) {
	vals := map[string]SignInRiskState{
		"atrisk":               SignInRiskStateatRisk,
		"confirmedcompromised": SignInRiskStateconfirmedCompromised,
		"confirmedsafe":        SignInRiskStateconfirmedSafe,
		"dismissed":            SignInRiskStatedismissed,
		"none":                 SignInRiskStatenone,
		"remediated":           SignInRiskStateremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInRiskState(input)
	return &out, nil
}
