package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertSeverity string

const (
	SecurityAlertSeverityhigh          SecurityAlertSeverity = "High"
	SecurityAlertSeverityinformational SecurityAlertSeverity = "Informational"
	SecurityAlertSeveritylow           SecurityAlertSeverity = "Low"
	SecurityAlertSeveritymedium        SecurityAlertSeverity = "Medium"
	SecurityAlertSeverityunknown       SecurityAlertSeverity = "Unknown"
)

func PossibleValuesForSecurityAlertSeverity() []string {
	return []string{
		string(SecurityAlertSeverityhigh),
		string(SecurityAlertSeverityinformational),
		string(SecurityAlertSeveritylow),
		string(SecurityAlertSeveritymedium),
		string(SecurityAlertSeverityunknown),
	}
}

func parseSecurityAlertSeverity(input string) (*SecurityAlertSeverity, error) {
	vals := map[string]SecurityAlertSeverity{
		"high":          SecurityAlertSeverityhigh,
		"informational": SecurityAlertSeverityinformational,
		"low":           SecurityAlertSeveritylow,
		"medium":        SecurityAlertSeveritymedium,
		"unknown":       SecurityAlertSeverityunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertSeverity(input)
	return &out, nil
}
