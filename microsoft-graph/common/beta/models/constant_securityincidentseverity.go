package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIncidentSeverity string

const (
	SecurityIncidentSeverityhigh          SecurityIncidentSeverity = "High"
	SecurityIncidentSeverityinformational SecurityIncidentSeverity = "Informational"
	SecurityIncidentSeveritylow           SecurityIncidentSeverity = "Low"
	SecurityIncidentSeveritymedium        SecurityIncidentSeverity = "Medium"
	SecurityIncidentSeverityunknown       SecurityIncidentSeverity = "Unknown"
)

func PossibleValuesForSecurityIncidentSeverity() []string {
	return []string{
		string(SecurityIncidentSeverityhigh),
		string(SecurityIncidentSeverityinformational),
		string(SecurityIncidentSeveritylow),
		string(SecurityIncidentSeveritymedium),
		string(SecurityIncidentSeverityunknown),
	}
}

func parseSecurityIncidentSeverity(input string) (*SecurityIncidentSeverity, error) {
	vals := map[string]SecurityIncidentSeverity{
		"high":          SecurityIncidentSeverityhigh,
		"informational": SecurityIncidentSeverityinformational,
		"low":           SecurityIncidentSeveritylow,
		"medium":        SecurityIncidentSeveritymedium,
		"unknown":       SecurityIncidentSeverityunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIncidentSeverity(input)
	return &out, nil
}
