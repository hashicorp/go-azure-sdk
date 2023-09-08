package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipalHistoryItemRiskState string

const (
	RiskyServicePrincipalHistoryItemRiskStateatRisk               RiskyServicePrincipalHistoryItemRiskState = "AtRisk"
	RiskyServicePrincipalHistoryItemRiskStateconfirmedCompromised RiskyServicePrincipalHistoryItemRiskState = "ConfirmedCompromised"
	RiskyServicePrincipalHistoryItemRiskStateconfirmedSafe        RiskyServicePrincipalHistoryItemRiskState = "ConfirmedSafe"
	RiskyServicePrincipalHistoryItemRiskStatedismissed            RiskyServicePrincipalHistoryItemRiskState = "Dismissed"
	RiskyServicePrincipalHistoryItemRiskStatenone                 RiskyServicePrincipalHistoryItemRiskState = "None"
	RiskyServicePrincipalHistoryItemRiskStateremediated           RiskyServicePrincipalHistoryItemRiskState = "Remediated"
)

func PossibleValuesForRiskyServicePrincipalHistoryItemRiskState() []string {
	return []string{
		string(RiskyServicePrincipalHistoryItemRiskStateatRisk),
		string(RiskyServicePrincipalHistoryItemRiskStateconfirmedCompromised),
		string(RiskyServicePrincipalHistoryItemRiskStateconfirmedSafe),
		string(RiskyServicePrincipalHistoryItemRiskStatedismissed),
		string(RiskyServicePrincipalHistoryItemRiskStatenone),
		string(RiskyServicePrincipalHistoryItemRiskStateremediated),
	}
}

func parseRiskyServicePrincipalHistoryItemRiskState(input string) (*RiskyServicePrincipalHistoryItemRiskState, error) {
	vals := map[string]RiskyServicePrincipalHistoryItemRiskState{
		"atrisk":               RiskyServicePrincipalHistoryItemRiskStateatRisk,
		"confirmedcompromised": RiskyServicePrincipalHistoryItemRiskStateconfirmedCompromised,
		"confirmedsafe":        RiskyServicePrincipalHistoryItemRiskStateconfirmedSafe,
		"dismissed":            RiskyServicePrincipalHistoryItemRiskStatedismissed,
		"none":                 RiskyServicePrincipalHistoryItemRiskStatenone,
		"remediated":           RiskyServicePrincipalHistoryItemRiskStateremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyServicePrincipalHistoryItemRiskState(input)
	return &out, nil
}
