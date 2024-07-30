package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionRiskState string

const (
	RiskDetectionRiskState_AtRisk               RiskDetectionRiskState = "atRisk"
	RiskDetectionRiskState_ConfirmedCompromised RiskDetectionRiskState = "confirmedCompromised"
	RiskDetectionRiskState_ConfirmedSafe        RiskDetectionRiskState = "confirmedSafe"
	RiskDetectionRiskState_Dismissed            RiskDetectionRiskState = "dismissed"
	RiskDetectionRiskState_None                 RiskDetectionRiskState = "none"
	RiskDetectionRiskState_Remediated           RiskDetectionRiskState = "remediated"
)

func PossibleValuesForRiskDetectionRiskState() []string {
	return []string{
		string(RiskDetectionRiskState_AtRisk),
		string(RiskDetectionRiskState_ConfirmedCompromised),
		string(RiskDetectionRiskState_ConfirmedSafe),
		string(RiskDetectionRiskState_Dismissed),
		string(RiskDetectionRiskState_None),
		string(RiskDetectionRiskState_Remediated),
	}
}

func (s *RiskDetectionRiskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskDetectionRiskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskDetectionRiskState(input string) (*RiskDetectionRiskState, error) {
	vals := map[string]RiskDetectionRiskState{
		"atrisk":               RiskDetectionRiskState_AtRisk,
		"confirmedcompromised": RiskDetectionRiskState_ConfirmedCompromised,
		"confirmedsafe":        RiskDetectionRiskState_ConfirmedSafe,
		"dismissed":            RiskDetectionRiskState_Dismissed,
		"none":                 RiskDetectionRiskState_None,
		"remediated":           RiskDetectionRiskState_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionRiskState(input)
	return &out, nil
}
