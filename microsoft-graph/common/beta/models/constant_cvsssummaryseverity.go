package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CvssSummarySeverity string

const (
	CvssSummarySeveritycritical CvssSummarySeverity = "Critical"
	CvssSummarySeverityhigh     CvssSummarySeverity = "High"
	CvssSummarySeveritylow      CvssSummarySeverity = "Low"
	CvssSummarySeveritymedium   CvssSummarySeverity = "Medium"
	CvssSummarySeveritynone     CvssSummarySeverity = "None"
)

func PossibleValuesForCvssSummarySeverity() []string {
	return []string{
		string(CvssSummarySeveritycritical),
		string(CvssSummarySeverityhigh),
		string(CvssSummarySeveritylow),
		string(CvssSummarySeveritymedium),
		string(CvssSummarySeveritynone),
	}
}

func parseCvssSummarySeverity(input string) (*CvssSummarySeverity, error) {
	vals := map[string]CvssSummarySeverity{
		"critical": CvssSummarySeveritycritical,
		"high":     CvssSummarySeverityhigh,
		"low":      CvssSummarySeveritylow,
		"medium":   CvssSummarySeveritymedium,
		"none":     CvssSummarySeveritynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CvssSummarySeverity(input)
	return &out, nil
}
