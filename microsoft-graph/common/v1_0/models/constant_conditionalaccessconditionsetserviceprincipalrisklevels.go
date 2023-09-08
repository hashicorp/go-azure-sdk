package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessConditionSetServicePrincipalRiskLevels string

const (
	ConditionalAccessConditionSetServicePrincipalRiskLevelshidden ConditionalAccessConditionSetServicePrincipalRiskLevels = "Hidden"
	ConditionalAccessConditionSetServicePrincipalRiskLevelshigh   ConditionalAccessConditionSetServicePrincipalRiskLevels = "High"
	ConditionalAccessConditionSetServicePrincipalRiskLevelslow    ConditionalAccessConditionSetServicePrincipalRiskLevels = "Low"
	ConditionalAccessConditionSetServicePrincipalRiskLevelsmedium ConditionalAccessConditionSetServicePrincipalRiskLevels = "Medium"
	ConditionalAccessConditionSetServicePrincipalRiskLevelsnone   ConditionalAccessConditionSetServicePrincipalRiskLevels = "None"
)

func PossibleValuesForConditionalAccessConditionSetServicePrincipalRiskLevels() []string {
	return []string{
		string(ConditionalAccessConditionSetServicePrincipalRiskLevelshidden),
		string(ConditionalAccessConditionSetServicePrincipalRiskLevelshigh),
		string(ConditionalAccessConditionSetServicePrincipalRiskLevelslow),
		string(ConditionalAccessConditionSetServicePrincipalRiskLevelsmedium),
		string(ConditionalAccessConditionSetServicePrincipalRiskLevelsnone),
	}
}

func parseConditionalAccessConditionSetServicePrincipalRiskLevels(input string) (*ConditionalAccessConditionSetServicePrincipalRiskLevels, error) {
	vals := map[string]ConditionalAccessConditionSetServicePrincipalRiskLevels{
		"hidden": ConditionalAccessConditionSetServicePrincipalRiskLevelshidden,
		"high":   ConditionalAccessConditionSetServicePrincipalRiskLevelshigh,
		"low":    ConditionalAccessConditionSetServicePrincipalRiskLevelslow,
		"medium": ConditionalAccessConditionSetServicePrincipalRiskLevelsmedium,
		"none":   ConditionalAccessConditionSetServicePrincipalRiskLevelsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessConditionSetServicePrincipalRiskLevels(input)
	return &out, nil
}
