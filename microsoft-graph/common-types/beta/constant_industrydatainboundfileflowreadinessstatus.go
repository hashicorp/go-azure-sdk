package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundFileFlowReadinessStatus string

const (
	IndustryDataInboundFileFlowReadinessStatus_Disabled IndustryDataInboundFileFlowReadinessStatus = "disabled"
	IndustryDataInboundFileFlowReadinessStatus_Expired  IndustryDataInboundFileFlowReadinessStatus = "expired"
	IndustryDataInboundFileFlowReadinessStatus_Failed   IndustryDataInboundFileFlowReadinessStatus = "failed"
	IndustryDataInboundFileFlowReadinessStatus_NotReady IndustryDataInboundFileFlowReadinessStatus = "notReady"
	IndustryDataInboundFileFlowReadinessStatus_Ready    IndustryDataInboundFileFlowReadinessStatus = "ready"
)

func PossibleValuesForIndustryDataInboundFileFlowReadinessStatus() []string {
	return []string{
		string(IndustryDataInboundFileFlowReadinessStatus_Disabled),
		string(IndustryDataInboundFileFlowReadinessStatus_Expired),
		string(IndustryDataInboundFileFlowReadinessStatus_Failed),
		string(IndustryDataInboundFileFlowReadinessStatus_NotReady),
		string(IndustryDataInboundFileFlowReadinessStatus_Ready),
	}
}

func (s *IndustryDataInboundFileFlowReadinessStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataInboundFileFlowReadinessStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataInboundFileFlowReadinessStatus(input string) (*IndustryDataInboundFileFlowReadinessStatus, error) {
	vals := map[string]IndustryDataInboundFileFlowReadinessStatus{
		"disabled": IndustryDataInboundFileFlowReadinessStatus_Disabled,
		"expired":  IndustryDataInboundFileFlowReadinessStatus_Expired,
		"failed":   IndustryDataInboundFileFlowReadinessStatus_Failed,
		"notready": IndustryDataInboundFileFlowReadinessStatus_NotReady,
		"ready":    IndustryDataInboundFileFlowReadinessStatus_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataInboundFileFlowReadinessStatus(input)
	return &out, nil
}
