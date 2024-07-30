package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCompliancePolicyRequiredPasswordComplexity string

const (
	AndroidCompliancePolicyRequiredPasswordComplexity_High   AndroidCompliancePolicyRequiredPasswordComplexity = "high"
	AndroidCompliancePolicyRequiredPasswordComplexity_Low    AndroidCompliancePolicyRequiredPasswordComplexity = "low"
	AndroidCompliancePolicyRequiredPasswordComplexity_Medium AndroidCompliancePolicyRequiredPasswordComplexity = "medium"
	AndroidCompliancePolicyRequiredPasswordComplexity_None   AndroidCompliancePolicyRequiredPasswordComplexity = "none"
)

func PossibleValuesForAndroidCompliancePolicyRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidCompliancePolicyRequiredPasswordComplexity_High),
		string(AndroidCompliancePolicyRequiredPasswordComplexity_Low),
		string(AndroidCompliancePolicyRequiredPasswordComplexity_Medium),
		string(AndroidCompliancePolicyRequiredPasswordComplexity_None),
	}
}

func (s *AndroidCompliancePolicyRequiredPasswordComplexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidCompliancePolicyRequiredPasswordComplexity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidCompliancePolicyRequiredPasswordComplexity(input string) (*AndroidCompliancePolicyRequiredPasswordComplexity, error) {
	vals := map[string]AndroidCompliancePolicyRequiredPasswordComplexity{
		"high":   AndroidCompliancePolicyRequiredPasswordComplexity_High,
		"low":    AndroidCompliancePolicyRequiredPasswordComplexity_Low,
		"medium": AndroidCompliancePolicyRequiredPasswordComplexity_Medium,
		"none":   AndroidCompliancePolicyRequiredPasswordComplexity_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCompliancePolicyRequiredPasswordComplexity(input)
	return &out, nil
}
