package usagedetails

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
