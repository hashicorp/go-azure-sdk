package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipalHistoryItemRiskLevel string

const (
	RiskyServicePrincipalHistoryItemRiskLevelhidden RiskyServicePrincipalHistoryItemRiskLevel = "Hidden"
	RiskyServicePrincipalHistoryItemRiskLevelhigh   RiskyServicePrincipalHistoryItemRiskLevel = "High"
	RiskyServicePrincipalHistoryItemRiskLevellow    RiskyServicePrincipalHistoryItemRiskLevel = "Low"
	RiskyServicePrincipalHistoryItemRiskLevelmedium RiskyServicePrincipalHistoryItemRiskLevel = "Medium"
	RiskyServicePrincipalHistoryItemRiskLevelnone   RiskyServicePrincipalHistoryItemRiskLevel = "None"
)

func PossibleValuesForRiskyServicePrincipalHistoryItemRiskLevel() []string {
	return []string{
		string(RiskyServicePrincipalHistoryItemRiskLevelhidden),
		string(RiskyServicePrincipalHistoryItemRiskLevelhigh),
		string(RiskyServicePrincipalHistoryItemRiskLevellow),
		string(RiskyServicePrincipalHistoryItemRiskLevelmedium),
		string(RiskyServicePrincipalHistoryItemRiskLevelnone),
	}
}

func parseRiskyServicePrincipalHistoryItemRiskLevel(input string) (*RiskyServicePrincipalHistoryItemRiskLevel, error) {
	vals := map[string]RiskyServicePrincipalHistoryItemRiskLevel{
		"hidden": RiskyServicePrincipalHistoryItemRiskLevelhidden,
		"high":   RiskyServicePrincipalHistoryItemRiskLevelhigh,
		"low":    RiskyServicePrincipalHistoryItemRiskLevellow,
		"medium": RiskyServicePrincipalHistoryItemRiskLevelmedium,
		"none":   RiskyServicePrincipalHistoryItemRiskLevelnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyServicePrincipalHistoryItemRiskLevel(input)
	return &out, nil
}
