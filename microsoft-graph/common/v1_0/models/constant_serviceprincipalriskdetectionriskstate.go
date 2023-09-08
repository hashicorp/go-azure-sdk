package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalRiskDetectionRiskState string

const (
	ServicePrincipalRiskDetectionRiskStateatRisk               ServicePrincipalRiskDetectionRiskState = "AtRisk"
	ServicePrincipalRiskDetectionRiskStateconfirmedCompromised ServicePrincipalRiskDetectionRiskState = "ConfirmedCompromised"
	ServicePrincipalRiskDetectionRiskStateconfirmedSafe        ServicePrincipalRiskDetectionRiskState = "ConfirmedSafe"
	ServicePrincipalRiskDetectionRiskStatedismissed            ServicePrincipalRiskDetectionRiskState = "Dismissed"
	ServicePrincipalRiskDetectionRiskStatenone                 ServicePrincipalRiskDetectionRiskState = "None"
	ServicePrincipalRiskDetectionRiskStateremediated           ServicePrincipalRiskDetectionRiskState = "Remediated"
)

func PossibleValuesForServicePrincipalRiskDetectionRiskState() []string {
	return []string{
		string(ServicePrincipalRiskDetectionRiskStateatRisk),
		string(ServicePrincipalRiskDetectionRiskStateconfirmedCompromised),
		string(ServicePrincipalRiskDetectionRiskStateconfirmedSafe),
		string(ServicePrincipalRiskDetectionRiskStatedismissed),
		string(ServicePrincipalRiskDetectionRiskStatenone),
		string(ServicePrincipalRiskDetectionRiskStateremediated),
	}
}

func parseServicePrincipalRiskDetectionRiskState(input string) (*ServicePrincipalRiskDetectionRiskState, error) {
	vals := map[string]ServicePrincipalRiskDetectionRiskState{
		"atrisk":               ServicePrincipalRiskDetectionRiskStateatRisk,
		"confirmedcompromised": ServicePrincipalRiskDetectionRiskStateconfirmedCompromised,
		"confirmedsafe":        ServicePrincipalRiskDetectionRiskStateconfirmedSafe,
		"dismissed":            ServicePrincipalRiskDetectionRiskStatedismissed,
		"none":                 ServicePrincipalRiskDetectionRiskStatenone,
		"remediated":           ServicePrincipalRiskDetectionRiskStateremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePrincipalRiskDetectionRiskState(input)
	return &out, nil
}
