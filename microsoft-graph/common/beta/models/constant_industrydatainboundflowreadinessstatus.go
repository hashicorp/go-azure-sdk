package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundFlowReadinessStatus string

const (
	IndustryDataInboundFlowReadinessStatusdisabled IndustryDataInboundFlowReadinessStatus = "Disabled"
	IndustryDataInboundFlowReadinessStatusexpired  IndustryDataInboundFlowReadinessStatus = "Expired"
	IndustryDataInboundFlowReadinessStatusfailed   IndustryDataInboundFlowReadinessStatus = "Failed"
	IndustryDataInboundFlowReadinessStatusnotReady IndustryDataInboundFlowReadinessStatus = "NotReady"
	IndustryDataInboundFlowReadinessStatusready    IndustryDataInboundFlowReadinessStatus = "Ready"
)

func PossibleValuesForIndustryDataInboundFlowReadinessStatus() []string {
	return []string{
		string(IndustryDataInboundFlowReadinessStatusdisabled),
		string(IndustryDataInboundFlowReadinessStatusexpired),
		string(IndustryDataInboundFlowReadinessStatusfailed),
		string(IndustryDataInboundFlowReadinessStatusnotReady),
		string(IndustryDataInboundFlowReadinessStatusready),
	}
}

func parseIndustryDataInboundFlowReadinessStatus(input string) (*IndustryDataInboundFlowReadinessStatus, error) {
	vals := map[string]IndustryDataInboundFlowReadinessStatus{
		"disabled": IndustryDataInboundFlowReadinessStatusdisabled,
		"expired":  IndustryDataInboundFlowReadinessStatusexpired,
		"failed":   IndustryDataInboundFlowReadinessStatusfailed,
		"notready": IndustryDataInboundFlowReadinessStatusnotReady,
		"ready":    IndustryDataInboundFlowReadinessStatusready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataInboundFlowReadinessStatus(input)
	return &out, nil
}
