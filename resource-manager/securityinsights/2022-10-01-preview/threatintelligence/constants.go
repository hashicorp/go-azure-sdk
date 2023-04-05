package threatintelligence

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatIntelligenceResourceKindEnum string

const (
	ThreatIntelligenceResourceKindEnumIndicator ThreatIntelligenceResourceKindEnum = "indicator"
)

func PossibleValuesForThreatIntelligenceResourceKindEnum() []string {
	return []string{
		string(ThreatIntelligenceResourceKindEnumIndicator),
	}
}

type ThreatIntelligenceSortingCriteriaEnum string

const (
	ThreatIntelligenceSortingCriteriaEnumAscending  ThreatIntelligenceSortingCriteriaEnum = "ascending"
	ThreatIntelligenceSortingCriteriaEnumDescending ThreatIntelligenceSortingCriteriaEnum = "descending"
	ThreatIntelligenceSortingCriteriaEnumUnsorted   ThreatIntelligenceSortingCriteriaEnum = "unsorted"
)

func PossibleValuesForThreatIntelligenceSortingCriteriaEnum() []string {
	return []string{
		string(ThreatIntelligenceSortingCriteriaEnumAscending),
		string(ThreatIntelligenceSortingCriteriaEnumDescending),
		string(ThreatIntelligenceSortingCriteriaEnumUnsorted),
	}
}
