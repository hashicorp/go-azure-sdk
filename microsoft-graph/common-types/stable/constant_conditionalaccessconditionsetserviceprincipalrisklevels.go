package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessConditionSetServicePrincipalRiskLevels string

const (
	ConditionalAccessConditionSetServicePrincipalRiskLevels_Hidden ConditionalAccessConditionSetServicePrincipalRiskLevels = "hidden"
	ConditionalAccessConditionSetServicePrincipalRiskLevels_High   ConditionalAccessConditionSetServicePrincipalRiskLevels = "high"
	ConditionalAccessConditionSetServicePrincipalRiskLevels_Low    ConditionalAccessConditionSetServicePrincipalRiskLevels = "low"
	ConditionalAccessConditionSetServicePrincipalRiskLevels_Medium ConditionalAccessConditionSetServicePrincipalRiskLevels = "medium"
	ConditionalAccessConditionSetServicePrincipalRiskLevels_None   ConditionalAccessConditionSetServicePrincipalRiskLevels = "none"
)

func PossibleValuesForConditionalAccessConditionSetServicePrincipalRiskLevels() []string {
	return []string{
		string(ConditionalAccessConditionSetServicePrincipalRiskLevels_Hidden),
		string(ConditionalAccessConditionSetServicePrincipalRiskLevels_High),
		string(ConditionalAccessConditionSetServicePrincipalRiskLevels_Low),
		string(ConditionalAccessConditionSetServicePrincipalRiskLevels_Medium),
		string(ConditionalAccessConditionSetServicePrincipalRiskLevels_None),
	}
}

func (s *ConditionalAccessConditionSetServicePrincipalRiskLevels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessConditionSetServicePrincipalRiskLevels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessConditionSetServicePrincipalRiskLevels(input string) (*ConditionalAccessConditionSetServicePrincipalRiskLevels, error) {
	vals := map[string]ConditionalAccessConditionSetServicePrincipalRiskLevels{
		"hidden": ConditionalAccessConditionSetServicePrincipalRiskLevels_Hidden,
		"high":   ConditionalAccessConditionSetServicePrincipalRiskLevels_High,
		"low":    ConditionalAccessConditionSetServicePrincipalRiskLevels_Low,
		"medium": ConditionalAccessConditionSetServicePrincipalRiskLevels_Medium,
		"none":   ConditionalAccessConditionSetServicePrincipalRiskLevels_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessConditionSetServicePrincipalRiskLevels(input)
	return &out, nil
}
