package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserRiskState string

const (
	RiskyUserRiskState_AtRisk               RiskyUserRiskState = "atRisk"
	RiskyUserRiskState_ConfirmedCompromised RiskyUserRiskState = "confirmedCompromised"
	RiskyUserRiskState_ConfirmedSafe        RiskyUserRiskState = "confirmedSafe"
	RiskyUserRiskState_Dismissed            RiskyUserRiskState = "dismissed"
	RiskyUserRiskState_None                 RiskyUserRiskState = "none"
	RiskyUserRiskState_Remediated           RiskyUserRiskState = "remediated"
)

func PossibleValuesForRiskyUserRiskState() []string {
	return []string{
		string(RiskyUserRiskState_AtRisk),
		string(RiskyUserRiskState_ConfirmedCompromised),
		string(RiskyUserRiskState_ConfirmedSafe),
		string(RiskyUserRiskState_Dismissed),
		string(RiskyUserRiskState_None),
		string(RiskyUserRiskState_Remediated),
	}
}

func (s *RiskyUserRiskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskyUserRiskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskyUserRiskState(input string) (*RiskyUserRiskState, error) {
	vals := map[string]RiskyUserRiskState{
		"atrisk":               RiskyUserRiskState_AtRisk,
		"confirmedcompromised": RiskyUserRiskState_ConfirmedCompromised,
		"confirmedsafe":        RiskyUserRiskState_ConfirmedSafe,
		"dismissed":            RiskyUserRiskState_Dismissed,
		"none":                 RiskyUserRiskState_None,
		"remediated":           RiskyUserRiskState_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyUserRiskState(input)
	return &out, nil
}
