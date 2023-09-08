package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInRiskLevelAggregated string

const (
	SignInRiskLevelAggregatedhidden SignInRiskLevelAggregated = "Hidden"
	SignInRiskLevelAggregatedhigh   SignInRiskLevelAggregated = "High"
	SignInRiskLevelAggregatedlow    SignInRiskLevelAggregated = "Low"
	SignInRiskLevelAggregatedmedium SignInRiskLevelAggregated = "Medium"
	SignInRiskLevelAggregatednone   SignInRiskLevelAggregated = "None"
)

func PossibleValuesForSignInRiskLevelAggregated() []string {
	return []string{
		string(SignInRiskLevelAggregatedhidden),
		string(SignInRiskLevelAggregatedhigh),
		string(SignInRiskLevelAggregatedlow),
		string(SignInRiskLevelAggregatedmedium),
		string(SignInRiskLevelAggregatednone),
	}
}

func parseSignInRiskLevelAggregated(input string) (*SignInRiskLevelAggregated, error) {
	vals := map[string]SignInRiskLevelAggregated{
		"hidden": SignInRiskLevelAggregatedhidden,
		"high":   SignInRiskLevelAggregatedhigh,
		"low":    SignInRiskLevelAggregatedlow,
		"medium": SignInRiskLevelAggregatedmedium,
		"none":   SignInRiskLevelAggregatednone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInRiskLevelAggregated(input)
	return &out, nil
}
