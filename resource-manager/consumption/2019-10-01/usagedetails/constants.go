package usagedetails

import "strings"

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
