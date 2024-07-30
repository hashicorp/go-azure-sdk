package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInRiskLevelAggregated string

const (
	SignInRiskLevelAggregated_Hidden SignInRiskLevelAggregated = "hidden"
	SignInRiskLevelAggregated_High   SignInRiskLevelAggregated = "high"
	SignInRiskLevelAggregated_Low    SignInRiskLevelAggregated = "low"
	SignInRiskLevelAggregated_Medium SignInRiskLevelAggregated = "medium"
	SignInRiskLevelAggregated_None   SignInRiskLevelAggregated = "none"
)

func PossibleValuesForSignInRiskLevelAggregated() []string {
	return []string{
		string(SignInRiskLevelAggregated_Hidden),
		string(SignInRiskLevelAggregated_High),
		string(SignInRiskLevelAggregated_Low),
		string(SignInRiskLevelAggregated_Medium),
		string(SignInRiskLevelAggregated_None),
	}
}

func (s *SignInRiskLevelAggregated) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInRiskLevelAggregated(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInRiskLevelAggregated(input string) (*SignInRiskLevelAggregated, error) {
	vals := map[string]SignInRiskLevelAggregated{
		"hidden": SignInRiskLevelAggregated_Hidden,
		"high":   SignInRiskLevelAggregated_High,
		"low":    SignInRiskLevelAggregated_Low,
		"medium": SignInRiskLevelAggregated_Medium,
		"none":   SignInRiskLevelAggregated_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInRiskLevelAggregated(input)
	return &out, nil
}
