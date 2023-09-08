package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertSeverity string

const (
	AlertSeverityhigh          AlertSeverity = "High"
	AlertSeverityinformational AlertSeverity = "Informational"
	AlertSeveritylow           AlertSeverity = "Low"
	AlertSeveritymedium        AlertSeverity = "Medium"
	AlertSeverityunknown       AlertSeverity = "Unknown"
)

func PossibleValuesForAlertSeverity() []string {
	return []string{
		string(AlertSeverityhigh),
		string(AlertSeverityinformational),
		string(AlertSeveritylow),
		string(AlertSeveritymedium),
		string(AlertSeverityunknown),
	}
}

func parseAlertSeverity(input string) (*AlertSeverity, error) {
	vals := map[string]AlertSeverity{
		"high":          AlertSeverityhigh,
		"informational": AlertSeverityinformational,
		"low":           AlertSeveritylow,
		"medium":        AlertSeveritymedium,
		"unknown":       AlertSeverityunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertSeverity(input)
	return &out, nil
}
