package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceRiskScore string

const (
	SecurityDeviceEvidenceRiskScore_High          SecurityDeviceEvidenceRiskScore = "high"
	SecurityDeviceEvidenceRiskScore_Informational SecurityDeviceEvidenceRiskScore = "informational"
	SecurityDeviceEvidenceRiskScore_Low           SecurityDeviceEvidenceRiskScore = "low"
	SecurityDeviceEvidenceRiskScore_Medium        SecurityDeviceEvidenceRiskScore = "medium"
	SecurityDeviceEvidenceRiskScore_None          SecurityDeviceEvidenceRiskScore = "none"
)

func PossibleValuesForSecurityDeviceEvidenceRiskScore() []string {
	return []string{
		string(SecurityDeviceEvidenceRiskScore_High),
		string(SecurityDeviceEvidenceRiskScore_Informational),
		string(SecurityDeviceEvidenceRiskScore_Low),
		string(SecurityDeviceEvidenceRiskScore_Medium),
		string(SecurityDeviceEvidenceRiskScore_None),
	}
}

func (s *SecurityDeviceEvidenceRiskScore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeviceEvidenceRiskScore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeviceEvidenceRiskScore(input string) (*SecurityDeviceEvidenceRiskScore, error) {
	vals := map[string]SecurityDeviceEvidenceRiskScore{
		"high":          SecurityDeviceEvidenceRiskScore_High,
		"informational": SecurityDeviceEvidenceRiskScore_Informational,
		"low":           SecurityDeviceEvidenceRiskScore_Low,
		"medium":        SecurityDeviceEvidenceRiskScore_Medium,
		"none":          SecurityDeviceEvidenceRiskScore_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceRiskScore(input)
	return &out, nil
}
