package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRunActivityStatus string

const (
	IndustryDataIndustryDataRunActivityStatuscompleted             IndustryDataIndustryDataRunActivityStatus = "Completed"
	IndustryDataIndustryDataRunActivityStatuscompletedWithErrors   IndustryDataIndustryDataRunActivityStatus = "CompletedWithErrors"
	IndustryDataIndustryDataRunActivityStatuscompletedWithWarnings IndustryDataIndustryDataRunActivityStatus = "CompletedWithWarnings"
	IndustryDataIndustryDataRunActivityStatusfailed                IndustryDataIndustryDataRunActivityStatus = "Failed"
	IndustryDataIndustryDataRunActivityStatusinProgress            IndustryDataIndustryDataRunActivityStatus = "InProgress"
	IndustryDataIndustryDataRunActivityStatusskipped               IndustryDataIndustryDataRunActivityStatus = "Skipped"
)

func PossibleValuesForIndustryDataIndustryDataRunActivityStatus() []string {
	return []string{
		string(IndustryDataIndustryDataRunActivityStatuscompleted),
		string(IndustryDataIndustryDataRunActivityStatuscompletedWithErrors),
		string(IndustryDataIndustryDataRunActivityStatuscompletedWithWarnings),
		string(IndustryDataIndustryDataRunActivityStatusfailed),
		string(IndustryDataIndustryDataRunActivityStatusinProgress),
		string(IndustryDataIndustryDataRunActivityStatusskipped),
	}
}

func parseIndustryDataIndustryDataRunActivityStatus(input string) (*IndustryDataIndustryDataRunActivityStatus, error) {
	vals := map[string]IndustryDataIndustryDataRunActivityStatus{
		"completed":             IndustryDataIndustryDataRunActivityStatuscompleted,
		"completedwitherrors":   IndustryDataIndustryDataRunActivityStatuscompletedWithErrors,
		"completedwithwarnings": IndustryDataIndustryDataRunActivityStatuscompletedWithWarnings,
		"failed":                IndustryDataIndustryDataRunActivityStatusfailed,
		"inprogress":            IndustryDataIndustryDataRunActivityStatusinProgress,
		"skipped":               IndustryDataIndustryDataRunActivityStatusskipped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataIndustryDataRunActivityStatus(input)
	return &out, nil
}
