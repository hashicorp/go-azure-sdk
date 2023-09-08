package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRunStatisticsStatus string

const (
	IndustryDataIndustryDataRunStatisticsStatuscompleted             IndustryDataIndustryDataRunStatisticsStatus = "Completed"
	IndustryDataIndustryDataRunStatisticsStatuscompletedWithErrors   IndustryDataIndustryDataRunStatisticsStatus = "CompletedWithErrors"
	IndustryDataIndustryDataRunStatisticsStatuscompletedWithWarnings IndustryDataIndustryDataRunStatisticsStatus = "CompletedWithWarnings"
	IndustryDataIndustryDataRunStatisticsStatusfailed                IndustryDataIndustryDataRunStatisticsStatus = "Failed"
	IndustryDataIndustryDataRunStatisticsStatusrunning               IndustryDataIndustryDataRunStatisticsStatus = "Running"
)

func PossibleValuesForIndustryDataIndustryDataRunStatisticsStatus() []string {
	return []string{
		string(IndustryDataIndustryDataRunStatisticsStatuscompleted),
		string(IndustryDataIndustryDataRunStatisticsStatuscompletedWithErrors),
		string(IndustryDataIndustryDataRunStatisticsStatuscompletedWithWarnings),
		string(IndustryDataIndustryDataRunStatisticsStatusfailed),
		string(IndustryDataIndustryDataRunStatisticsStatusrunning),
	}
}

func parseIndustryDataIndustryDataRunStatisticsStatus(input string) (*IndustryDataIndustryDataRunStatisticsStatus, error) {
	vals := map[string]IndustryDataIndustryDataRunStatisticsStatus{
		"completed":             IndustryDataIndustryDataRunStatisticsStatuscompleted,
		"completedwitherrors":   IndustryDataIndustryDataRunStatisticsStatuscompletedWithErrors,
		"completedwithwarnings": IndustryDataIndustryDataRunStatisticsStatuscompletedWithWarnings,
		"failed":                IndustryDataIndustryDataRunStatisticsStatusfailed,
		"running":               IndustryDataIndustryDataRunStatisticsStatusrunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataIndustryDataRunStatisticsStatus(input)
	return &out, nil
}
