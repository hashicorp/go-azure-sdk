package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundFlowDataDomain string

const (
	IndustryDataInboundFlowDataDomain_EducationRostering IndustryDataInboundFlowDataDomain = "educationRostering"
)

func PossibleValuesForIndustryDataInboundFlowDataDomain() []string {
	return []string{
		string(IndustryDataInboundFlowDataDomain_EducationRostering),
	}
}

func (s *IndustryDataInboundFlowDataDomain) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataInboundFlowDataDomain(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataInboundFlowDataDomain(input string) (*IndustryDataInboundFlowDataDomain, error) {
	vals := map[string]IndustryDataInboundFlowDataDomain{
		"educationrostering": IndustryDataInboundFlowDataDomain_EducationRostering,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataInboundFlowDataDomain(input)
	return &out, nil
}
