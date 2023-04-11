package usagedetails

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Metrictype string

const (
	MetrictypeActualcost    Metrictype = "actualcost"
	MetrictypeAmortizedcost Metrictype = "amortizedcost"
	MetrictypeUsage         Metrictype = "usage"
)

func PossibleValuesForMetrictype() []string {
	return []string{
		string(MetrictypeActualcost),
		string(MetrictypeAmortizedcost),
		string(MetrictypeUsage),
	}
}

func (s *Metrictype) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMetrictype(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMetrictype(input string) (*Metrictype, error) {
	vals := map[string]Metrictype{
		"actualcost":    MetrictypeActualcost,
		"amortizedcost": MetrictypeAmortizedcost,
		"usage":         MetrictypeUsage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Metrictype(input)
	return &out, nil
}

type PricingModelType string

const (
	PricingModelTypeOnDemand    PricingModelType = "On Demand"
	PricingModelTypeReservation PricingModelType = "Reservation"
	PricingModelTypeSpot        PricingModelType = "Spot"
)

func PossibleValuesForPricingModelType() []string {
	return []string{
		string(PricingModelTypeOnDemand),
		string(PricingModelTypeReservation),
		string(PricingModelTypeSpot),
	}
}

func (s *PricingModelType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePricingModelType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePricingModelType(input string) (*PricingModelType, error) {
	vals := map[string]PricingModelType{
		"on demand":   PricingModelTypeOnDemand,
		"reservation": PricingModelTypeReservation,
		"spot":        PricingModelTypeSpot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PricingModelType(input)
	return &out, nil
}

type UsageDetailsKind string

const (
	UsageDetailsKindLegacy UsageDetailsKind = "legacy"
	UsageDetailsKindModern UsageDetailsKind = "modern"
)

func PossibleValuesForUsageDetailsKind() []string {
	return []string{
		string(UsageDetailsKindLegacy),
		string(UsageDetailsKindModern),
	}
}

func (s *UsageDetailsKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUsageDetailsKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUsageDetailsKind(input string) (*UsageDetailsKind, error) {
	vals := map[string]UsageDetailsKind{
		"legacy": UsageDetailsKindLegacy,
		"modern": UsageDetailsKindModern,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UsageDetailsKind(input)
	return &out, nil
}
