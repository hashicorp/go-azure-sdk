package iotsecuritysolutionsanalytics

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportedSeverity string

const (
	ReportedSeverityHigh          ReportedSeverity = "High"
	ReportedSeverityInformational ReportedSeverity = "Informational"
	ReportedSeverityLow           ReportedSeverity = "Low"
	ReportedSeverityMedium        ReportedSeverity = "Medium"
)

func PossibleValuesForReportedSeverity() []string {
	return []string{
		string(ReportedSeverityHigh),
		string(ReportedSeverityInformational),
		string(ReportedSeverityLow),
		string(ReportedSeverityMedium),
	}
}

func parseReportedSeverity(input string) (*ReportedSeverity, error) {
	vals := map[string]ReportedSeverity{
		"high":          ReportedSeverityHigh,
		"informational": ReportedSeverityInformational,
		"low":           ReportedSeverityLow,
		"medium":        ReportedSeverityMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReportedSeverity(input)
	return &out, nil
}
