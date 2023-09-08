package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertTemplateSeverity string

const (
	SecurityAlertTemplateSeverityhigh          SecurityAlertTemplateSeverity = "High"
	SecurityAlertTemplateSeverityinformational SecurityAlertTemplateSeverity = "Informational"
	SecurityAlertTemplateSeveritylow           SecurityAlertTemplateSeverity = "Low"
	SecurityAlertTemplateSeveritymedium        SecurityAlertTemplateSeverity = "Medium"
	SecurityAlertTemplateSeverityunknown       SecurityAlertTemplateSeverity = "Unknown"
)

func PossibleValuesForSecurityAlertTemplateSeverity() []string {
	return []string{
		string(SecurityAlertTemplateSeverityhigh),
		string(SecurityAlertTemplateSeverityinformational),
		string(SecurityAlertTemplateSeveritylow),
		string(SecurityAlertTemplateSeveritymedium),
		string(SecurityAlertTemplateSeverityunknown),
	}
}

func parseSecurityAlertTemplateSeverity(input string) (*SecurityAlertTemplateSeverity, error) {
	vals := map[string]SecurityAlertTemplateSeverity{
		"high":          SecurityAlertTemplateSeverityhigh,
		"informational": SecurityAlertTemplateSeverityinformational,
		"low":           SecurityAlertTemplateSeveritylow,
		"medium":        SecurityAlertTemplateSeveritymedium,
		"unknown":       SecurityAlertTemplateSeverityunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertTemplateSeverity(input)
	return &out, nil
}
