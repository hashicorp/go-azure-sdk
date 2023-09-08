package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundFlowDataDomain string

const (
	IndustryDataInboundFlowDataDomaineducationRostering IndustryDataInboundFlowDataDomain = "EducationRostering"
)

func PossibleValuesForIndustryDataInboundFlowDataDomain() []string {
	return []string{
		string(IndustryDataInboundFlowDataDomaineducationRostering),
	}
}

func parseIndustryDataInboundFlowDataDomain(input string) (*IndustryDataInboundFlowDataDomain, error) {
	vals := map[string]IndustryDataInboundFlowDataDomain{
		"educationrostering": IndustryDataInboundFlowDataDomaineducationRostering,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataInboundFlowDataDomain(input)
	return &out, nil
}
