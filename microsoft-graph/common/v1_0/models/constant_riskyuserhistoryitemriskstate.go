package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserHistoryItemRiskState string

const (
	RiskyUserHistoryItemRiskStateatRisk               RiskyUserHistoryItemRiskState = "AtRisk"
	RiskyUserHistoryItemRiskStateconfirmedCompromised RiskyUserHistoryItemRiskState = "ConfirmedCompromised"
	RiskyUserHistoryItemRiskStateconfirmedSafe        RiskyUserHistoryItemRiskState = "ConfirmedSafe"
	RiskyUserHistoryItemRiskStatedismissed            RiskyUserHistoryItemRiskState = "Dismissed"
	RiskyUserHistoryItemRiskStatenone                 RiskyUserHistoryItemRiskState = "None"
	RiskyUserHistoryItemRiskStateremediated           RiskyUserHistoryItemRiskState = "Remediated"
)

func PossibleValuesForRiskyUserHistoryItemRiskState() []string {
	return []string{
		string(RiskyUserHistoryItemRiskStateatRisk),
		string(RiskyUserHistoryItemRiskStateconfirmedCompromised),
		string(RiskyUserHistoryItemRiskStateconfirmedSafe),
		string(RiskyUserHistoryItemRiskStatedismissed),
		string(RiskyUserHistoryItemRiskStatenone),
		string(RiskyUserHistoryItemRiskStateremediated),
	}
}

func parseRiskyUserHistoryItemRiskState(input string) (*RiskyUserHistoryItemRiskState, error) {
	vals := map[string]RiskyUserHistoryItemRiskState{
		"atrisk":               RiskyUserHistoryItemRiskStateatRisk,
		"confirmedcompromised": RiskyUserHistoryItemRiskStateconfirmedCompromised,
		"confirmedsafe":        RiskyUserHistoryItemRiskStateconfirmedSafe,
		"dismissed":            RiskyUserHistoryItemRiskStatedismissed,
		"none":                 RiskyUserHistoryItemRiskStatenone,
		"remediated":           RiskyUserHistoryItemRiskStateremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyUserHistoryItemRiskState(input)
	return &out, nil
}
