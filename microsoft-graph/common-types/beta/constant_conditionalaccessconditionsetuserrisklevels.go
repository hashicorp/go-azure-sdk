package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessConditionSetUserRiskLevels string

const (
	ConditionalAccessConditionSetUserRiskLevels_Hidden ConditionalAccessConditionSetUserRiskLevels = "hidden"
	ConditionalAccessConditionSetUserRiskLevels_High   ConditionalAccessConditionSetUserRiskLevels = "high"
	ConditionalAccessConditionSetUserRiskLevels_Low    ConditionalAccessConditionSetUserRiskLevels = "low"
	ConditionalAccessConditionSetUserRiskLevels_Medium ConditionalAccessConditionSetUserRiskLevels = "medium"
	ConditionalAccessConditionSetUserRiskLevels_None   ConditionalAccessConditionSetUserRiskLevels = "none"
)

func PossibleValuesForConditionalAccessConditionSetUserRiskLevels() []string {
	return []string{
		string(ConditionalAccessConditionSetUserRiskLevels_Hidden),
		string(ConditionalAccessConditionSetUserRiskLevels_High),
		string(ConditionalAccessConditionSetUserRiskLevels_Low),
		string(ConditionalAccessConditionSetUserRiskLevels_Medium),
		string(ConditionalAccessConditionSetUserRiskLevels_None),
	}
}

func (s *ConditionalAccessConditionSetUserRiskLevels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessConditionSetUserRiskLevels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessConditionSetUserRiskLevels(input string) (*ConditionalAccessConditionSetUserRiskLevels, error) {
	vals := map[string]ConditionalAccessConditionSetUserRiskLevels{
		"hidden": ConditionalAccessConditionSetUserRiskLevels_Hidden,
		"high":   ConditionalAccessConditionSetUserRiskLevels_High,
		"low":    ConditionalAccessConditionSetUserRiskLevels_Low,
		"medium": ConditionalAccessConditionSetUserRiskLevels_Medium,
		"none":   ConditionalAccessConditionSetUserRiskLevels_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessConditionSetUserRiskLevels(input)
	return &out, nil
}
