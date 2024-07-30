package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessConditionSetSignInRiskLevels string

const (
	ConditionalAccessConditionSetSignInRiskLevels_Hidden ConditionalAccessConditionSetSignInRiskLevels = "hidden"
	ConditionalAccessConditionSetSignInRiskLevels_High   ConditionalAccessConditionSetSignInRiskLevels = "high"
	ConditionalAccessConditionSetSignInRiskLevels_Low    ConditionalAccessConditionSetSignInRiskLevels = "low"
	ConditionalAccessConditionSetSignInRiskLevels_Medium ConditionalAccessConditionSetSignInRiskLevels = "medium"
	ConditionalAccessConditionSetSignInRiskLevels_None   ConditionalAccessConditionSetSignInRiskLevels = "none"
)

func PossibleValuesForConditionalAccessConditionSetSignInRiskLevels() []string {
	return []string{
		string(ConditionalAccessConditionSetSignInRiskLevels_Hidden),
		string(ConditionalAccessConditionSetSignInRiskLevels_High),
		string(ConditionalAccessConditionSetSignInRiskLevels_Low),
		string(ConditionalAccessConditionSetSignInRiskLevels_Medium),
		string(ConditionalAccessConditionSetSignInRiskLevels_None),
	}
}

func (s *ConditionalAccessConditionSetSignInRiskLevels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessConditionSetSignInRiskLevels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessConditionSetSignInRiskLevels(input string) (*ConditionalAccessConditionSetSignInRiskLevels, error) {
	vals := map[string]ConditionalAccessConditionSetSignInRiskLevels{
		"hidden": ConditionalAccessConditionSetSignInRiskLevels_Hidden,
		"high":   ConditionalAccessConditionSetSignInRiskLevels_High,
		"low":    ConditionalAccessConditionSetSignInRiskLevels_Low,
		"medium": ConditionalAccessConditionSetSignInRiskLevels_Medium,
		"none":   ConditionalAccessConditionSetSignInRiskLevels_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessConditionSetSignInRiskLevels(input)
	return &out, nil
}
