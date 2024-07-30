package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInRiskState string

const (
	SignInRiskState_AtRisk               SignInRiskState = "atRisk"
	SignInRiskState_ConfirmedCompromised SignInRiskState = "confirmedCompromised"
	SignInRiskState_ConfirmedSafe        SignInRiskState = "confirmedSafe"
	SignInRiskState_Dismissed            SignInRiskState = "dismissed"
	SignInRiskState_None                 SignInRiskState = "none"
	SignInRiskState_Remediated           SignInRiskState = "remediated"
)

func PossibleValuesForSignInRiskState() []string {
	return []string{
		string(SignInRiskState_AtRisk),
		string(SignInRiskState_ConfirmedCompromised),
		string(SignInRiskState_ConfirmedSafe),
		string(SignInRiskState_Dismissed),
		string(SignInRiskState_None),
		string(SignInRiskState_Remediated),
	}
}

func (s *SignInRiskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInRiskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInRiskState(input string) (*SignInRiskState, error) {
	vals := map[string]SignInRiskState{
		"atrisk":               SignInRiskState_AtRisk,
		"confirmedcompromised": SignInRiskState_ConfirmedCompromised,
		"confirmedsafe":        SignInRiskState_ConfirmedSafe,
		"dismissed":            SignInRiskState_Dismissed,
		"none":                 SignInRiskState_None,
		"remediated":           SignInRiskState_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInRiskState(input)
	return &out, nil
}
