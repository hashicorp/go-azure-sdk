package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataOutboundFlowActivityStatus string

const (
	IndustryDataOutboundFlowActivityStatuscompleted             IndustryDataOutboundFlowActivityStatus = "Completed"
	IndustryDataOutboundFlowActivityStatuscompletedWithErrors   IndustryDataOutboundFlowActivityStatus = "CompletedWithErrors"
	IndustryDataOutboundFlowActivityStatuscompletedWithWarnings IndustryDataOutboundFlowActivityStatus = "CompletedWithWarnings"
	IndustryDataOutboundFlowActivityStatusfailed                IndustryDataOutboundFlowActivityStatus = "Failed"
	IndustryDataOutboundFlowActivityStatusinProgress            IndustryDataOutboundFlowActivityStatus = "InProgress"
	IndustryDataOutboundFlowActivityStatusskipped               IndustryDataOutboundFlowActivityStatus = "Skipped"
)

func PossibleValuesForIndustryDataOutboundFlowActivityStatus() []string {
	return []string{
		string(IndustryDataOutboundFlowActivityStatuscompleted),
		string(IndustryDataOutboundFlowActivityStatuscompletedWithErrors),
		string(IndustryDataOutboundFlowActivityStatuscompletedWithWarnings),
		string(IndustryDataOutboundFlowActivityStatusfailed),
		string(IndustryDataOutboundFlowActivityStatusinProgress),
		string(IndustryDataOutboundFlowActivityStatusskipped),
	}
}

func parseIndustryDataOutboundFlowActivityStatus(input string) (*IndustryDataOutboundFlowActivityStatus, error) {
	vals := map[string]IndustryDataOutboundFlowActivityStatus{
		"completed":             IndustryDataOutboundFlowActivityStatuscompleted,
		"completedwitherrors":   IndustryDataOutboundFlowActivityStatuscompletedWithErrors,
		"completedwithwarnings": IndustryDataOutboundFlowActivityStatuscompletedWithWarnings,
		"failed":                IndustryDataOutboundFlowActivityStatusfailed,
		"inprogress":            IndustryDataOutboundFlowActivityStatusinProgress,
		"skipped":               IndustryDataOutboundFlowActivityStatusskipped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataOutboundFlowActivityStatus(input)
	return &out, nil
}
