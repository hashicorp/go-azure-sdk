package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity string

const (
	AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity_High   AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity = "high"
	AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity_Low    AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity = "low"
	AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity_Medium AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity = "medium"
	AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity_None   AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity = "none"
)

func PossibleValuesForAndroidWorkProfileCompliancePolicyRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity_High),
		string(AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity_Low),
		string(AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity_Medium),
		string(AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity_None),
	}
}

func (s *AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileCompliancePolicyRequiredPasswordComplexity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileCompliancePolicyRequiredPasswordComplexity(input string) (*AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity, error) {
	vals := map[string]AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity{
		"high":   AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity_High,
		"low":    AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity_Low,
		"medium": AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity_Medium,
		"none":   AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCompliancePolicyRequiredPasswordComplexity(input)
	return &out, nil
}
