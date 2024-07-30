package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionRiskLevel string

const (
	RiskDetectionRiskLevel_Hidden RiskDetectionRiskLevel = "hidden"
	RiskDetectionRiskLevel_High   RiskDetectionRiskLevel = "high"
	RiskDetectionRiskLevel_Low    RiskDetectionRiskLevel = "low"
	RiskDetectionRiskLevel_Medium RiskDetectionRiskLevel = "medium"
	RiskDetectionRiskLevel_None   RiskDetectionRiskLevel = "none"
)

func PossibleValuesForRiskDetectionRiskLevel() []string {
	return []string{
		string(RiskDetectionRiskLevel_Hidden),
		string(RiskDetectionRiskLevel_High),
		string(RiskDetectionRiskLevel_Low),
		string(RiskDetectionRiskLevel_Medium),
		string(RiskDetectionRiskLevel_None),
	}
}

func (s *RiskDetectionRiskLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskDetectionRiskLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskDetectionRiskLevel(input string) (*RiskDetectionRiskLevel, error) {
	vals := map[string]RiskDetectionRiskLevel{
		"hidden": RiskDetectionRiskLevel_Hidden,
		"high":   RiskDetectionRiskLevel_High,
		"low":    RiskDetectionRiskLevel_Low,
		"medium": RiskDetectionRiskLevel_Medium,
		"none":   RiskDetectionRiskLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionRiskLevel(input)
	return &out, nil
}
