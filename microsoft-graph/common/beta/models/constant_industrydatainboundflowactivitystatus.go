package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundFlowActivityStatus string

const (
	IndustryDataInboundFlowActivityStatuscompleted             IndustryDataInboundFlowActivityStatus = "Completed"
	IndustryDataInboundFlowActivityStatuscompletedWithErrors   IndustryDataInboundFlowActivityStatus = "CompletedWithErrors"
	IndustryDataInboundFlowActivityStatuscompletedWithWarnings IndustryDataInboundFlowActivityStatus = "CompletedWithWarnings"
	IndustryDataInboundFlowActivityStatusfailed                IndustryDataInboundFlowActivityStatus = "Failed"
	IndustryDataInboundFlowActivityStatusinProgress            IndustryDataInboundFlowActivityStatus = "InProgress"
	IndustryDataInboundFlowActivityStatusskipped               IndustryDataInboundFlowActivityStatus = "Skipped"
)

func PossibleValuesForIndustryDataInboundFlowActivityStatus() []string {
	return []string{
		string(IndustryDataInboundFlowActivityStatuscompleted),
		string(IndustryDataInboundFlowActivityStatuscompletedWithErrors),
		string(IndustryDataInboundFlowActivityStatuscompletedWithWarnings),
		string(IndustryDataInboundFlowActivityStatusfailed),
		string(IndustryDataInboundFlowActivityStatusinProgress),
		string(IndustryDataInboundFlowActivityStatusskipped),
	}
}

func parseIndustryDataInboundFlowActivityStatus(input string) (*IndustryDataInboundFlowActivityStatus, error) {
	vals := map[string]IndustryDataInboundFlowActivityStatus{
		"completed":             IndustryDataInboundFlowActivityStatuscompleted,
		"completedwitherrors":   IndustryDataInboundFlowActivityStatuscompletedWithErrors,
		"completedwithwarnings": IndustryDataInboundFlowActivityStatuscompletedWithWarnings,
		"failed":                IndustryDataInboundFlowActivityStatusfailed,
		"inprogress":            IndustryDataInboundFlowActivityStatusinProgress,
		"skipped":               IndustryDataInboundFlowActivityStatusskipped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataInboundFlowActivityStatus(input)
	return &out, nil
}
