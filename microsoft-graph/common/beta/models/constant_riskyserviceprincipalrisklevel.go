package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipalRiskLevel string

const (
	RiskyServicePrincipalRiskLevelhidden RiskyServicePrincipalRiskLevel = "Hidden"
	RiskyServicePrincipalRiskLevelhigh   RiskyServicePrincipalRiskLevel = "High"
	RiskyServicePrincipalRiskLevellow    RiskyServicePrincipalRiskLevel = "Low"
	RiskyServicePrincipalRiskLevelmedium RiskyServicePrincipalRiskLevel = "Medium"
	RiskyServicePrincipalRiskLevelnone   RiskyServicePrincipalRiskLevel = "None"
)

func PossibleValuesForRiskyServicePrincipalRiskLevel() []string {
	return []string{
		string(RiskyServicePrincipalRiskLevelhidden),
		string(RiskyServicePrincipalRiskLevelhigh),
		string(RiskyServicePrincipalRiskLevellow),
		string(RiskyServicePrincipalRiskLevelmedium),
		string(RiskyServicePrincipalRiskLevelnone),
	}
}

func parseRiskyServicePrincipalRiskLevel(input string) (*RiskyServicePrincipalRiskLevel, error) {
	vals := map[string]RiskyServicePrincipalRiskLevel{
		"hidden": RiskyServicePrincipalRiskLevelhidden,
		"high":   RiskyServicePrincipalRiskLevelhigh,
		"low":    RiskyServicePrincipalRiskLevellow,
		"medium": RiskyServicePrincipalRiskLevelmedium,
		"none":   RiskyServicePrincipalRiskLevelnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyServicePrincipalRiskLevel(input)
	return &out, nil
}
