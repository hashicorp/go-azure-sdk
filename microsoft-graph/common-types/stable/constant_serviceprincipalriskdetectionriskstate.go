package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalRiskDetectionRiskState string

const (
	ServicePrincipalRiskDetectionRiskState_AtRisk               ServicePrincipalRiskDetectionRiskState = "atRisk"
	ServicePrincipalRiskDetectionRiskState_ConfirmedCompromised ServicePrincipalRiskDetectionRiskState = "confirmedCompromised"
	ServicePrincipalRiskDetectionRiskState_ConfirmedSafe        ServicePrincipalRiskDetectionRiskState = "confirmedSafe"
	ServicePrincipalRiskDetectionRiskState_Dismissed            ServicePrincipalRiskDetectionRiskState = "dismissed"
	ServicePrincipalRiskDetectionRiskState_None                 ServicePrincipalRiskDetectionRiskState = "none"
	ServicePrincipalRiskDetectionRiskState_Remediated           ServicePrincipalRiskDetectionRiskState = "remediated"
)

func PossibleValuesForServicePrincipalRiskDetectionRiskState() []string {
	return []string{
		string(ServicePrincipalRiskDetectionRiskState_AtRisk),
		string(ServicePrincipalRiskDetectionRiskState_ConfirmedCompromised),
		string(ServicePrincipalRiskDetectionRiskState_ConfirmedSafe),
		string(ServicePrincipalRiskDetectionRiskState_Dismissed),
		string(ServicePrincipalRiskDetectionRiskState_None),
		string(ServicePrincipalRiskDetectionRiskState_Remediated),
	}
}

func (s *ServicePrincipalRiskDetectionRiskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServicePrincipalRiskDetectionRiskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServicePrincipalRiskDetectionRiskState(input string) (*ServicePrincipalRiskDetectionRiskState, error) {
	vals := map[string]ServicePrincipalRiskDetectionRiskState{
		"atrisk":               ServicePrincipalRiskDetectionRiskState_AtRisk,
		"confirmedcompromised": ServicePrincipalRiskDetectionRiskState_ConfirmedCompromised,
		"confirmedsafe":        ServicePrincipalRiskDetectionRiskState_ConfirmedSafe,
		"dismissed":            ServicePrincipalRiskDetectionRiskState_Dismissed,
		"none":                 ServicePrincipalRiskDetectionRiskState_None,
		"remediated":           ServicePrincipalRiskDetectionRiskState_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServicePrincipalRiskDetectionRiskState(input)
	return &out, nil
}
