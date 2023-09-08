package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundActivityResultsStatus string

const (
	IndustryDataInboundActivityResultsStatuscompleted             IndustryDataInboundActivityResultsStatus = "Completed"
	IndustryDataInboundActivityResultsStatuscompletedWithErrors   IndustryDataInboundActivityResultsStatus = "CompletedWithErrors"
	IndustryDataInboundActivityResultsStatuscompletedWithWarnings IndustryDataInboundActivityResultsStatus = "CompletedWithWarnings"
	IndustryDataInboundActivityResultsStatusfailed                IndustryDataInboundActivityResultsStatus = "Failed"
	IndustryDataInboundActivityResultsStatusinProgress            IndustryDataInboundActivityResultsStatus = "InProgress"
	IndustryDataInboundActivityResultsStatusskipped               IndustryDataInboundActivityResultsStatus = "Skipped"
)

func PossibleValuesForIndustryDataInboundActivityResultsStatus() []string {
	return []string{
		string(IndustryDataInboundActivityResultsStatuscompleted),
		string(IndustryDataInboundActivityResultsStatuscompletedWithErrors),
		string(IndustryDataInboundActivityResultsStatuscompletedWithWarnings),
		string(IndustryDataInboundActivityResultsStatusfailed),
		string(IndustryDataInboundActivityResultsStatusinProgress),
		string(IndustryDataInboundActivityResultsStatusskipped),
	}
}

func parseIndustryDataInboundActivityResultsStatus(input string) (*IndustryDataInboundActivityResultsStatus, error) {
	vals := map[string]IndustryDataInboundActivityResultsStatus{
		"completed":             IndustryDataInboundActivityResultsStatuscompleted,
		"completedwitherrors":   IndustryDataInboundActivityResultsStatuscompletedWithErrors,
		"completedwithwarnings": IndustryDataInboundActivityResultsStatuscompletedWithWarnings,
		"failed":                IndustryDataInboundActivityResultsStatusfailed,
		"inprogress":            IndustryDataInboundActivityResultsStatusinProgress,
		"skipped":               IndustryDataInboundActivityResultsStatusskipped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataInboundActivityResultsStatus(input)
	return &out, nil
}
