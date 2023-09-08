package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserRiskState string

const (
	RiskyUserRiskStateatRisk               RiskyUserRiskState = "AtRisk"
	RiskyUserRiskStateconfirmedCompromised RiskyUserRiskState = "ConfirmedCompromised"
	RiskyUserRiskStateconfirmedSafe        RiskyUserRiskState = "ConfirmedSafe"
	RiskyUserRiskStatedismissed            RiskyUserRiskState = "Dismissed"
	RiskyUserRiskStatenone                 RiskyUserRiskState = "None"
	RiskyUserRiskStateremediated           RiskyUserRiskState = "Remediated"
)

func PossibleValuesForRiskyUserRiskState() []string {
	return []string{
		string(RiskyUserRiskStateatRisk),
		string(RiskyUserRiskStateconfirmedCompromised),
		string(RiskyUserRiskStateconfirmedSafe),
		string(RiskyUserRiskStatedismissed),
		string(RiskyUserRiskStatenone),
		string(RiskyUserRiskStateremediated),
	}
}

func parseRiskyUserRiskState(input string) (*RiskyUserRiskState, error) {
	vals := map[string]RiskyUserRiskState{
		"atrisk":               RiskyUserRiskStateatRisk,
		"confirmedcompromised": RiskyUserRiskStateconfirmedCompromised,
		"confirmedsafe":        RiskyUserRiskStateconfirmedSafe,
		"dismissed":            RiskyUserRiskStatedismissed,
		"none":                 RiskyUserRiskStatenone,
		"remediated":           RiskyUserRiskStateremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyUserRiskState(input)
	return &out, nil
}
