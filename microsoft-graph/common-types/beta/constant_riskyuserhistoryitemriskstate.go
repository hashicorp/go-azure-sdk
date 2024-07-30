package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserHistoryItemRiskState string

const (
	RiskyUserHistoryItemRiskState_AtRisk               RiskyUserHistoryItemRiskState = "atRisk"
	RiskyUserHistoryItemRiskState_ConfirmedCompromised RiskyUserHistoryItemRiskState = "confirmedCompromised"
	RiskyUserHistoryItemRiskState_ConfirmedSafe        RiskyUserHistoryItemRiskState = "confirmedSafe"
	RiskyUserHistoryItemRiskState_Dismissed            RiskyUserHistoryItemRiskState = "dismissed"
	RiskyUserHistoryItemRiskState_None                 RiskyUserHistoryItemRiskState = "none"
	RiskyUserHistoryItemRiskState_Remediated           RiskyUserHistoryItemRiskState = "remediated"
)

func PossibleValuesForRiskyUserHistoryItemRiskState() []string {
	return []string{
		string(RiskyUserHistoryItemRiskState_AtRisk),
		string(RiskyUserHistoryItemRiskState_ConfirmedCompromised),
		string(RiskyUserHistoryItemRiskState_ConfirmedSafe),
		string(RiskyUserHistoryItemRiskState_Dismissed),
		string(RiskyUserHistoryItemRiskState_None),
		string(RiskyUserHistoryItemRiskState_Remediated),
	}
}

func (s *RiskyUserHistoryItemRiskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskyUserHistoryItemRiskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskyUserHistoryItemRiskState(input string) (*RiskyUserHistoryItemRiskState, error) {
	vals := map[string]RiskyUserHistoryItemRiskState{
		"atrisk":               RiskyUserHistoryItemRiskState_AtRisk,
		"confirmedcompromised": RiskyUserHistoryItemRiskState_ConfirmedCompromised,
		"confirmedsafe":        RiskyUserHistoryItemRiskState_ConfirmedSafe,
		"dismissed":            RiskyUserHistoryItemRiskState_Dismissed,
		"none":                 RiskyUserHistoryItemRiskState_None,
		"remediated":           RiskyUserHistoryItemRiskState_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyUserHistoryItemRiskState(input)
	return &out, nil
}
