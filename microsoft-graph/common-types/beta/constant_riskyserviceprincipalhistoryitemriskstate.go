package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipalHistoryItemRiskState string

const (
	RiskyServicePrincipalHistoryItemRiskState_AtRisk               RiskyServicePrincipalHistoryItemRiskState = "atRisk"
	RiskyServicePrincipalHistoryItemRiskState_ConfirmedCompromised RiskyServicePrincipalHistoryItemRiskState = "confirmedCompromised"
	RiskyServicePrincipalHistoryItemRiskState_ConfirmedSafe        RiskyServicePrincipalHistoryItemRiskState = "confirmedSafe"
	RiskyServicePrincipalHistoryItemRiskState_Dismissed            RiskyServicePrincipalHistoryItemRiskState = "dismissed"
	RiskyServicePrincipalHistoryItemRiskState_None                 RiskyServicePrincipalHistoryItemRiskState = "none"
	RiskyServicePrincipalHistoryItemRiskState_Remediated           RiskyServicePrincipalHistoryItemRiskState = "remediated"
)

func PossibleValuesForRiskyServicePrincipalHistoryItemRiskState() []string {
	return []string{
		string(RiskyServicePrincipalHistoryItemRiskState_AtRisk),
		string(RiskyServicePrincipalHistoryItemRiskState_ConfirmedCompromised),
		string(RiskyServicePrincipalHistoryItemRiskState_ConfirmedSafe),
		string(RiskyServicePrincipalHistoryItemRiskState_Dismissed),
		string(RiskyServicePrincipalHistoryItemRiskState_None),
		string(RiskyServicePrincipalHistoryItemRiskState_Remediated),
	}
}

func (s *RiskyServicePrincipalHistoryItemRiskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskyServicePrincipalHistoryItemRiskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskyServicePrincipalHistoryItemRiskState(input string) (*RiskyServicePrincipalHistoryItemRiskState, error) {
	vals := map[string]RiskyServicePrincipalHistoryItemRiskState{
		"atrisk":               RiskyServicePrincipalHistoryItemRiskState_AtRisk,
		"confirmedcompromised": RiskyServicePrincipalHistoryItemRiskState_ConfirmedCompromised,
		"confirmedsafe":        RiskyServicePrincipalHistoryItemRiskState_ConfirmedSafe,
		"dismissed":            RiskyServicePrincipalHistoryItemRiskState_Dismissed,
		"none":                 RiskyServicePrincipalHistoryItemRiskState_None,
		"remediated":           RiskyServicePrincipalHistoryItemRiskState_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyServicePrincipalHistoryItemRiskState(input)
	return &out, nil
}
