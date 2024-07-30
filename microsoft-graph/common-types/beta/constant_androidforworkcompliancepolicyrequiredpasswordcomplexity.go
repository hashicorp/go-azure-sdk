package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCompliancePolicyRequiredPasswordComplexity string

const (
	AndroidForWorkCompliancePolicyRequiredPasswordComplexity_High   AndroidForWorkCompliancePolicyRequiredPasswordComplexity = "high"
	AndroidForWorkCompliancePolicyRequiredPasswordComplexity_Low    AndroidForWorkCompliancePolicyRequiredPasswordComplexity = "low"
	AndroidForWorkCompliancePolicyRequiredPasswordComplexity_Medium AndroidForWorkCompliancePolicyRequiredPasswordComplexity = "medium"
	AndroidForWorkCompliancePolicyRequiredPasswordComplexity_None   AndroidForWorkCompliancePolicyRequiredPasswordComplexity = "none"
)

func PossibleValuesForAndroidForWorkCompliancePolicyRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidForWorkCompliancePolicyRequiredPasswordComplexity_High),
		string(AndroidForWorkCompliancePolicyRequiredPasswordComplexity_Low),
		string(AndroidForWorkCompliancePolicyRequiredPasswordComplexity_Medium),
		string(AndroidForWorkCompliancePolicyRequiredPasswordComplexity_None),
	}
}

func (s *AndroidForWorkCompliancePolicyRequiredPasswordComplexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkCompliancePolicyRequiredPasswordComplexity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkCompliancePolicyRequiredPasswordComplexity(input string) (*AndroidForWorkCompliancePolicyRequiredPasswordComplexity, error) {
	vals := map[string]AndroidForWorkCompliancePolicyRequiredPasswordComplexity{
		"high":   AndroidForWorkCompliancePolicyRequiredPasswordComplexity_High,
		"low":    AndroidForWorkCompliancePolicyRequiredPasswordComplexity_Low,
		"medium": AndroidForWorkCompliancePolicyRequiredPasswordComplexity_Medium,
		"none":   AndroidForWorkCompliancePolicyRequiredPasswordComplexity_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCompliancePolicyRequiredPasswordComplexity(input)
	return &out, nil
}
