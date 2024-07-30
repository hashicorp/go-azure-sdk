package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundFileFlowDataDomain string

const (
	IndustryDataInboundFileFlowDataDomain_EducationRostering IndustryDataInboundFileFlowDataDomain = "educationRostering"
)

func PossibleValuesForIndustryDataInboundFileFlowDataDomain() []string {
	return []string{
		string(IndustryDataInboundFileFlowDataDomain_EducationRostering),
	}
}

func (s *IndustryDataInboundFileFlowDataDomain) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataInboundFileFlowDataDomain(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataInboundFileFlowDataDomain(input string) (*IndustryDataInboundFileFlowDataDomain, error) {
	vals := map[string]IndustryDataInboundFileFlowDataDomain{
		"educationrostering": IndustryDataInboundFileFlowDataDomain_EducationRostering,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataInboundFileFlowDataDomain(input)
	return &out, nil
}
