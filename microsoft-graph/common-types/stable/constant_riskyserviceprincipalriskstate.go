package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipalRiskState string

const (
	RiskyServicePrincipalRiskState_AtRisk               RiskyServicePrincipalRiskState = "atRisk"
	RiskyServicePrincipalRiskState_ConfirmedCompromised RiskyServicePrincipalRiskState = "confirmedCompromised"
	RiskyServicePrincipalRiskState_ConfirmedSafe        RiskyServicePrincipalRiskState = "confirmedSafe"
	RiskyServicePrincipalRiskState_Dismissed            RiskyServicePrincipalRiskState = "dismissed"
	RiskyServicePrincipalRiskState_None                 RiskyServicePrincipalRiskState = "none"
	RiskyServicePrincipalRiskState_Remediated           RiskyServicePrincipalRiskState = "remediated"
)

func PossibleValuesForRiskyServicePrincipalRiskState() []string {
	return []string{
		string(RiskyServicePrincipalRiskState_AtRisk),
		string(RiskyServicePrincipalRiskState_ConfirmedCompromised),
		string(RiskyServicePrincipalRiskState_ConfirmedSafe),
		string(RiskyServicePrincipalRiskState_Dismissed),
		string(RiskyServicePrincipalRiskState_None),
		string(RiskyServicePrincipalRiskState_Remediated),
	}
}

func (s *RiskyServicePrincipalRiskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskyServicePrincipalRiskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskyServicePrincipalRiskState(input string) (*RiskyServicePrincipalRiskState, error) {
	vals := map[string]RiskyServicePrincipalRiskState{
		"atrisk":               RiskyServicePrincipalRiskState_AtRisk,
		"confirmedcompromised": RiskyServicePrincipalRiskState_ConfirmedCompromised,
		"confirmedsafe":        RiskyServicePrincipalRiskState_ConfirmedSafe,
		"dismissed":            RiskyServicePrincipalRiskState_Dismissed,
		"none":                 RiskyServicePrincipalRiskState_None,
		"remediated":           RiskyServicePrincipalRiskState_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyServicePrincipalRiskState(input)
	return &out, nil
}
