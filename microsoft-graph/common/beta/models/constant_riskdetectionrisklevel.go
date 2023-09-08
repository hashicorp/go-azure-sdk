package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionRiskLevel string

const (
	RiskDetectionRiskLevelhidden RiskDetectionRiskLevel = "Hidden"
	RiskDetectionRiskLevelhigh   RiskDetectionRiskLevel = "High"
	RiskDetectionRiskLevellow    RiskDetectionRiskLevel = "Low"
	RiskDetectionRiskLevelmedium RiskDetectionRiskLevel = "Medium"
	RiskDetectionRiskLevelnone   RiskDetectionRiskLevel = "None"
)

func PossibleValuesForRiskDetectionRiskLevel() []string {
	return []string{
		string(RiskDetectionRiskLevelhidden),
		string(RiskDetectionRiskLevelhigh),
		string(RiskDetectionRiskLevellow),
		string(RiskDetectionRiskLevelmedium),
		string(RiskDetectionRiskLevelnone),
	}
}

func parseRiskDetectionRiskLevel(input string) (*RiskDetectionRiskLevel, error) {
	vals := map[string]RiskDetectionRiskLevel{
		"hidden": RiskDetectionRiskLevelhidden,
		"high":   RiskDetectionRiskLevelhigh,
		"low":    RiskDetectionRiskLevellow,
		"medium": RiskDetectionRiskLevelmedium,
		"none":   RiskDetectionRiskLevelnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionRiskLevel(input)
	return &out, nil
}
