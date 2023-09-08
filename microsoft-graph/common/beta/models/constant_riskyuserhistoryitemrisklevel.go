package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserHistoryItemRiskLevel string

const (
	RiskyUserHistoryItemRiskLevelhidden RiskyUserHistoryItemRiskLevel = "Hidden"
	RiskyUserHistoryItemRiskLevelhigh   RiskyUserHistoryItemRiskLevel = "High"
	RiskyUserHistoryItemRiskLevellow    RiskyUserHistoryItemRiskLevel = "Low"
	RiskyUserHistoryItemRiskLevelmedium RiskyUserHistoryItemRiskLevel = "Medium"
	RiskyUserHistoryItemRiskLevelnone   RiskyUserHistoryItemRiskLevel = "None"
)

func PossibleValuesForRiskyUserHistoryItemRiskLevel() []string {
	return []string{
		string(RiskyUserHistoryItemRiskLevelhidden),
		string(RiskyUserHistoryItemRiskLevelhigh),
		string(RiskyUserHistoryItemRiskLevellow),
		string(RiskyUserHistoryItemRiskLevelmedium),
		string(RiskyUserHistoryItemRiskLevelnone),
	}
}

func parseRiskyUserHistoryItemRiskLevel(input string) (*RiskyUserHistoryItemRiskLevel, error) {
	vals := map[string]RiskyUserHistoryItemRiskLevel{
		"hidden": RiskyUserHistoryItemRiskLevelhidden,
		"high":   RiskyUserHistoryItemRiskLevelhigh,
		"low":    RiskyUserHistoryItemRiskLevellow,
		"medium": RiskyUserHistoryItemRiskLevelmedium,
		"none":   RiskyUserHistoryItemRiskLevelnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyUserHistoryItemRiskLevel(input)
	return &out, nil
}
