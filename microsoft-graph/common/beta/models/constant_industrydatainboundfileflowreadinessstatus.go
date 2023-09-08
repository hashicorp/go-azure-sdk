package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundFileFlowReadinessStatus string

const (
	IndustryDataInboundFileFlowReadinessStatusdisabled IndustryDataInboundFileFlowReadinessStatus = "Disabled"
	IndustryDataInboundFileFlowReadinessStatusexpired  IndustryDataInboundFileFlowReadinessStatus = "Expired"
	IndustryDataInboundFileFlowReadinessStatusfailed   IndustryDataInboundFileFlowReadinessStatus = "Failed"
	IndustryDataInboundFileFlowReadinessStatusnotReady IndustryDataInboundFileFlowReadinessStatus = "NotReady"
	IndustryDataInboundFileFlowReadinessStatusready    IndustryDataInboundFileFlowReadinessStatus = "Ready"
)

func PossibleValuesForIndustryDataInboundFileFlowReadinessStatus() []string {
	return []string{
		string(IndustryDataInboundFileFlowReadinessStatusdisabled),
		string(IndustryDataInboundFileFlowReadinessStatusexpired),
		string(IndustryDataInboundFileFlowReadinessStatusfailed),
		string(IndustryDataInboundFileFlowReadinessStatusnotReady),
		string(IndustryDataInboundFileFlowReadinessStatusready),
	}
}

func parseIndustryDataInboundFileFlowReadinessStatus(input string) (*IndustryDataInboundFileFlowReadinessStatus, error) {
	vals := map[string]IndustryDataInboundFileFlowReadinessStatus{
		"disabled": IndustryDataInboundFileFlowReadinessStatusdisabled,
		"expired":  IndustryDataInboundFileFlowReadinessStatusexpired,
		"failed":   IndustryDataInboundFileFlowReadinessStatusfailed,
		"notready": IndustryDataInboundFileFlowReadinessStatusnotReady,
		"ready":    IndustryDataInboundFileFlowReadinessStatusready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataInboundFileFlowReadinessStatus(input)
	return &out, nil
}
