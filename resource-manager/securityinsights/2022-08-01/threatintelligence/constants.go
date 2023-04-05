package threatintelligence

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatIntelligenceResourceInnerKind string

const (
	ThreatIntelligenceResourceInnerKindIndicator ThreatIntelligenceResourceInnerKind = "indicator"
)

func PossibleValuesForThreatIntelligenceResourceInnerKind() []string {
	return []string{
		string(ThreatIntelligenceResourceInnerKindIndicator),
	}
}

type ThreatIntelligenceSortingOrder string

const (
	ThreatIntelligenceSortingOrderAscending  ThreatIntelligenceSortingOrder = "ascending"
	ThreatIntelligenceSortingOrderDescending ThreatIntelligenceSortingOrder = "descending"
	ThreatIntelligenceSortingOrderUnsorted   ThreatIntelligenceSortingOrder = "unsorted"
)

func PossibleValuesForThreatIntelligenceSortingOrder() []string {
	return []string{
		string(ThreatIntelligenceSortingOrderAscending),
		string(ThreatIntelligenceSortingOrderDescending),
		string(ThreatIntelligenceSortingOrderUnsorted),
	}
}
