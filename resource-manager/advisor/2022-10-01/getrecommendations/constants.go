package getrecommendations

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

type Risk string

const (
	RiskError   Risk = "Error"
	RiskNone    Risk = "None"
	RiskWarning Risk = "Warning"
)

func PossibleValuesForRisk() []string {
	return []string{
		string(RiskError),
		string(RiskNone),
		string(RiskWarning),
	}
}
