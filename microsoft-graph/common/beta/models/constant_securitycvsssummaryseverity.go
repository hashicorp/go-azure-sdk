package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCvssSummarySeverity string

const (
	SecurityCvssSummarySeveritycritical SecurityCvssSummarySeverity = "Critical"
	SecurityCvssSummarySeverityhigh     SecurityCvssSummarySeverity = "High"
	SecurityCvssSummarySeveritylow      SecurityCvssSummarySeverity = "Low"
	SecurityCvssSummarySeveritymedium   SecurityCvssSummarySeverity = "Medium"
	SecurityCvssSummarySeveritynone     SecurityCvssSummarySeverity = "None"
)

func PossibleValuesForSecurityCvssSummarySeverity() []string {
	return []string{
		string(SecurityCvssSummarySeveritycritical),
		string(SecurityCvssSummarySeverityhigh),
		string(SecurityCvssSummarySeveritylow),
		string(SecurityCvssSummarySeveritymedium),
		string(SecurityCvssSummarySeveritynone),
	}
}

func parseSecurityCvssSummarySeverity(input string) (*SecurityCvssSummarySeverity, error) {
	vals := map[string]SecurityCvssSummarySeverity{
		"critical": SecurityCvssSummarySeveritycritical,
		"high":     SecurityCvssSummarySeverityhigh,
		"low":      SecurityCvssSummarySeveritylow,
		"medium":   SecurityCvssSummarySeveritymedium,
		"none":     SecurityCvssSummarySeveritynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCvssSummarySeverity(input)
	return &out, nil
}
