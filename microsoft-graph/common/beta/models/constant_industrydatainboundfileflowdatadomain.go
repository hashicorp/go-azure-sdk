package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundFileFlowDataDomain string

const (
	IndustryDataInboundFileFlowDataDomaineducationRostering IndustryDataInboundFileFlowDataDomain = "EducationRostering"
)

func PossibleValuesForIndustryDataInboundFileFlowDataDomain() []string {
	return []string{
		string(IndustryDataInboundFileFlowDataDomaineducationRostering),
	}
}

func parseIndustryDataInboundFileFlowDataDomain(input string) (*IndustryDataInboundFileFlowDataDomain, error) {
	vals := map[string]IndustryDataInboundFileFlowDataDomain{
		"educationrostering": IndustryDataInboundFileFlowDataDomaineducationRostering,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataInboundFileFlowDataDomain(input)
	return &out, nil
}
