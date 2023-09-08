package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionRiskState string

const (
	RiskDetectionRiskStateatRisk               RiskDetectionRiskState = "AtRisk"
	RiskDetectionRiskStateconfirmedCompromised RiskDetectionRiskState = "ConfirmedCompromised"
	RiskDetectionRiskStateconfirmedSafe        RiskDetectionRiskState = "ConfirmedSafe"
	RiskDetectionRiskStatedismissed            RiskDetectionRiskState = "Dismissed"
	RiskDetectionRiskStatenone                 RiskDetectionRiskState = "None"
	RiskDetectionRiskStateremediated           RiskDetectionRiskState = "Remediated"
)

func PossibleValuesForRiskDetectionRiskState() []string {
	return []string{
		string(RiskDetectionRiskStateatRisk),
		string(RiskDetectionRiskStateconfirmedCompromised),
		string(RiskDetectionRiskStateconfirmedSafe),
		string(RiskDetectionRiskStatedismissed),
		string(RiskDetectionRiskStatenone),
		string(RiskDetectionRiskStateremediated),
	}
}

func parseRiskDetectionRiskState(input string) (*RiskDetectionRiskState, error) {
	vals := map[string]RiskDetectionRiskState{
		"atrisk":               RiskDetectionRiskStateatRisk,
		"confirmedcompromised": RiskDetectionRiskStateconfirmedCompromised,
		"confirmedsafe":        RiskDetectionRiskStateconfirmedSafe,
		"dismissed":            RiskDetectionRiskStatedismissed,
		"none":                 RiskDetectionRiskStatenone,
		"remediated":           RiskDetectionRiskStateremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionRiskState(input)
	return &out, nil
}
