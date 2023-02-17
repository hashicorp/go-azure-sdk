package prediction

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Category string

const (
	CategoryCost                  Category = "Cost"
	CategoryHighAvailability      Category = "HighAvailability"
	CategoryOperationalExcellence Category = "OperationalExcellence"
	CategoryPerformance           Category = "Performance"
	CategorySecurity              Category = "Security"
)

func PossibleValuesForCategory() []string {
	return []string{
		string(CategoryCost),
		string(CategoryHighAvailability),
		string(CategoryOperationalExcellence),
		string(CategoryPerformance),
		string(CategorySecurity),
	}
}

func parseCategory(input string) (*Category, error) {
	vals := map[string]Category{
		"cost":                  CategoryCost,
		"highavailability":      CategoryHighAvailability,
		"operationalexcellence": CategoryOperationalExcellence,
		"performance":           CategoryPerformance,
		"security":              CategorySecurity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Category(input)
	return &out, nil
}

type Impact string

const (
	ImpactHigh   Impact = "High"
	ImpactLow    Impact = "Low"
	ImpactMedium Impact = "Medium"
)

func PossibleValuesForImpact() []string {
	return []string{
		string(ImpactHigh),
		string(ImpactLow),
		string(ImpactMedium),
	}
}

func parseImpact(input string) (*Impact, error) {
	vals := map[string]Impact{
		"high":   ImpactHigh,
		"low":    ImpactLow,
		"medium": ImpactMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Impact(input)
	return &out, nil
}

type PredictionType string

const (
	PredictionTypePredictiveRightsizing PredictionType = "PredictiveRightsizing"
)

func PossibleValuesForPredictionType() []string {
	return []string{
		string(PredictionTypePredictiveRightsizing),
	}
}

func parsePredictionType(input string) (*PredictionType, error) {
	vals := map[string]PredictionType{
		"predictiverightsizing": PredictionTypePredictiveRightsizing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PredictionType(input)
	return &out, nil
}
