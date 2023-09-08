package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserRiskLevel string

const (
	RiskyUserRiskLevelhidden RiskyUserRiskLevel = "Hidden"
	RiskyUserRiskLevelhigh   RiskyUserRiskLevel = "High"
	RiskyUserRiskLevellow    RiskyUserRiskLevel = "Low"
	RiskyUserRiskLevelmedium RiskyUserRiskLevel = "Medium"
	RiskyUserRiskLevelnone   RiskyUserRiskLevel = "None"
)

func PossibleValuesForRiskyUserRiskLevel() []string {
	return []string{
		string(RiskyUserRiskLevelhidden),
		string(RiskyUserRiskLevelhigh),
		string(RiskyUserRiskLevellow),
		string(RiskyUserRiskLevelmedium),
		string(RiskyUserRiskLevelnone),
	}
}

func parseRiskyUserRiskLevel(input string) (*RiskyUserRiskLevel, error) {
	vals := map[string]RiskyUserRiskLevel{
		"hidden": RiskyUserRiskLevelhidden,
		"high":   RiskyUserRiskLevelhigh,
		"low":    RiskyUserRiskLevellow,
		"medium": RiskyUserRiskLevelmedium,
		"none":   RiskyUserRiskLevelnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyUserRiskLevel(input)
	return &out, nil
}
