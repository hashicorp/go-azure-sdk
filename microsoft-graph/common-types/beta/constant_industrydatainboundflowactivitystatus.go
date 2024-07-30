package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundFlowActivityStatus string

const (
	IndustryDataInboundFlowActivityStatus_Completed             IndustryDataInboundFlowActivityStatus = "completed"
	IndustryDataInboundFlowActivityStatus_CompletedWithErrors   IndustryDataInboundFlowActivityStatus = "completedWithErrors"
	IndustryDataInboundFlowActivityStatus_CompletedWithWarnings IndustryDataInboundFlowActivityStatus = "completedWithWarnings"
	IndustryDataInboundFlowActivityStatus_Failed                IndustryDataInboundFlowActivityStatus = "failed"
	IndustryDataInboundFlowActivityStatus_InProgress            IndustryDataInboundFlowActivityStatus = "inProgress"
	IndustryDataInboundFlowActivityStatus_Skipped               IndustryDataInboundFlowActivityStatus = "skipped"
)

func PossibleValuesForIndustryDataInboundFlowActivityStatus() []string {
	return []string{
		string(IndustryDataInboundFlowActivityStatus_Completed),
		string(IndustryDataInboundFlowActivityStatus_CompletedWithErrors),
		string(IndustryDataInboundFlowActivityStatus_CompletedWithWarnings),
		string(IndustryDataInboundFlowActivityStatus_Failed),
		string(IndustryDataInboundFlowActivityStatus_InProgress),
		string(IndustryDataInboundFlowActivityStatus_Skipped),
	}
}

func (s *IndustryDataInboundFlowActivityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataInboundFlowActivityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataInboundFlowActivityStatus(input string) (*IndustryDataInboundFlowActivityStatus, error) {
	vals := map[string]IndustryDataInboundFlowActivityStatus{
		"completed":             IndustryDataInboundFlowActivityStatus_Completed,
		"completedwitherrors":   IndustryDataInboundFlowActivityStatus_CompletedWithErrors,
		"completedwithwarnings": IndustryDataInboundFlowActivityStatus_CompletedWithWarnings,
		"failed":                IndustryDataInboundFlowActivityStatus_Failed,
		"inprogress":            IndustryDataInboundFlowActivityStatus_InProgress,
		"skipped":               IndustryDataInboundFlowActivityStatus_Skipped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataInboundFlowActivityStatus(input)
	return &out, nil
}
