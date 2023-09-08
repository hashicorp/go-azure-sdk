package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessConditionSetUserRiskLevels string

const (
	ConditionalAccessConditionSetUserRiskLevelshidden ConditionalAccessConditionSetUserRiskLevels = "Hidden"
	ConditionalAccessConditionSetUserRiskLevelshigh   ConditionalAccessConditionSetUserRiskLevels = "High"
	ConditionalAccessConditionSetUserRiskLevelslow    ConditionalAccessConditionSetUserRiskLevels = "Low"
	ConditionalAccessConditionSetUserRiskLevelsmedium ConditionalAccessConditionSetUserRiskLevels = "Medium"
	ConditionalAccessConditionSetUserRiskLevelsnone   ConditionalAccessConditionSetUserRiskLevels = "None"
)

func PossibleValuesForConditionalAccessConditionSetUserRiskLevels() []string {
	return []string{
		string(ConditionalAccessConditionSetUserRiskLevelshidden),
		string(ConditionalAccessConditionSetUserRiskLevelshigh),
		string(ConditionalAccessConditionSetUserRiskLevelslow),
		string(ConditionalAccessConditionSetUserRiskLevelsmedium),
		string(ConditionalAccessConditionSetUserRiskLevelsnone),
	}
}

func parseConditionalAccessConditionSetUserRiskLevels(input string) (*ConditionalAccessConditionSetUserRiskLevels, error) {
	vals := map[string]ConditionalAccessConditionSetUserRiskLevels{
		"hidden": ConditionalAccessConditionSetUserRiskLevelshidden,
		"high":   ConditionalAccessConditionSetUserRiskLevelshigh,
		"low":    ConditionalAccessConditionSetUserRiskLevelslow,
		"medium": ConditionalAccessConditionSetUserRiskLevelsmedium,
		"none":   ConditionalAccessConditionSetUserRiskLevelsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessConditionSetUserRiskLevels(input)
	return &out, nil
}
