package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantAlertRuleSeverity string

const (
	ManagedTenantsManagedTenantAlertRuleSeverityhigh          ManagedTenantsManagedTenantAlertRuleSeverity = "High"
	ManagedTenantsManagedTenantAlertRuleSeverityinformational ManagedTenantsManagedTenantAlertRuleSeverity = "Informational"
	ManagedTenantsManagedTenantAlertRuleSeveritylow           ManagedTenantsManagedTenantAlertRuleSeverity = "Low"
	ManagedTenantsManagedTenantAlertRuleSeveritymedium        ManagedTenantsManagedTenantAlertRuleSeverity = "Medium"
	ManagedTenantsManagedTenantAlertRuleSeverityunknown       ManagedTenantsManagedTenantAlertRuleSeverity = "Unknown"
)

func PossibleValuesForManagedTenantsManagedTenantAlertRuleSeverity() []string {
	return []string{
		string(ManagedTenantsManagedTenantAlertRuleSeverityhigh),
		string(ManagedTenantsManagedTenantAlertRuleSeverityinformational),
		string(ManagedTenantsManagedTenantAlertRuleSeveritylow),
		string(ManagedTenantsManagedTenantAlertRuleSeveritymedium),
		string(ManagedTenantsManagedTenantAlertRuleSeverityunknown),
	}
}

func parseManagedTenantsManagedTenantAlertRuleSeverity(input string) (*ManagedTenantsManagedTenantAlertRuleSeverity, error) {
	vals := map[string]ManagedTenantsManagedTenantAlertRuleSeverity{
		"high":          ManagedTenantsManagedTenantAlertRuleSeverityhigh,
		"informational": ManagedTenantsManagedTenantAlertRuleSeverityinformational,
		"low":           ManagedTenantsManagedTenantAlertRuleSeveritylow,
		"medium":        ManagedTenantsManagedTenantAlertRuleSeveritymedium,
		"unknown":       ManagedTenantsManagedTenantAlertRuleSeverityunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagedTenantAlertRuleSeverity(input)
	return &out, nil
}
