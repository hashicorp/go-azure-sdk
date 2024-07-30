package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataOutboundFlowActivityStatus string

const (
	IndustryDataOutboundFlowActivityStatus_Completed             IndustryDataOutboundFlowActivityStatus = "completed"
	IndustryDataOutboundFlowActivityStatus_CompletedWithErrors   IndustryDataOutboundFlowActivityStatus = "completedWithErrors"
	IndustryDataOutboundFlowActivityStatus_CompletedWithWarnings IndustryDataOutboundFlowActivityStatus = "completedWithWarnings"
	IndustryDataOutboundFlowActivityStatus_Failed                IndustryDataOutboundFlowActivityStatus = "failed"
	IndustryDataOutboundFlowActivityStatus_InProgress            IndustryDataOutboundFlowActivityStatus = "inProgress"
	IndustryDataOutboundFlowActivityStatus_Skipped               IndustryDataOutboundFlowActivityStatus = "skipped"
)

func PossibleValuesForIndustryDataOutboundFlowActivityStatus() []string {
	return []string{
		string(IndustryDataOutboundFlowActivityStatus_Completed),
		string(IndustryDataOutboundFlowActivityStatus_CompletedWithErrors),
		string(IndustryDataOutboundFlowActivityStatus_CompletedWithWarnings),
		string(IndustryDataOutboundFlowActivityStatus_Failed),
		string(IndustryDataOutboundFlowActivityStatus_InProgress),
		string(IndustryDataOutboundFlowActivityStatus_Skipped),
	}
}

func (s *IndustryDataOutboundFlowActivityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataOutboundFlowActivityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataOutboundFlowActivityStatus(input string) (*IndustryDataOutboundFlowActivityStatus, error) {
	vals := map[string]IndustryDataOutboundFlowActivityStatus{
		"completed":             IndustryDataOutboundFlowActivityStatus_Completed,
		"completedwitherrors":   IndustryDataOutboundFlowActivityStatus_CompletedWithErrors,
		"completedwithwarnings": IndustryDataOutboundFlowActivityStatus_CompletedWithWarnings,
		"failed":                IndustryDataOutboundFlowActivityStatus_Failed,
		"inprogress":            IndustryDataOutboundFlowActivityStatus_InProgress,
		"skipped":               IndustryDataOutboundFlowActivityStatus_Skipped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataOutboundFlowActivityStatus(input)
	return &out, nil
}
