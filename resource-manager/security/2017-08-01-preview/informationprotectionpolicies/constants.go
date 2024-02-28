package informationprotectionpolicies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionPolicyName string

const (
	InformationProtectionPolicyNameCustom    InformationProtectionPolicyName = "custom"
	InformationProtectionPolicyNameEffective InformationProtectionPolicyName = "effective"
)

func PossibleValuesForInformationProtectionPolicyName() []string {
	return []string{
		string(InformationProtectionPolicyNameCustom),
		string(InformationProtectionPolicyNameEffective),
	}
}

func (s *InformationProtectionPolicyName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInformationProtectionPolicyName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInformationProtectionPolicyName(input string) (*InformationProtectionPolicyName, error) {
	vals := map[string]InformationProtectionPolicyName{
		"custom":    InformationProtectionPolicyNameCustom,
		"effective": InformationProtectionPolicyNameEffective,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InformationProtectionPolicyName(input)
	return &out, nil
}

type Rank string

const (
	RankCritical Rank = "Critical"
	RankHigh     Rank = "High"
	RankLow      Rank = "Low"
	RankMedium   Rank = "Medium"
	RankNone     Rank = "None"
)

func PossibleValuesForRank() []string {
	return []string{
		string(RankCritical),
		string(RankHigh),
		string(RankLow),
		string(RankMedium),
		string(RankNone),
	}
}

func (s *Rank) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRank(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRank(input string) (*Rank, error) {
	vals := map[string]Rank{
		"critical": RankCritical,
		"high":     RankHigh,
		"low":      RankLow,
		"medium":   RankMedium,
		"none":     RankNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Rank(input)
	return &out, nil
}
