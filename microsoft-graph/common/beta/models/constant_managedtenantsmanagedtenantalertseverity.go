package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantAlertSeverity string

const (
	ManagedTenantsManagedTenantAlertSeverityhigh          ManagedTenantsManagedTenantAlertSeverity = "High"
	ManagedTenantsManagedTenantAlertSeverityinformational ManagedTenantsManagedTenantAlertSeverity = "Informational"
	ManagedTenantsManagedTenantAlertSeveritylow           ManagedTenantsManagedTenantAlertSeverity = "Low"
	ManagedTenantsManagedTenantAlertSeveritymedium        ManagedTenantsManagedTenantAlertSeverity = "Medium"
	ManagedTenantsManagedTenantAlertSeverityunknown       ManagedTenantsManagedTenantAlertSeverity = "Unknown"
)

func PossibleValuesForManagedTenantsManagedTenantAlertSeverity() []string {
	return []string{
		string(ManagedTenantsManagedTenantAlertSeverityhigh),
		string(ManagedTenantsManagedTenantAlertSeverityinformational),
		string(ManagedTenantsManagedTenantAlertSeveritylow),
		string(ManagedTenantsManagedTenantAlertSeveritymedium),
		string(ManagedTenantsManagedTenantAlertSeverityunknown),
	}
}

func parseManagedTenantsManagedTenantAlertSeverity(input string) (*ManagedTenantsManagedTenantAlertSeverity, error) {
	vals := map[string]ManagedTenantsManagedTenantAlertSeverity{
		"high":          ManagedTenantsManagedTenantAlertSeverityhigh,
		"informational": ManagedTenantsManagedTenantAlertSeverityinformational,
		"low":           ManagedTenantsManagedTenantAlertSeveritylow,
		"medium":        ManagedTenantsManagedTenantAlertSeveritymedium,
		"unknown":       ManagedTenantsManagedTenantAlertSeverityunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagedTenantAlertSeverity(input)
	return &out, nil
}
