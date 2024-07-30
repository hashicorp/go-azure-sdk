package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundActivityResultsStatus string

const (
	IndustryDataInboundActivityResultsStatus_Completed             IndustryDataInboundActivityResultsStatus = "completed"
	IndustryDataInboundActivityResultsStatus_CompletedWithErrors   IndustryDataInboundActivityResultsStatus = "completedWithErrors"
	IndustryDataInboundActivityResultsStatus_CompletedWithWarnings IndustryDataInboundActivityResultsStatus = "completedWithWarnings"
	IndustryDataInboundActivityResultsStatus_Failed                IndustryDataInboundActivityResultsStatus = "failed"
	IndustryDataInboundActivityResultsStatus_InProgress            IndustryDataInboundActivityResultsStatus = "inProgress"
	IndustryDataInboundActivityResultsStatus_Skipped               IndustryDataInboundActivityResultsStatus = "skipped"
)

func PossibleValuesForIndustryDataInboundActivityResultsStatus() []string {
	return []string{
		string(IndustryDataInboundActivityResultsStatus_Completed),
		string(IndustryDataInboundActivityResultsStatus_CompletedWithErrors),
		string(IndustryDataInboundActivityResultsStatus_CompletedWithWarnings),
		string(IndustryDataInboundActivityResultsStatus_Failed),
		string(IndustryDataInboundActivityResultsStatus_InProgress),
		string(IndustryDataInboundActivityResultsStatus_Skipped),
	}
}

func (s *IndustryDataInboundActivityResultsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataInboundActivityResultsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataInboundActivityResultsStatus(input string) (*IndustryDataInboundActivityResultsStatus, error) {
	vals := map[string]IndustryDataInboundActivityResultsStatus{
		"completed":             IndustryDataInboundActivityResultsStatus_Completed,
		"completedwitherrors":   IndustryDataInboundActivityResultsStatus_CompletedWithErrors,
		"completedwithwarnings": IndustryDataInboundActivityResultsStatus_CompletedWithWarnings,
		"failed":                IndustryDataInboundActivityResultsStatus_Failed,
		"inprogress":            IndustryDataInboundActivityResultsStatus_InProgress,
		"skipped":               IndustryDataInboundActivityResultsStatus_Skipped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataInboundActivityResultsStatus(input)
	return &out, nil
}
