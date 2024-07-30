package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundFlowReadinessStatus string

const (
	IndustryDataInboundFlowReadinessStatus_Disabled IndustryDataInboundFlowReadinessStatus = "disabled"
	IndustryDataInboundFlowReadinessStatus_Expired  IndustryDataInboundFlowReadinessStatus = "expired"
	IndustryDataInboundFlowReadinessStatus_Failed   IndustryDataInboundFlowReadinessStatus = "failed"
	IndustryDataInboundFlowReadinessStatus_NotReady IndustryDataInboundFlowReadinessStatus = "notReady"
	IndustryDataInboundFlowReadinessStatus_Ready    IndustryDataInboundFlowReadinessStatus = "ready"
)

func PossibleValuesForIndustryDataInboundFlowReadinessStatus() []string {
	return []string{
		string(IndustryDataInboundFlowReadinessStatus_Disabled),
		string(IndustryDataInboundFlowReadinessStatus_Expired),
		string(IndustryDataInboundFlowReadinessStatus_Failed),
		string(IndustryDataInboundFlowReadinessStatus_NotReady),
		string(IndustryDataInboundFlowReadinessStatus_Ready),
	}
}

func (s *IndustryDataInboundFlowReadinessStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataInboundFlowReadinessStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataInboundFlowReadinessStatus(input string) (*IndustryDataInboundFlowReadinessStatus, error) {
	vals := map[string]IndustryDataInboundFlowReadinessStatus{
		"disabled": IndustryDataInboundFlowReadinessStatus_Disabled,
		"expired":  IndustryDataInboundFlowReadinessStatus_Expired,
		"failed":   IndustryDataInboundFlowReadinessStatus_Failed,
		"notready": IndustryDataInboundFlowReadinessStatus_NotReady,
		"ready":    IndustryDataInboundFlowReadinessStatus_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataInboundFlowReadinessStatus(input)
	return &out, nil
}
