package threatintelligence

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *ThreatIntelligenceResourceInnerKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatIntelligenceResourceInnerKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatIntelligenceResourceInnerKind(input string) (*ThreatIntelligenceResourceInnerKind, error) {
	vals := map[string]ThreatIntelligenceResourceInnerKind{
		"indicator": ThreatIntelligenceResourceInnerKindIndicator,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatIntelligenceResourceInnerKind(input)
	return &out, nil
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

func (s *ThreatIntelligenceSortingOrder) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatIntelligenceSortingOrder(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatIntelligenceSortingOrder(input string) (*ThreatIntelligenceSortingOrder, error) {
	vals := map[string]ThreatIntelligenceSortingOrder{
		"ascending":  ThreatIntelligenceSortingOrderAscending,
		"descending": ThreatIntelligenceSortingOrderDescending,
		"unsorted":   ThreatIntelligenceSortingOrderUnsorted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatIntelligenceSortingOrder(input)
	return &out, nil
}
