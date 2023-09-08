package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessConditionSetSignInRiskLevels string

const (
	ConditionalAccessConditionSetSignInRiskLevelshidden ConditionalAccessConditionSetSignInRiskLevels = "Hidden"
	ConditionalAccessConditionSetSignInRiskLevelshigh   ConditionalAccessConditionSetSignInRiskLevels = "High"
	ConditionalAccessConditionSetSignInRiskLevelslow    ConditionalAccessConditionSetSignInRiskLevels = "Low"
	ConditionalAccessConditionSetSignInRiskLevelsmedium ConditionalAccessConditionSetSignInRiskLevels = "Medium"
	ConditionalAccessConditionSetSignInRiskLevelsnone   ConditionalAccessConditionSetSignInRiskLevels = "None"
)

func PossibleValuesForConditionalAccessConditionSetSignInRiskLevels() []string {
	return []string{
		string(ConditionalAccessConditionSetSignInRiskLevelshidden),
		string(ConditionalAccessConditionSetSignInRiskLevelshigh),
		string(ConditionalAccessConditionSetSignInRiskLevelslow),
		string(ConditionalAccessConditionSetSignInRiskLevelsmedium),
		string(ConditionalAccessConditionSetSignInRiskLevelsnone),
	}
}

func parseConditionalAccessConditionSetSignInRiskLevels(input string) (*ConditionalAccessConditionSetSignInRiskLevels, error) {
	vals := map[string]ConditionalAccessConditionSetSignInRiskLevels{
		"hidden": ConditionalAccessConditionSetSignInRiskLevelshidden,
		"high":   ConditionalAccessConditionSetSignInRiskLevelshigh,
		"low":    ConditionalAccessConditionSetSignInRiskLevelslow,
		"medium": ConditionalAccessConditionSetSignInRiskLevelsmedium,
		"none":   ConditionalAccessConditionSetSignInRiskLevelsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessConditionSetSignInRiskLevels(input)
	return &out, nil
}
