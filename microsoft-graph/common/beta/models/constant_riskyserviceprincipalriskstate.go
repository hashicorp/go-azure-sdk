package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipalRiskState string

const (
	RiskyServicePrincipalRiskStateatRisk               RiskyServicePrincipalRiskState = "AtRisk"
	RiskyServicePrincipalRiskStateconfirmedCompromised RiskyServicePrincipalRiskState = "ConfirmedCompromised"
	RiskyServicePrincipalRiskStateconfirmedSafe        RiskyServicePrincipalRiskState = "ConfirmedSafe"
	RiskyServicePrincipalRiskStatedismissed            RiskyServicePrincipalRiskState = "Dismissed"
	RiskyServicePrincipalRiskStatenone                 RiskyServicePrincipalRiskState = "None"
	RiskyServicePrincipalRiskStateremediated           RiskyServicePrincipalRiskState = "Remediated"
)

func PossibleValuesForRiskyServicePrincipalRiskState() []string {
	return []string{
		string(RiskyServicePrincipalRiskStateatRisk),
		string(RiskyServicePrincipalRiskStateconfirmedCompromised),
		string(RiskyServicePrincipalRiskStateconfirmedSafe),
		string(RiskyServicePrincipalRiskStatedismissed),
		string(RiskyServicePrincipalRiskStatenone),
		string(RiskyServicePrincipalRiskStateremediated),
	}
}

func parseRiskyServicePrincipalRiskState(input string) (*RiskyServicePrincipalRiskState, error) {
	vals := map[string]RiskyServicePrincipalRiskState{
		"atrisk":               RiskyServicePrincipalRiskStateatRisk,
		"confirmedcompromised": RiskyServicePrincipalRiskStateconfirmedCompromised,
		"confirmedsafe":        RiskyServicePrincipalRiskStateconfirmedSafe,
		"dismissed":            RiskyServicePrincipalRiskStatedismissed,
		"none":                 RiskyServicePrincipalRiskStatenone,
		"remediated":           RiskyServicePrincipalRiskStateremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyServicePrincipalRiskState(input)
	return &out, nil
}
