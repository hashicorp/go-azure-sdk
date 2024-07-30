package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipalRiskLevel string

const (
	RiskyServicePrincipalRiskLevel_Hidden RiskyServicePrincipalRiskLevel = "hidden"
	RiskyServicePrincipalRiskLevel_High   RiskyServicePrincipalRiskLevel = "high"
	RiskyServicePrincipalRiskLevel_Low    RiskyServicePrincipalRiskLevel = "low"
	RiskyServicePrincipalRiskLevel_Medium RiskyServicePrincipalRiskLevel = "medium"
	RiskyServicePrincipalRiskLevel_None   RiskyServicePrincipalRiskLevel = "none"
)

func PossibleValuesForRiskyServicePrincipalRiskLevel() []string {
	return []string{
		string(RiskyServicePrincipalRiskLevel_Hidden),
		string(RiskyServicePrincipalRiskLevel_High),
		string(RiskyServicePrincipalRiskLevel_Low),
		string(RiskyServicePrincipalRiskLevel_Medium),
		string(RiskyServicePrincipalRiskLevel_None),
	}
}

func (s *RiskyServicePrincipalRiskLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskyServicePrincipalRiskLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskyServicePrincipalRiskLevel(input string) (*RiskyServicePrincipalRiskLevel, error) {
	vals := map[string]RiskyServicePrincipalRiskLevel{
		"hidden": RiskyServicePrincipalRiskLevel_Hidden,
		"high":   RiskyServicePrincipalRiskLevel_High,
		"low":    RiskyServicePrincipalRiskLevel_Low,
		"medium": RiskyServicePrincipalRiskLevel_Medium,
		"none":   RiskyServicePrincipalRiskLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyServicePrincipalRiskLevel(input)
	return &out, nil
}
