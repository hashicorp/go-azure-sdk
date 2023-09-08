package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertTemplateSeverity string

const (
	AlertTemplateSeverityhigh          AlertTemplateSeverity = "High"
	AlertTemplateSeverityinformational AlertTemplateSeverity = "Informational"
	AlertTemplateSeveritylow           AlertTemplateSeverity = "Low"
	AlertTemplateSeveritymedium        AlertTemplateSeverity = "Medium"
	AlertTemplateSeverityunknown       AlertTemplateSeverity = "Unknown"
)

func PossibleValuesForAlertTemplateSeverity() []string {
	return []string{
		string(AlertTemplateSeverityhigh),
		string(AlertTemplateSeverityinformational),
		string(AlertTemplateSeveritylow),
		string(AlertTemplateSeveritymedium),
		string(AlertTemplateSeverityunknown),
	}
}

func parseAlertTemplateSeverity(input string) (*AlertTemplateSeverity, error) {
	vals := map[string]AlertTemplateSeverity{
		"high":          AlertTemplateSeverityhigh,
		"informational": AlertTemplateSeverityinformational,
		"low":           AlertTemplateSeveritylow,
		"medium":        AlertTemplateSeveritymedium,
		"unknown":       AlertTemplateSeverityunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertTemplateSeverity(input)
	return &out, nil
}
